package impl

import (
	"context"

	"github.com/lyft/datacatalog/pkg/errors"
	"github.com/lyft/datacatalog/pkg/manager/impl/validators"
	"github.com/lyft/datacatalog/pkg/manager/interfaces"
	"github.com/lyft/datacatalog/pkg/repositories"
	datacatalog "github.com/lyft/datacatalog/protos/gen"

	"github.com/lyft/datacatalog/pkg/repositories/models"
	"github.com/lyft/datacatalog/pkg/repositories/transformers"

	"github.com/lyft/flytestdlib/contextutils"
	"github.com/lyft/flytestdlib/logger"
	"github.com/lyft/flytestdlib/promutils"
	"github.com/lyft/flytestdlib/promutils/labeled"
	"github.com/lyft/flytestdlib/storage"
	"google.golang.org/grpc/codes"
)

type artifactMetrics struct {
	scope                    promutils.Scope
	createSuccessCounter     labeled.Counter
	createFailureCounter     labeled.Counter
	getSuccessCounter        labeled.Counter
	getFailureCounter        labeled.Counter
	createDataFailureCounter labeled.Counter
	createDataSuccessCounter labeled.Counter
	transformerErrorCounter  labeled.Counter
	validationErrorCounter   labeled.Counter
	alreadyExistsCounter     labeled.Counter
	doesNotExistCounter      labeled.Counter
}

type artifactManager struct {
	repo          repositories.RepositoryInterface
	artifactStore ArtifactDataStore
	systemMetrics artifactMetrics
}

// Create an Artifact along with the associated ArtifactData. The ArtifactData will be stored in an offloaded location.
func (m *artifactManager) CreateArtifact(ctx context.Context, request datacatalog.CreateArtifactRequest) (*datacatalog.CreateArtifactResponse, error) {
	artifact := request.Artifact
	err := validators.ValidateArtifact(artifact)
	if err != nil {
		logger.Warningf(ctx, "Invalid create artifact request %v, err: %v", request, err)
		m.systemMetrics.validationErrorCounter.Inc(ctx)
		return nil, err
	}

	ctx = contextutils.WithProjectDomain(ctx, artifact.Dataset.Project, artifact.Dataset.Domain)
	datasetKey := transformers.FromDatasetID(*artifact.Dataset)

	// The dataset must exist for the artifact, let's verify that first
	_, err = m.repo.DatasetRepo().Get(ctx, datasetKey)
	if err != nil {
		logger.Warnf(ctx, "Failed to get dataset for artifact creation %v, err: %v", datasetKey, err)
		m.systemMetrics.createFailureCounter.Inc(ctx)
		return nil, err
	}

	// create Artifact Data offloaded storage files
	artifactDataModels := make([]models.ArtifactData, len(request.Artifact.Data))
	for i, artifactData := range request.Artifact.Data {
		dataLocation, err := m.artifactStore.PutData(ctx, *artifact, *artifactData)
		if err != nil {
			logger.Errorf(ctx, "Failed to store artifact data err: %v", err)
			m.systemMetrics.createDataFailureCounter.Inc(ctx)
			return nil, err
		}

		artifactDataModels[i].Name = artifactData.Name
		artifactDataModels[i].Location = dataLocation.String()
		m.systemMetrics.createDataSuccessCounter.Inc(ctx)
	}

	logger.Debugf(ctx, "Stored %v data for artifact %+v", len(artifactDataModels), artifact.Id)

	artifactModel, err := transformers.CreateArtifactModel(request, artifactDataModels)
	if err != nil {
		logger.Errorf(ctx, "Failed to transform artifact err: %v", err)
		m.systemMetrics.transformerErrorCounter.Inc(ctx)
		return nil, err
	}

	err = m.repo.ArtifactRepo().Create(ctx, artifactModel)
	if err != nil {
		if errors.IsAlreadyExistsError(err) {
			logger.Warnf(ctx, "Artifact already exists key: %+v, err %v", artifact.Id, err)
			m.systemMetrics.alreadyExistsCounter.Inc(ctx)
		} else {
			logger.Errorf(ctx, "Failed to create artifact %v, err: %v", artifactDataModels, err)
			m.systemMetrics.createFailureCounter.Inc(ctx)
		}
		return nil, err
	}

	logger.Debugf(ctx, "Successfully created artifact id: %v", artifact.Id)

	m.systemMetrics.createSuccessCounter.Inc(ctx)
	return &datacatalog.CreateArtifactResponse{}, nil
}

