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
	"database/sql"
	"github.com/nvwa-io/nvwa-io/nvwa-server/libs/errs"
	. "github.com/nvwa-io/nvwa-io/nvwa-server/svrs"
)

type ProjectRoleController struct {
	BaseAuthController
}

// @Title Get project role detail
// @router /:project_role_id [put]
func (t *ProjectRoleController) Detail() {

}

// @Title Get all project roles
// @router / [get]
func (t *ProjectRoleController) List() {
	list, err := DefaultProjectRoleSvr.ListAll()
	if err != nil {
		if err != sql.ErrNoRows {
			t.FailJson(errs.ERR_OPERATE, err.Error())
			return
		}
	}

	// get project role permissions
	rolePerms, err := DefaultProjectRolePermSvr.GetAllProjectRoleIdMap()
	if err != nil {
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}

	// format project roles and role permissions
	res := make([]map[string]interface{}, 0)
	for _, v := range list {
		tmp := map[string]interface{}{
			"project_role":      v,
			"project_role_perm": rolePerms[v.Id],
		}
		res = append(res, tmp)
	}

	t.SuccJson(RespData{
		"list": res,
	})
}
