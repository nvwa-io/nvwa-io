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
	"github.com/astaxie/beego/validation"
	"github.com/go-ozzo/ozzo-dbx"
	. "github.com/nvwa-io/nvwa-io/nvwa-server/daos"
	"github.com/nvwa-io/nvwa-io/nvwa-server/entities/vo"
	"github.com/nvwa-io/nvwa-io/nvwa-server/lang"
	"github.com/nvwa-io/nvwa-io/nvwa-server/libs/errs"
	"github.com/nvwa-io/nvwa-io/nvwa-server/libs/logger"
	. "github.com/nvwa-io/nvwa-io/nvwa-server/svrs"
)

// @Title Create new project role
// @router /admin/ [post]
func (t *ProjectRoleController) Create() {
	// json decode request
	req := new(vo.ReqProjectRole)
	err := t.ReadRequestJson(&req)
	if err != nil {
		t.FailJson(errs.ERR_PARAM, err.Error())
		return
	}

	// @TODO validate request params

	// check whether project role exist
	ok, err := DefaultProjectRoleSvr.IsExist(req.ProjectRole.Name)
	if err != nil {
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}
	if ok {
		t.FailJson(errs.ERR_PROJECT_ROLE_EXIST, req.ProjectRole.Name)
		return
	}

	id, err := DefaultProjectRoleSvr.Create(req.ProjectRole.Name)
	if err != nil {
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}

	t.SuccJson(map[string]interface{}{
		"id": id,
	})
}

// @Title Update project role
// @router /admin/:project_role_id [put]
func (t *ProjectRoleController) Update() {
	id, err := t.GetInt64(":project_role_id")
	if err != nil {
		t.FailJson(errs.ERR_PARAM, ":project_role_id")
		return
	}

	req := new(vo.ReqProjectRole)
	err = t.ReadRequestJson(req)
	logger.Debugf("Role %v", req)
	if err != nil {
		t.FailJson(errs.ERR_PARAM, err.Error())
		return
	}

	// project role detail
	project, err := DefaultProjectRoleSvr.GetById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			t.FailJson(errs.ERR_NO_RECORD)
			return
		}
	}

	// @TODO check where has permission to operate
	logger.Debugf("project %v", project)

	// check whether project role exist
	ok, err := DefaultProjectRoleSvr.IsExist(req.ProjectRole.Name, id)
	if err != nil {
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}
	if ok {
		t.FailJson(errs.ERR_PROJECT_ROLE_EXIST, req.ProjectRole.Name)
		return
	}

	// update project role
	_, err = DefaultProjectRoleDao.UpdateById(id, dbx.Params{
		"name": req.ProjectRole.Name,
	})
	if err != nil {
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}

	t.SuccJson()
}

// @Title Delete project role
// @router /:project_role_id [delete]
func (t *ProjectRoleController) Delete() {
	id, err := t.GetInt64(":project_role_id")
	if err != nil {
		t.FailJson(errs.ERR_PARAM, ":project_role_id")
		return
	}

	valid := validation.Validation{}
	valid.Min(id, 1, "project_role_id").Message(lang.I("project_role.id.invalid"))
	if valid.HasErrors() {
		t.FailJson(errs.ERR_PARAM, valid.Errors[0].Error())
		return
	}

	err = DefaultProjectRoleSvr.DeleteById(id)
	if err != nil {
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}

	t.SuccJson()
}
