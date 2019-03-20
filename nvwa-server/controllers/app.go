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
)

// App Api
type AppController struct {
	BaseAuthController
}

// @Title Create app
// @router / [post]
func (t *AppController) Create() {
	// json decode request
	req := new(vo.ReqApp)
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

	// check whether app exist
	ok, err := DefaultAppSvr.IsExist(req.App.Name)
	if err != nil {
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}
	if ok {
		t.FailJson(errs.ERR_APP_EXIST, req.App.Name)
		return
	}

	// create app and init env and cluster
	req.App.Uid = t.uid()
	id, err := DefaultAppSvr.CreateAndInitEnvCluster(&req.App)
	if err != nil {
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}

	t.SuccJson(RespData{
		"id": id,
	})
}

// @Title Update app
// @router /:app_id [put]
func (t *AppController) Update() {
	appId, err := t.GetInt64(":app_id")
	if err != nil {
		t.FailJson(errs.ERR_PARAM, err.Error())
		return
	}

	// json decode request
	req := new(vo.ReqApp)
	err = t.ReadRequestJson(&req)
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

	// check whether app exist
	ok, err := DefaultAppSvr.IsExist(req.App.Name, appId)
	if err != nil {
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}
	if ok {
		t.FailJson(errs.ERR_APP_EXIST, req.App.Name)
		return
	}

	err = DefaultAppSvr.UpdateById(appId, dbx.Params{
		"name":          req.App.Name,
		"description":   req.App.Description,
		"deploy_user":   req.App.DeployUser,
		"deploy_path":   req.App.DeployPath,
		"repo_url":      req.App.RepoUrl,
		"repo_username": req.App.RepoUsername,
		"repo_password": req.App.RepoPassword,
	})
	if err != nil {
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}

	t.SuccJson(RespData{})
}

// @Title update commands
// @router /:app_id/commands [put]
func (t *AppController) UpdateCmd() {
	id, err := t.GetInt64(":app_id")
	if err != nil {
		t.FailJson(errs.ERR_PARAM, ":app_id")
		return
	}

	valid := validation.Validation{}
	valid.Min(id, 1, "app_id").Message(lang.I("app.id.invalid"))
	if valid.HasErrors() {
		t.FailJson(errs.ERR_PARAM, valid.Errors[0].Error())
		return
	}

	params := dbx.Params{}
	err = t.ReadRequestJson(&params)
	if err != nil {
		t.FailJson(errs.ERR_PARAM, err.Error())
		return
	}
	_, err = DefaultAppDao.UpdateById(id, params)
	if err != nil {
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}

	t.SuccJson()
}

// @Title Get detail of app
// @router /:app_id [get]
func (t *AppController) Detail() {
	id, err := t.GetInt64(":app_id")
	if err != nil {
		t.FailJson(errs.ERR_PARAM, ":app_id")
		return
	}

	valid := validation.Validation{}
	valid.Min(id, 1, "app_id").Message(lang.I("app.id.invalid"))
	if valid.HasErrors() {
		t.FailJson(errs.ERR_PARAM, valid.Errors[0].Error())
		return
	}

	app := new(AppEntity)
	err = DefaultAppDao.GetById(id, app)
	if err != nil {
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}

	t.SuccJson(RespData{
		"app": app,
	})
}

// @Title Get app list of user's project
// @router /project/:project_id [get]
func (t *AppController) List() {
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

	// list all apps of project
	list, err := DefaultAppSvr.ListAllByProjectId(projectId)
	if err != nil {
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}

	// get app users who create the apps
	uids := make([]int64, 0)
	for _, a := range list {
		uids = append(uids, a.Uid)
	}
	users, err := DefaultUserSvr.GetByUids(uids)
	if err != nil {
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}

	// reformat response data
	res := make([]map[string]interface{}, 0)
	for _, a := range list {
		res = append(res, map[string]interface{}{
			"app":  a,
			"user": users[a.Uid],
		})
	}

	t.SuccJson(RespData{
		"list": res,
	})
}

