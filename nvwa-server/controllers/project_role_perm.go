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

package controllers

import (
	"github.com/astaxie/beego/validation"
	"github.com/nvwa-io/nvwa-io/nvwa-server/entities"
	"github.com/nvwa-io/nvwa-io/nvwa-server/lang"
	"github.com/nvwa-io/nvwa-io/nvwa-server/libs/errs"
	. "github.com/nvwa-io/nvwa-io/nvwa-server/svrs"
)

type ProjectRolePermController struct {
	BaseAuthController
}

// @Title List all perm of project role
// @router /project-roles/:project_role_id [get]
func (t *ProjectRolePermController) ProjectRolePermList() {
	projectRoleId, err := t.GetInt64(":project_role_id")
	if err != nil {
		t.FailJson(errs.ERR_PARAM, ":project_role_id")
		return
	}

	valid := validation.Validation{}
	valid.Min(projectRoleId, 1, "project_role_id").Message(lang.I("project_role.id.invalid"))
	if valid.HasErrors() {
		t.FailJson(errs.ERR_PARAM, valid.Errors[0].Error())
		return
	}

	list, err := DefaultProjectRolePermSvr.ListAll(projectRoleId)
	if err != nil {
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}

	if len(list) == 0 {
		t.SuccJson(RespData{
			"project_role_perms": []entities.ProjectRolePermEntity{},
		})
		return
	}

	t.SuccJson(RespData{
		"project_role_perms": list,
	})
}
