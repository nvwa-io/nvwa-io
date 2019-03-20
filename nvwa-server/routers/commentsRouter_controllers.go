package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:AppController"] = append(beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:AppController"],
		beego.ControllerComments{
			Method: "Create",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:AppController"] = append(beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:AppController"],
		beego.ControllerComments{
			Method: "Update",
			Router: `/:app_id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:AppController"] = append(beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:AppController"],
		beego.ControllerComments{
			Method: "Detail",
			Router: `/:app_id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:AppController"] = append(beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:AppController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:app_id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:AppController"] = append(beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:AppController"],
		beego.ControllerComments{
			Method: "Branches",
			Router: `/:app_id/branches`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:AppController"] = append(beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:AppController"],
		beego.ControllerComments{
			Method: "UpdateCmd",
			Router: `/:app_id/commands`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:AppController"] = append(beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:AppController"],
		beego.ControllerComments{
			Method: "AdminList",
			Router: `/admin/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:AppController"] = append(beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:AppController"],
		beego.ControllerComments{
			Method: "ListAppAndEnvByAppId",
			Router: `/app-and-envs/app/:app_id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:AppController"] = append(beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:AppController"],
		beego.ControllerComments{
			Method: "ListAppAndEnv",
			Router: `/app-and-envs/project/:project_id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:AppController"] = append(beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:AppController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/project/:project_id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:AuditController"] = append(beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:AuditController"],
		beego.ControllerComments{
			Method: "Cancel",
			Router: `/:audit_id/cancel`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:AuditController"] = append(beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:AuditController"],
		beego.ControllerComments{
			Method: "Pass",
			Router: `/:audit_id/pass`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:AuditController"] = append(beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:AuditController"],
		beego.ControllerComments{
			Method: "Reject",
			Router: `/:audit_id/reject`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:AuditController"] = append(beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:AuditController"],
		beego.ControllerComments{
			Method: "ListByStatus",
			Router: `/admin/status/:status`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:AuditController"] = append(beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:AuditController"],
		beego.ControllerComments{
			Method: "ListAudited",
			Router: `/audited`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:AuditController"] = append(beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:AuditController"],
		beego.ControllerComments{
			Method: "ListMine",
			Router: `/mine`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:AuditController"] = append(beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:AuditController"],
		beego.ControllerComments{
			Method: "ListWait",
			Router: `/wait`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:AuditController"] = append(beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:AuditController"],
		beego.ControllerComments{
			Method: "GetWaitCount",
			Router: `/wait-num`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:BuildController"] = append(beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:BuildController"],
		beego.ControllerComments{
			Method: "Create",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:BuildController"] = append(beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:BuildController"],
		beego.ControllerComments{
			Method: "Detail",
			Router: `/:build_id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:BuildController"] = append(beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:BuildController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/project/:project_id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:ClusterController"] = append(beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:ClusterController"],
		beego.ControllerComments{
			Method: "Create",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:ClusterController"] = append(beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:ClusterController"],
		beego.ControllerComments{
			Method: "Update",
			Router: `/:cluster_id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:ClusterController"] = append(beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:ClusterController"],
		beego.ControllerComments{
			Method: "Detail",
			Router: `/:cluster_id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:ClusterController"] = append(beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:ClusterController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:cluster_id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:ClusterController"] = append(beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:ClusterController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/env/:env_id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:DeploymentController"] = append(beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:DeploymentController"],
		beego.ControllerComments{
			Method: "Create",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:DeploymentController"] = append(beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:DeploymentController"],
		beego.ControllerComments{
			Method: "Audit",
			Router: `/audit/:deployment_id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:DeploymentController"] = append(beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:DeploymentController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/project/:project_id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:EnvController"] = append(beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:EnvController"],
		beego.ControllerComments{
			Method: "Create",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:EnvController"] = append(beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:EnvController"],
		beego.ControllerComments{
			Method: "Update",
			Router: `/:env_id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:EnvController"] = append(beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:EnvController"],
		beego.ControllerComments{
			Method: "Detail",
			Router: `/:env_id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:EnvController"] = append(beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:EnvController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:env_id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:EnvController"] = append(beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:EnvController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/app/:app_id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:EventController"] = append(beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:EventController"],
		beego.ControllerComments{
			Method: "Create",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:EventController"] = append(beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:EventController"],
		beego.ControllerComments{
			Method: "Update",
			Router: `/:event_id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:EventController"] = append(beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:EventController"],
		beego.ControllerComments{
			Method: "Detail",
			Router: `/:event_id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:EventController"] = append(beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:EventController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:event_id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:EventController"] = append(beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:EventController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/user/:uid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:JobController"] = append(beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:JobController"],
		beego.ControllerComments{
			Method: "StartJob",
			Router: `/:job_id/start`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:JobStepController"] = append(beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:JobStepController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/job/:job_id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:MemberController"] = append(beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:MemberController"],
		beego.ControllerComments{
			Method: "Add",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:MemberController"] = append(beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:MemberController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:member_id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:MemberController"] = append(beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:MemberController"],
		beego.ControllerComments{
			Method: "UpdateRole",
			Router: `/:member_id/role`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:MemberController"] = append(beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:MemberController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/project/:project_id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:PkgController"] = append(beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:PkgController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/env/:env_id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:ProjectController"] = append(beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:ProjectController"],
		beego.ControllerComments{
			Method: "Create",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:ProjectController"] = append(beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:ProjectController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:ProjectController"] = append(beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:ProjectController"],
		beego.ControllerComments{
			Method: "Update",
			Router: `/:project_id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:ProjectController"] = append(beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:ProjectController"],
		beego.ControllerComments{
			Method: "Detail",
			Router: `/:project_id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:ProjectController"] = append(beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:ProjectController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:project_id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:ProjectController"] = append(beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:ProjectController"],
		beego.ControllerComments{
			Method: "AdminList",
			Router: `/admin/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:ProjectRoleController"] = append(beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:ProjectRoleController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:ProjectRoleController"] = append(beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:ProjectRoleController"],
		beego.ControllerComments{
			Method: "Detail",
			Router: `/:project_role_id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:ProjectRoleController"] = append(beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:ProjectRoleController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:project_role_id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:ProjectRoleController"] = append(beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:ProjectRoleController"],
		beego.ControllerComments{
			Method: "Create",
			Router: `/admin/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:ProjectRoleController"] = append(beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:ProjectRoleController"],
		beego.ControllerComments{
			Method: "Update",
			Router: `/admin/:project_role_id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:ProjectRolePermController"] = append(beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:ProjectRolePermController"],
		beego.ControllerComments{
			Method: "Create",
			Router: `/admin/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:ProjectRolePermController"] = append(beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:ProjectRolePermController"],
		beego.ControllerComments{
			Method: "All",
			Router: `/admin/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:ProjectRolePermController"] = append(beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:ProjectRolePermController"],
		beego.ControllerComments{
			Method: "BatchCreate",
			Router: `/admin/batch-create`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:ProjectRolePermController"] = append(beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:ProjectRolePermController"],
		beego.ControllerComments{
			Method: "BatchUpdate",
			Router: `/admin/batch-update`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:ProjectRolePermController"] = append(beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:ProjectRolePermController"],
		beego.ControllerComments{
			Method: "ProjectRolePermList",
			Router: `/project-roles/:project_role_id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:SiteController"] = append(beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:SiteController"],
		beego.ControllerComments{
			Method: "Index",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:SiteController"] = append(beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:SiteController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:SiteController"] = append(beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:SiteController"],
		beego.ControllerComments{
			Method: "Register",
			Router: `/register`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:SiteController"] = append(beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:SiteController"],
		beego.ControllerComments{
			Method: "GetUserByToken",
			Router: `/token/user`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:SystemController"] = append(beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:SystemController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:UserController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/admin/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:UserController"],
		beego.ControllerComments{
			Method: "UpdateRole",
			Router: `/admin/role`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:UserController"],
		beego.ControllerComments{
			Method: "All",
			Router: `/all`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/nvwa-io/nvwa-io/nvwa-server/controllers:UserController"],
		beego.ControllerComments{
			Method: "Detail",
			Router: `/detail`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
