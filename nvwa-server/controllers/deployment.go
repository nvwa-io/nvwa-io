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

type DeploymentController struct {
	BaseAuthController
}

// @Title Create deployment
// @router / [post]
func (t *DeploymentController) Create() {
	// json decode request
	req := new(vo.ReqDeployment)
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

	// create app and init env and cluster
	req.Deployment.Uid = t.uid()
	id, err := DefaultDeploymentSvr.CreateAndInitJob(&req.Deployment)
	if err != nil {
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}

	t.SuccJson(RespData{
		"id": id,
	})
}

// @Title @TODO audit deployment
// @router /audit/:deployment_id [put]
func (t *DeploymentController) Audit() {

}

// @Title Get page of deployments
// @router /project/:project_id [get]
func (t *DeploymentController) List() {
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

	// get page deployment records
	list, err := DefaultDeploymentSvr.ListPageDetailByAppIds(appIds, page, pagesize)
	if err != nil {
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}

	total, err := DefaultDeploymentSvr.CountByAppIds(appIds)
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
