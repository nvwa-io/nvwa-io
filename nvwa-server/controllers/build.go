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
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/nvwa-io/nvwa-io/nvwa-server/entities/vo"
	"github.com/nvwa-io/nvwa-io/nvwa-server/lang"
	"github.com/nvwa-io/nvwa-io/nvwa-server/libs"
	"github.com/nvwa-io/nvwa-io/nvwa-server/libs/errs"
	. "github.com/nvwa-io/nvwa-io/nvwa-server/svrs"
)

type BuildController struct {
	BaseAuthController
}

// @Title Crate build
// @router / [post]
func (t *BuildController) Create() {
	// json decode request
	req := new(vo.ReqBuild)
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

	// insert project
	req.Build.Uid = t.uid()
	id, err := DefaultBuildSvr.Create(t.uid(), req.Build.AppId, req.Build.Branch)
	if err != nil {
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}

	t.SuccJson(RespData{
		"id": id,
	})
}

// @Title View build detail
// @router /:build_id [get]
func (t *BuildController) Detail() {
	id, err := t.GetInt64(":build_id")
	if err != nil {
		t.FailJson(errs.ERR_PARAM, ":build_id")
		return
	}

	valid := validation.Validation{}
	valid.Min(id, 1, "build_id").Message(lang.I("build.id.invalid"))
	if valid.HasErrors() {
		t.FailJson(errs.ERR_PARAM, valid.Errors[0].Error())
		return
	}

	entity, err := DefaultBuildSvr.GetById(id)
	if err != nil {
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}

	t.SuccJson(RespData{
		"build": entity,
	})
}

// @Title View build list
// @router /project/:project_id [get]
func (t *BuildController) List() {
	projectId, err := t.GetInt64(":project_id")
	if err != nil {
		t.FailJson(errs.ERR_PARAM, ":project_id")
		return
	}
	valid := validation.Validation{}
	valid.Min(projectId, 1, "project_id").Message(lang.I("project.id.invalid"))
	if valid.HasErrors() {
		t.FailJson(errs.ERR_PARAM, valid.Errors[0].Error())
		return
	}

	// @TODO permission validate

	appId, _ := t.GetInt64("app_id", 0)
	appIds, err := DefaultAppSvr.GetAppIdsByProjectId(projectId)
	if err != nil {
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}

	// check whether appId is belong to project
	if appId > 0 {
		if !libs.InArray(appIds, appId) {
			t.FailJson(errs.ERR_PARAM, fmt.Sprintf("app_id=%d", appId))
			return
		}
		appIds = []int64{appId}
	}

	page, _ := t.GetInt("page", 1)
	pagesize, _ := t.GetInt("pagesize", 20)

	// get page build records
	list, err := DefaultBuildSvr.ListDetailPageByAppIds(appIds, page, pagesize)
	if err != nil {
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}

	total, err := DefaultBuildSvr.CountByAppIds(appIds)
	if err != nil {
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}

	t.SuccJson(RespData{
		"list":     list,
		"total":    total,
		"page":     page,
		"pagesize": pagesize,
	})
}
