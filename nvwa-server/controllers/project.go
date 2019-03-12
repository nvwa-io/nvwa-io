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
	. "github.com/nvwa-io/nvwa-io/nvwa-server/entities"
	"github.com/nvwa-io/nvwa-io/nvwa-server/entities/vo"
	"github.com/nvwa-io/nvwa-io/nvwa-server/lang"
	"github.com/nvwa-io/nvwa-io/nvwa-server/libs/errs"
	. "github.com/nvwa-io/nvwa-io/nvwa-server/svrs"
	"github.com/polaris1119/logger"
)

type ProjectController struct {
	BaseAuthController
}

// @Title Create new project
// @router / [post]
func (t *ProjectController) Create() {
	// json decode request
	req := new(vo.ReqProject)
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

	// check whether project exist
	ok, err := DefaultProjectSvr.IsExist(req.Project.Name)
	if err != nil {
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}
	if ok {
		t.FailJson(errs.ERR_PROJECT_EXIST, req.Project.Name)
		return
	}

	// insert project
	id, err := DefaultProjectSvr.Create(t.uid(), req.Project.Name, req.Project.Description)
	if err != nil {
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}

	t.SuccJson(map[string]interface{}{
		"id": id,
	})
}

// @Title Update project
// @router /:project_id [put]
func (t *ProjectController) Update() {
	id, err := t.GetInt64(":project_id")
	if err != nil {
		t.FailJson(errs.ERR_PARAM, ":project_id")
		return
	}
	req := new(vo.ReqProject)
	err = t.ReadRequestJson(req)
	if err != nil {
		t.FailJson(errs.ERR_PARAM, err.Error())
		return
	}

	// project detail
	project, err := DefaultProjectSvr.GetById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			t.FailJson(errs.ERR_NO_RECORD)
			return
		}
	}

	// @TODO check where has permission to operate this project
	logger.Debugf("project %v", project)

	// check whether project exist
	ok, err := DefaultProjectSvr.IsExist(req.Project.Name, id)
	if err != nil {
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}
	if ok {
		t.FailJson(errs.ERR_PROJECT_EXIST, req.Project.Name)
		return
	}

	// update project
	_, err = DefaultProjectDao.UpdateById(id, dbx.Params{
		"name":        req.Project.Name,
		"description": req.Project.Description,
	})
	if err != nil {
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}

	t.SuccJson()
}

// @Title Get detail info of project
// @router /:project_id [get]
func (t *ProjectController) Detail() {
	id, err := t.GetInt64(":project_id")
	if err != nil {
		t.FailJson(errs.ERR_PARAM, ":project_id")
		return
	}

	valid := validation.Validation{}
	valid.Min(id, 1, "project_id").Message(lang.I("project.id.invalid"))
	if valid.HasErrors() {
		t.FailJson(errs.ERR_PARAM, valid.Errors[0].Error())
		return
	}

	p := new(ProjectEntity)
	err = DefaultProjectDao.GetById(id, p)
	if err != nil {
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}

	t.SuccJson(RespData{
		"project": p,
	})
}

// @Title Delete project
// @router /:project_id [delete]
func (t *ProjectController) Delete() {
	id, err := t.GetInt64(":project_id")
	if err != nil {
		t.FailJson(errs.ERR_PARAM, ":project_id")
		return
	}

	valid := validation.Validation{}
	valid.Min(id, 1, "project_id").Message(lang.I("project.id.invalid"))
	if valid.HasErrors() {
		t.FailJson(errs.ERR_PARAM, valid.Errors[0].Error())
		return
	}

	err = DefaultProjectSvr.DeleteById(id)
	if err != nil {
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}

	t.SuccJson()
}

// @Title get user related projects
// @router /
func (t *ProjectController) List() {

	// user's projects
	projectIds, err := DefaultMemberSvr.GetProjectIdsByUid(t.uid())
	if err != nil {
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}

	projects := make([]ProjectEntity, 0)
	err = DefaultProjectDao.GetAllByIdsInt64(projectIds, &projects)
	if err != nil {
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}

	// project user who created
	uids := make([]int64, 0)
	for _, v := range projects {
		uids = append(uids, v.Uid)
	}
	users, err := DefaultUserSvr.GetByUids(uids)
	if err != nil {
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}

	// format response data struct
	data := make([]map[string]interface{}, 0)
	for _, p := range projects {
		tmp := map[string]interface{}{
			"project": p,
			"user":    users[p.Uid],
		}
		data = append(data, tmp)
	}

	t.SuccJson(RespData{
		"list": data,
	})
}
