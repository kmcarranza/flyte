/*
 * flyteidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package flyteadmin

type CoreTaskLog struct {
	Uri string `json:"uri,omitempty"`
	Name string `json:"name,omitempty"`
	MessageFormat *TaskLogMessageFormat `json:"message_format,omitempty"`
	Ttl string `json:"ttl,omitempty"`
}
