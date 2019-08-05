// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: flyteidl/service/admin.proto

#include "flyteidl/service/admin.pb.h"

#include <algorithm>

#include <google/protobuf/stubs/common.h>
#include <google/protobuf/stubs/port.h>
#include <google/protobuf/stubs/once.h>
#include <google/protobuf/io/coded_stream.h>
#include <google/protobuf/wire_format_lite_inl.h>
#include <google/protobuf/descriptor.h>
#include <google/protobuf/generated_message_reflection.h>
#include <google/protobuf/reflection_ops.h>
#include <google/protobuf/wire_format.h>
// This is a temporary google only hack
#ifdef GOOGLE_PROTOBUF_ENFORCE_UNIQUENESS
#include "third_party/protobuf/version.h"
#endif
// @@protoc_insertion_point(includes)
namespace flyteidl {
namespace service {
}  // namespace service
}  // namespace flyteidl
namespace protobuf_flyteidl_2fservice_2fadmin_2eproto {
const ::google::protobuf::uint32 TableStruct::offsets[1] = {};
static const ::google::protobuf::internal::MigrationSchema* schemas = NULL;
static const ::google::protobuf::Message* const* file_default_instances = NULL;

void protobuf_AssignDescriptors() {
  AddDescriptors();
  ::google::protobuf::MessageFactory* factory = NULL;
  AssignDescriptors(
      "flyteidl/service/admin.proto", schemas, file_default_instances, TableStruct::offsets, factory,
      NULL, NULL, NULL);
}

void protobuf_AssignDescriptorsOnce() {
  static GOOGLE_PROTOBUF_DECLARE_ONCE(once);
  ::google::protobuf::GoogleOnceInit(&once, &protobuf_AssignDescriptors);
}

void protobuf_RegisterTypes(const ::std::string&) GOOGLE_PROTOBUF_ATTRIBUTE_COLD;
void protobuf_RegisterTypes(const ::std::string&) {
  protobuf_AssignDescriptorsOnce();
}

void AddDescriptorsImpl() {
  InitDefaults();
  static const char descriptor[] GOOGLE_PROTOBUF_ATTRIBUTE_SECTION_VARIABLE(protodesc_cold) = {
      "\n\034flyteidl/service/admin.proto\022\020flyteidl"
      ".service\032\034google/api/annotations.proto\032\034"
      "flyteidl/admin/project.proto\032\031flyteidl/a"
      "dmin/task.proto\032\035flyteidl/admin/workflow"
      ".proto\032 flyteidl/admin/launch_plan.proto"
      "\032\032flyteidl/admin/event.proto\032\036flyteidl/a"
      "dmin/execution.proto\032#flyteidl/admin/nod"
      "e_execution.proto\032#flyteidl/admin/task_e"
      "xecution.proto\032\033flyteidl/admin/common.pr"
      "oto\032,protoc-gen-swagger/options/annotati"
      "ons.proto2\3018\n\014AdminService\022\304\002\n\nCreateTas"
      "k\022!.flyteidl.admin.TaskCreateRequest\032\".f"
      "lyteidl.admin.TaskCreateResponse\"\356\001\202\323\344\223\002"
      "\022\"\r/api/v1/tasks:\001*\222A\322\001\032%Create and regi"
      "ster a task definitionJB\n\003400\022;\n9Returne"
      "d for bad request that may have failed v"
      "alidation.Je\n\003409\022^\n\\Returned for a requ"
      "est that references an identical entity "
      "that has already been registered.\022\210\001\n\007Ge"
      "tTask\022 .flyteidl.admin.ObjectGetRequest\032"
      "\024.flyteidl.admin.Task\"E\202\323\344\223\002\?\022=/api/v1/t"
      "asks/{id.project}/{id.domain}/{id.name}/"
      "{id.version}\022\227\001\n\013ListTaskIds\0220.flyteidl."
      "admin.NamedEntityIdentifierListRequest\032)"
      ".flyteidl.admin.NamedEntityIdentifierLis"
      "t\"+\202\323\344\223\002%\022#/api/v1/task_ids/{project}/{d"
      "omain}\022\256\001\n\tListTasks\022#.flyteidl.admin.Re"
      "sourceListRequest\032\030.flyteidl.admin.TaskL"
      "ist\"b\202\323\344\223\002\\\0220/api/v1/tasks/{id.project}/"
      "{id.domain}/{id.name}Z(\022&/api/v1/tasks/{"
      "id.project}/{id.domain}\022\330\002\n\016CreateWorkfl"
      "ow\022%.flyteidl.admin.WorkflowCreateReques"
      "t\032&.flyteidl.admin.WorkflowCreateRespons"
      "e\"\366\001\202\323\344\223\002\026\"\021/api/v1/workflows:\001*\222A\326\001\032)Cr"
      "eate and register a workflow definitionJ"
      "B\n\003400\022;\n9Returned for bad request that "
      "may have failed validation.Je\n\003409\022^\n\\Re"
      "turned for a request that references an "
      "identical entity that has already been r"
      "egistered.\022\224\001\n\013GetWorkflow\022 .flyteidl.ad"
      "min.ObjectGetRequest\032\030.flyteidl.admin.Wo"
      "rkflow\"I\202\323\344\223\002C\022A/api/v1/workflows/{id.pr"
      "oject}/{id.domain}/{id.name}/{id.version"
      "}\022\237\001\n\017ListWorkflowIds\0220.flyteidl.admin.N"
      "amedEntityIdentifierListRequest\032).flytei"
      "dl.admin.NamedEntityIdentifierList\"/\202\323\344\223"
      "\002)\022\'/api/v1/workflow_ids/{project}/{doma"
      "in}\022\276\001\n\rListWorkflows\022#.flyteidl.admin.R"
      "esourceListRequest\032\034.flyteidl.admin.Work"
      "flowList\"j\202\323\344\223\002d\0224/api/v1/workflows/{id."
      "project}/{id.domain}/{id.name}Z,\022*/api/v"
      "1/workflows/{id.project}/{id.domain}\022\344\002\n"
      "\020CreateLaunchPlan\022\'.flyteidl.admin.Launc"
      "hPlanCreateRequest\032(.flyteidl.admin.Laun"
      "chPlanCreateResponse\"\374\001\202\323\344\223\002\031\"\024/api/v1/l"
      "aunch_plans:\001*\222A\331\001\032,Create and register "
      "a launch plan definitionJB\n\003400\022;\n9Retur"
      "ned for bad request that may have failed"
      " validation.Je\n\003409\022^\n\\Returned for a re"
      "quest that references an identical entit"
      "y that has already been registered.\022\233\001\n\r"
      "GetLaunchPlan\022 .flyteidl.admin.ObjectGet"
      "Request\032\032.flyteidl.admin.LaunchPlan\"L\202\323\344"
      "\223\002F\022D/api/v1/launch_plans/{id.project}/{"
      "id.domain}/{id.name}/{id.version}\022\244\001\n\021Li"
      "stLaunchPlanIds\0220.flyteidl.admin.NamedEn"
      "tityIdentifierListRequest\032).flyteidl.adm"
      "in.NamedEntityIdentifierList\"2\202\323\344\223\002,\022*/a"
      "pi/v1/launch_plan_ids/{project}/{domain}"
      "\022\310\001\n\017ListLaunchPlans\022#.flyteidl.admin.Re"
      "sourceListRequest\032\036.flyteidl.admin.Launc"
      "hPlanList\"p\202\323\344\223\002j\0227/api/v1/launch_plans/"
      "{id.project}/{id.domain}/{id.name}Z/\022-/a"
      "pi/v1/launch_plans/{id.project}/{id.doma"
      "in}\022\266\001\n\020UpdateLaunchPlan\022\'.flyteidl.admi"
      "n.LaunchPlanUpdateRequest\032(.flyteidl.adm"
      "in.LaunchPlanUpdateResponse\"O\202\323\344\223\002I\032D/ap"
      "i/v1/launch_plans/{id.project}/{id.domai"
      "n}/{id.name}/{id.version}:\001*\022\316\002\n\017CreateE"
      "xecution\022&.flyteidl.admin.ExecutionCreat"
      "eRequest\032\'.flyteidl.admin.ExecutionCreat"
      "eResponse\"\351\001\202\323\344\223\002\027\"\022/api/v1/executions:\001"
      "*\222A\310\001\032\033Create a workflow executionJB\n\00340"
      "0\022;\n9Returned for bad request that may h"
      "ave failed validation.Je\n\003409\022^\n\\Returne"
      "d for a request that references an ident"
      "ical entity that has already been regist"
      "ered.\022\335\002\n\021RelaunchExecution\022(.flyteidl.a"
      "dmin.ExecutionRelaunchRequest\032\'.flyteidl"
      ".admin.ExecutionCreateResponse\"\364\001\202\323\344\223\002 \""
      "\033/api/v1/executions/relaunch:\001*\222A\312\001\032\035Rel"
      "aunch a workflow executionJB\n\003400\022;\n9Ret"
      "urned for bad request that may have fail"
      "ed validation.Je\n\003409\022^\n\\Returned for a "
      "request that references an identical ent"
      "ity that has already been registered.\022\225\001"
      "\n\014GetExecution\022+.flyteidl.admin.Workflow"
      "ExecutionGetRequest\032\031.flyteidl.admin.Exe"
      "cution\"=\202\323\344\223\0027\0225/api/v1/executions/{id.p"
      "roject}/{id.domain}/{id.name}\022\271\001\n\020GetExe"
      "cutionData\022/.flyteidl.admin.WorkflowExec"
      "utionGetDataRequest\0320.flyteidl.admin.Wor"
      "kflowExecutionGetDataResponse\"B\202\323\344\223\002<\022:/"
      "api/v1/data/executions/{id.project}/{id."
      "domain}/{id.name}\022\211\001\n\016ListExecutions\022#.f"
      "lyteidl.admin.ResourceListRequest\032\035.flyt"
      "eidl.admin.ExecutionList\"3\202\323\344\223\002-\022+/api/v"
      "1/executions/{id.project}/{id.domain}\022\255\001"
      "\n\022TerminateExecution\022).flyteidl.admin.Ex"
      "ecutionTerminateRequest\032*.flyteidl.admin"
      ".ExecutionTerminateResponse\"@\202\323\344\223\002:*5/ap"
      "i/v1/executions/{id.project}/{id.domain}"
      "/{id.name}:\001*\022\322\001\n\020GetNodeExecution\022\'.fly"
      "teidl.admin.NodeExecutionGetRequest\032\035.fl"
      "yteidl.admin.NodeExecution\"v\202\323\344\223\002p\022n/api"
      "/v1/node_executions/{id.execution_id.pro"
      "ject}/{id.execution_id.domain}/{id.execu"
      "tion_id.name}/{id.node_id}\022\336\001\n\022ListNodeE"
      "xecutions\022(.flyteidl.admin.NodeExecution"
      "ListRequest\032!.flyteidl.admin.NodeExecuti"
      "onList\"{\202\323\344\223\002u\022s/api/v1/node_executions/"
      "{workflow_execution_id.project}/{workflo"
      "w_execution_id.domain}/{workflow_executi"
      "on_id.name}\022\245\004\n\031ListNodeExecutionsForTas"
      "k\022/.flyteidl.admin.NodeExecutionForTaskL"
      "istRequest\032!.flyteidl.admin.NodeExecutio"
      "nList\"\263\003\202\323\344\223\002\254\003\022\251\003/api/v1/children/task_"
      "executions/{task_execution_id.node_execu"
      "tion_id.execution_id.project}/{task_exec"
      "ution_id.node_execution_id.execution_id."
      "domain}/{task_execution_id.node_executio"
      "n_id.execution_id.name}/{task_execution_"
      "id.node_execution_id.node_id}/{task_exec"
      "ution_id.task_id.project}/{task_executio"
      "n_id.task_id.domain}/{task_execution_id."
      "task_id.name}/{task_execution_id.task_id"
      ".version}/{task_execution_id.retry_attem"
      "pt}\022\356\001\n\024GetNodeExecutionData\022+.flyteidl."
      "admin.NodeExecutionGetDataRequest\032,.flyt"
      "eidl.admin.NodeExecutionGetDataResponse\""
      "{\202\323\344\223\002u\022s/api/v1/data/node_executions/{i"
      "d.execution_id.project}/{id.execution_id"
      ".domain}/{id.execution_id.name}/{id.node"
      "_id}\022\245\002\n\017RegisterProject\022&.flyteidl.admi"
      "n.ProjectRegisterRequest\032\'.flyteidl.admi"
      "n.ProjectRegisterResponse\"\300\001\202\323\344\223\002\025\"\020/api"
      "/v1/projects:\001*\222A\241\001\032+Register a project "
      "along with valid domainsJ.\n\003201\022\'\n%Retur"
      "ned for successful registration.JB\n\003400\022"
      ";\n9Returned for bad request that may hav"
      "e failed validation.\022f\n\014ListProjects\022\".f"
      "lyteidl.admin.ProjectListRequest\032\030.flyte"
      "idl.admin.Projects\"\030\202\323\344\223\002\022\022\020/api/v1/proj"
      "ects\022\231\001\n\023CreateWorkflowEvent\022-.flyteidl."
      "admin.WorkflowExecutionEventRequest\032..fl"
      "yteidl.admin.WorkflowExecutionEventRespo"
      "nse\"#\202\323\344\223\002\035\"\030/api/v1/events/workflows:\001*"
      "\022\211\001\n\017CreateNodeEvent\022).flyteidl.admin.No"
      "deExecutionEventRequest\032*.flyteidl.admin"
      ".NodeExecutionEventResponse\"\037\202\323\344\223\002\031\"\024/ap"
      "i/v1/events/nodes:\001*\022\211\001\n\017CreateTaskEvent"
      "\022).flyteidl.admin.TaskExecutionEventRequ"
      "est\032*.flyteidl.admin.TaskExecutionEventR"
      "esponse\"\037\202\323\344\223\002\031\"\024/api/v1/events/tasks:\001*"
      "\022\200\003\n\020GetTaskExecution\022\'.flyteidl.admin.T"
      "askExecutionGetRequest\032\035.flyteidl.admin."
      "TaskExecution\"\243\002\202\323\344\223\002\234\002\022\231\002/api/v1/task_e"
      "xecutions/{id.node_execution_id.executio"
      "n_id.project}/{id.node_execution_id.exec"
      "ution_id.domain}/{id.node_execution_id.e"
      "xecution_id.name}/{id.node_execution_id."
      "node_id}/{id.task_id.project}/{id.task_i"
      "d.domain}/{id.task_id.name}/{id.task_id."
      "version}/{id.retry_attempt}\022\230\002\n\022ListTask"
      "Executions\022(.flyteidl.admin.TaskExecutio"
      "nListRequest\032!.flyteidl.admin.TaskExecut"
      "ionList\"\264\001\202\323\344\223\002\255\001\022\252\001/api/v1/task_executi"
      "ons/{node_execution_id.execution_id.proj"
      "ect}/{node_execution_id.execution_id.dom"
      "ain}/{node_execution_id.execution_id.nam"
      "e}/{node_execution_id.node_id}\022\234\003\n\024GetTa"
      "skExecutionData\022+.flyteidl.admin.TaskExe"
      "cutionGetDataRequest\032,.flyteidl.admin.Ta"
      "skExecutionGetDataResponse\"\250\002\202\323\344\223\002\241\002\022\236\002/"
      "api/v1/data/task_executions/{id.node_exe"
      "cution_id.execution_id.project}/{id.node"
      "_execution_id.execution_id.domain}/{id.n"
      "ode_execution_id.execution_id.name}/{id."
      "node_execution_id.node_id}/{id.task_id.p"
      "roject}/{id.task_id.domain}/{id.task_id."
      "name}/{id.task_id.version}/{id.retry_att"
      "empt}B5Z3github.com/lyft/flyteidl/gen/pb"
      "-go/flyteidl/serviceb\006proto3"
  };
  ::google::protobuf::DescriptorPool::InternalAddGeneratedFile(
      descriptor, 7708);
  ::google::protobuf::MessageFactory::InternalRegisterGeneratedFile(
    "flyteidl/service/admin.proto", &protobuf_RegisterTypes);
  ::protobuf_google_2fapi_2fannotations_2eproto::AddDescriptors();
  ::protobuf_flyteidl_2fadmin_2fproject_2eproto::AddDescriptors();
  ::protobuf_flyteidl_2fadmin_2ftask_2eproto::AddDescriptors();
  ::protobuf_flyteidl_2fadmin_2fworkflow_2eproto::AddDescriptors();
  ::protobuf_flyteidl_2fadmin_2flaunch_5fplan_2eproto::AddDescriptors();
  ::protobuf_flyteidl_2fadmin_2fevent_2eproto::AddDescriptors();
  ::protobuf_flyteidl_2fadmin_2fexecution_2eproto::AddDescriptors();
  ::protobuf_flyteidl_2fadmin_2fnode_5fexecution_2eproto::AddDescriptors();
  ::protobuf_flyteidl_2fadmin_2ftask_5fexecution_2eproto::AddDescriptors();
  ::protobuf_flyteidl_2fadmin_2fcommon_2eproto::AddDescriptors();
  ::protobuf_protoc_2dgen_2dswagger_2foptions_2fannotations_2eproto::AddDescriptors();
}

void AddDescriptors() {
  static GOOGLE_PROTOBUF_DECLARE_ONCE(once);
  ::google::protobuf::GoogleOnceInit(&once, &AddDescriptorsImpl);
}
// Force AddDescriptors() to be called at dynamic initialization time.
struct StaticDescriptorInitializer {
  StaticDescriptorInitializer() {
    AddDescriptors();
  }
} static_descriptor_initializer;
}  // namespace protobuf_flyteidl_2fservice_2fadmin_2eproto
namespace flyteidl {
namespace service {

// @@protoc_insertion_point(namespace_scope)
}  // namespace service
}  // namespace flyteidl

// @@protoc_insertion_point(global_scope)