// @Title Get app list of user's project
// @router /app-and-envs/project/:project_id [get]
func (t *AppController) ListAppAndEnv() {
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

	// list all apps of project
	list, err := DefaultAppSvr.ListAllByProjectId(projectId)
	if err != nil {
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}

	// get all env by appIds
	appIds := make([]int64, 0)
	for _, v := range list {
		appIds = append(appIds, v.Id)
	}

	envs, err := DefaultEnvSvr.GetByAppIds(appIds)
	if err != nil {
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}

	// get cluster of envs
	envIds := make([]int64, 0)
	for _, v := range envs {
		for _, ev := range v {
			envIds = append(envIds, ev.Id)
		}
	}
	envClusters, err := DefaultClusterSvr.GetByEnvIds(envIds)
	if err != nil {
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}

	// format return data
	//{
	//    "app": AppEntity,
	//    "envs": [
	//        {
	//            "env": EnvEntity,
	//            "clusters": []ClusterEntity,
	//        }
	//    ]
	//}
	res := make([]map[string]interface{}, 0)
	for _, v := range list {
		tmpEnvs := make([]map[string]interface{}, 0)
		for _, e := range envs[v.Id] {
			tmpEnvs = append(tmpEnvs, map[string]interface{}{
				"env":      e,
				"clusters": envClusters[e.Id],
			})
		}

		tmp := map[string]interface{}{
			"app":  v,
			"envs": tmpEnvs,
		}

		res = append(res, tmp)
	}

	t.SuccJson(RespData{
		"list": res,
	})
}

// @Title Get app's env and cluster
// @router /app-and-envs/app/:app_id [get]
func (t *AppController) ListAppAndEnvByAppId() {
	appId, err := t.GetInt64(":app_id")
	if err != nil {
		t.FailJson(errs.ERR_PARAM, ":app_id")
		return
	}

	valid := validation.Validation{}
	valid.Min(appId, 1, "app_id").Message(lang.I("app.id.invalid"))
	if valid.HasErrors() {
		t.FailJson(errs.ERR_PARAM, valid.Errors[0].Error())
		return
	}

	app, err := DefaultAppSvr.GetById(appId)
	if err != nil {
		if err == sql.ErrNoRows {
			t.FailJson(errs.ERR_NO_RECORD, err.Error())
			return
		}

		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}

	// get all env by appId
	envs, err := DefaultEnvSvr.ListAllByAppId(appId)
	if err != nil {
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}

	// get cluster of envs
	envIds := make([]int64, 0)
	for _, v := range envs {
		envIds = append(envIds, v.Id)
	}
	envClusters, err := DefaultClusterSvr.GetByEnvIds(envIds)
	if err != nil {
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}

	// format return data
	//   [
	//        {
	//            "env": EnvEntity,
	//            "clusters": []ClusterEntity,
	//        }
	//    ]
	tmpEnvs := make([]map[string]interface{}, 0)
	for _, e := range envs {
		tmpEnvs = append(tmpEnvs, map[string]interface{}{
			"env":      e,
			"clusters": envClusters[e.Id],
		})
	}

	t.SuccJson(RespData{
		"app":  app,
		"list": tmpEnvs,
	})
}

// @Title Delete app
// @router /:app_id [delete]
func (t *AppController) Delete() {
	id, err := t.GetInt64(":app_id")
	if err != nil {
		t.FailJson(errs.ERR_PARAM, ":app_id")
		return
	}

	valid := validation.Validation{}
	valid.Min(id, 1, "app_id").Message(lang.I("app.id.invalid"))
	if valid.HasErrors() {
		t.FailJson(errs.ERR_PARAM, valid.Errors[0].Error())
		return
	}

	err = DefaultAppSvr.DeleteById(id)
	if err != nil {
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}

	t.SuccJson()
}

// @Title View app branches
// @router /:app_id/branches [get]
func (t *AppController) Branches() {
	appId, err := t.GetInt64(":app_id")
	if err != nil {
		t.FailJson(errs.ERR_PARAM, ":app_id")
		return
	}

	branches, err := DefaultAppSvr.GetBranchesByAppId(appId)
	if err != nil {
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}

	t.SuccJson(RespData{
		"list": branches,
	})
}
