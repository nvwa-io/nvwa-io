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
	"github.com/nvwa-io/nvwa-io/nvwa-server/entities"
	"github.com/nvwa-io/nvwa-io/nvwa-server/entities/vo"
	"github.com/nvwa-io/nvwa-io/nvwa-server/lang"
	"github.com/nvwa-io/nvwa-io/nvwa-server/libs/errs"
	. "github.com/nvwa-io/nvwa-io/nvwa-server/svrs"
)

// @Title Add perm to role
// @router /admin/ [post]
func (t *ProjectRolePermController) Create() {
	// json decode request
	req := new(vo.ReqProjectRolePerm)
	err := t.ReadRequestJson(&req)
	if err != nil {
		t.FailJson(errs.ERR_PARAM, err.Error())
		return
	}

	// validate request params
	err = req.Valid()
	if err != nil {
		t.FailJson(errs.ERR_PARAM, err.Error())
		return
	}

	// check whether role exist
	if tmp, _ := DefaultProjectRoleSvr.GetById(req.Perm.ProjectRoleId); tmp == nil {
		t.FailJson(errs.ERR_NO_RECORD, lang.I("project_role.not.exist"))
		return
	}

	// check whether perm has been bind to project role
	ok, err := DefaultProjectRolePermSvr.IsExist(req.Perm.ProjectRoleId, req.Perm.Perm)
	if err != nil {
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}
	if ok {
		t.SuccJson()
		return
	}

	// bind perm to project role
	id, err := DefaultProjectRolePermSvr.Create(req.Perm.ProjectRoleId, req.Perm.Perm)
	if err != nil {
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}

	t.SuccJson(RespData{"id": id})
}

// @Title Batch create to project role permissions
// @router /admin/batch-create [post]
func (t *ProjectRolePermController) BatchCreate() {
	// json decode request
	req := new(vo.ReqBatchCreateProjectRolePerm)
	err := t.ReadRequestJson(&req)
	if err != nil {
		t.FailJson(errs.ERR_PARAM, err.Error())
		return
	}

	// validate request params
	err = req.Valid()
	if err != nil {
		t.FailJson(errs.ERR_PARAM, err.Error())
		return
	}

	// check whether role exist
	if tmp, _ := DefaultProjectRoleSvr.GetByName(req.ProjectRoleName); tmp != nil {
		t.FailJson(errs.ERR_NO_RECORD, lang.I("project_role.exist"))
		return
	}

	err = DefaultProjectRolePermSvr.CreateProjectRoleAndBindPerms(req.ProjectRoleName, req.Perms)
	if err != nil {
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}

	t.SuccJson(RespData{})
}

// @Title Batch update to project role
// @router /admin/batch-update [post]
func (t *ProjectRolePermController) BatchUpdate() {
	// json decode request
	req := new(vo.ReqBatchProjectRolePerm)
	err := t.ReadRequestJson(&req)
	if err != nil {
		t.FailJson(errs.ERR_PARAM, err.Error())
		return
	}

	// validate request params
	err = req.Valid()
	if err != nil {
		t.FailJson(errs.ERR_PARAM, err.Error())
		return
	}

	// check whether role exist
	if tmp, _ := DefaultProjectRoleSvr.GetById(req.ProjectRoleId); tmp == nil {
		t.FailJson(errs.ERR_NO_RECORD, lang.I("project_role.not.exist"))
		return
	}

	// check whether duplicated project role name
	tmp, err := DefaultProjectRoleSvr.GetByName(req.ProjectRoleName)
	if err != nil {
		if err != sql.ErrNoRows {
			t.FailJson(errs.ERR_OPERATE, err.Error())
			return
		}
	}
	if tmp != nil && tmp.Id != req.ProjectRoleId {
		t.FailJson(errs.ERR_OPERATE, lang.I("project_role.exist"))
		return
	}

	// batch update project role's permissions
	err = DefaultProjectRolePermSvr.BatchUpdate(req.ProjectRoleId, req.ProjectRoleName, req.Perms)
	if err != nil {
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}

	t.SuccJson()
}

// @Title get all perms
// @router /admin/ [get]
func (t *ProjectRolePermController) All() {
	t.SuccJson(RespData{
		"list": entities.PERM_LABELS,
	})
}