// Get the Artifact and its associated ArtifactData. The request can query by ArtifactID or TagName.
func (m *artifactManager) GetArtifact(ctx context.Context, request datacatalog.GetArtifactRequest) (*datacatalog.GetArtifactResponse, error) {
	datasetID := request.Dataset
	err := validators.ValidateGetArtifactRequest(request)
	if err != nil {
		logger.Warningf(ctx, "Invalid get artifact request %v, err: %v", request, err)
		m.systemMetrics.validationErrorCounter.Inc(ctx)
		return nil, err
	}

	ctx = contextutils.WithProjectDomain(ctx, datasetID.Project, datasetID.Domain)
	var artifactModel models.Artifact
	switch request.QueryHandle.(type) {
	case *datacatalog.GetArtifactRequest_ArtifactId:
		logger.Debugf(ctx, "Get artifact by id %v", request.GetArtifactId())
		artifactKey := transformers.ToArtifactKey(*datasetID, request.GetArtifactId())
		artifactModel, err = m.repo.ArtifactRepo().Get(ctx, artifactKey)

		if err != nil {
			if errors.IsDoesNotExistError(err) {
				logger.Warnf(ctx, "Artifact does not exist id: %+v, err %v", request.GetArtifactId(), err)
				m.systemMetrics.doesNotExistCounter.Inc(ctx)
			} else {
				logger.Errorf(ctx, "Unable to retrieve artifact by id: %+v, err %v", request.GetArtifactId(), err)
				m.systemMetrics.getFailureCounter.Inc(ctx)
			}
			return nil, err
		}
	case *datacatalog.GetArtifactRequest_TagName:
		logger.Debugf(ctx, "Get artifact by id %v", request.GetTagName())
		tagKey := transformers.ToTagKey(*datasetID, request.GetTagName())
		tag, err := m.repo.TagRepo().Get(ctx, tagKey)

		if err != nil {
			if errors.IsDoesNotExistError(err) {
				logger.Warnf(ctx, "Artifact does not exist tag: %+v, err %v", request.GetTagName(), err)
				m.systemMetrics.doesNotExistCounter.Inc(ctx)
			} else {
				logger.Errorf(ctx, "Unable to retrieve Artifact by tag %v, err: %v", request.GetTagName(), err)
				m.systemMetrics.getFailureCounter.Inc(ctx)
			}
			return nil, err
		}

		artifactModel = tag.Artifact
	}

	if len(artifactModel.ArtifactData) == 0 {
		return nil, errors.NewDataCatalogErrorf(codes.Internal, "artifact [%+v] does not have artifact data associated", request)
	}

	artifact, err := transformers.FromArtifactModel(artifactModel)
	if err != nil {
		logger.Errorf(ctx, "Error in transforming get artifact request %+v, err %v", artifactModel, err)
		m.systemMetrics.transformerErrorCounter.Inc(ctx)
		return nil, err
	}

	artifactDataList := make([]*datacatalog.ArtifactData, len(artifactModel.ArtifactData))
	for i, artifactData := range artifactModel.ArtifactData {
		value, err := m.artifactStore.GetData(ctx, artifactData)
		if err != nil {
			logger.Errorf(ctx, "Error in getting artifact data from datastore %+v, err %v", artifactData.Location, err)
			return nil, err
		}

		artifactDataList[i] = &datacatalog.ArtifactData{
			Name:  artifactData.Name,
			Value: value,
		}
	}
	artifact.Data = artifactDataList

	logger.Debugf(ctx, "Retrieved artifact dataset %v, id: %v", artifact.Dataset, artifact.Id)
	m.systemMetrics.getSuccessCounter.Inc(ctx)
	return &datacatalog.GetArtifactResponse{
		Artifact: &artifact,
	}, nil
}

func NewArtifactManager(repo repositories.RepositoryInterface, store *storage.DataStore, storagePrefix storage.DataReference, artifactScope promutils.Scope) interfaces.ArtifactManager {
	artifactMetrics := artifactMetrics{
		scope:                    artifactScope,
		createSuccessCounter:     labeled.NewCounter("create_artifact_success_count", "The number of times create artifact was called", artifactScope, labeled.EmitUnlabeledMetric),
		getSuccessCounter:        labeled.NewCounter("get_artifact_success_count", "The number of times get artifact was called", artifactScope, labeled.EmitUnlabeledMetric),
		createFailureCounter:     labeled.NewCounter("create_artifact_failure_count", "The number of times create artifact failed", artifactScope, labeled.EmitUnlabeledMetric),
		getFailureCounter:        labeled.NewCounter("get_artifact_failure_count", "The number of times get artifact failed", artifactScope, labeled.EmitUnlabeledMetric),
		createDataFailureCounter: labeled.NewCounter("create_artifact_data_failure_count", "The number of times create artifact data failed", artifactScope, labeled.EmitUnlabeledMetric),
		createDataSuccessCounter: labeled.NewCounter("create_artifact_data_succeeded_count", "The number of times create artifact data succeeded", artifactScope, labeled.EmitUnlabeledMetric),
		transformerErrorCounter:  labeled.NewCounter("transformer_failed_count", "The number of times transformations failed", artifactScope, labeled.EmitUnlabeledMetric),
		validationErrorCounter:   labeled.NewCounter("validation_failed_count", "The number of times validation failed", artifactScope, labeled.EmitUnlabeledMetric),
		alreadyExistsCounter:     labeled.NewCounter("already_exists_count", "The number of times an artifact already exists", artifactScope, labeled.EmitUnlabeledMetric),
		doesNotExistCounter:      labeled.NewCounter("does_not_exists_count", "The number of times an artifact was not found", artifactScope, labeled.EmitUnlabeledMetric),
	}

	return &artifactManager{
		repo:          repo,
		artifactStore: NewArtifactDataStore(store, storagePrefix),
		systemMetrics: artifactMetrics,
	}
}
