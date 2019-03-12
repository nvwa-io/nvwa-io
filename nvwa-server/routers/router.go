// Copyright 2019 - now The https://github.com/nvwa-io/nvwa-io Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/nvwa-io/nvwa-io/nvwa-server/controllers"
	"github.com/nvwa-io/nvwa-io/nvwa-server/controllers/apis"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/site",
			beego.NSInclude(
				&controllers.SiteController{},
			),
		),
		beego.NSNamespace("/users",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/projects",
			beego.NSInclude(
				&controllers.ProjectController{},
			),
		),
		beego.NSNamespace("/audits",
			beego.NSInclude(
				&controllers.AuditController{},
			),
		),
		beego.NSNamespace("/project-roles",
			beego.NSInclude(
				&controllers.ProjectRoleController{},
			),
		),
		beego.NSNamespace("/project-role-perms",
			beego.NSInclude(
				&controllers.ProjectRolePermController{},
			),
		),
		beego.NSNamespace("/members",
			beego.NSInclude(
				&controllers.MemberController{},
			),
		),
		beego.NSNamespace("/apps",
			beego.NSInclude(
				&controllers.AppController{},
			),
		),
		beego.NSNamespace("/envs",
			beego.NSInclude(
				&controllers.EnvController{},
			),
		),
		beego.NSNamespace("/clusters",
			beego.NSInclude(
				&controllers.ClusterController{},
			),
		),
		beego.NSNamespace("/builds",
			beego.NSInclude(
				&controllers.BuildController{},
			),
		),
		beego.NSNamespace("/pkgs",
			beego.NSInclude(
				&controllers.PkgController{},
			),
		),
		beego.NSNamespace("/deployments",
			beego.NSInclude(
				&controllers.DeploymentController{},
			),
		),
		beego.NSNamespace("/jobs",
			beego.NSInclude(
				&controllers.JobController{},
			),
		),
		beego.NSNamespace("/job-steps",
			beego.NSInclude(
				&controllers.JobStepController{},
			),
		),
		beego.NSNamespace("/systems",
			beego.NSInclude(
				&controllers.SystemController{},
			),
		),
	)

	apiNs := beego.NewNamespace("/api/v1",
		beego.NSNamespace("/build",
			beego.NSInclude(
				&apis.BuildApi{},
			),
		),
		beego.NSNamespace("/notify",
			beego.NSInclude(
				&apis.NotifyApi{},
			),
		),
	)

	beego.AddNamespace(ns)
	beego.AddNamespace(apiNs)
}
