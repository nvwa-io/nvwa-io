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
	"github.com/nvwa-io/nvwa-io/nvwa-server/daos"
	"github.com/nvwa-io/nvwa-io/nvwa-server/entities/vo"
	"github.com/nvwa-io/nvwa-io/nvwa-server/lang"
	"github.com/nvwa-io/nvwa-io/nvwa-server/libs/errs"
	. "github.com/nvwa-io/nvwa-io/nvwa-server/svrs"
)

type EnvController struct {
	BaseAuthController
}

// @Title Create app environment
// @router / [post]
func (t *EnvController) Create() {
	// json decode request
	req := new(vo.ReqEnv)
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

	// check whether env exist
	ok, err := DefaultEnvSvr.IsExist(req.Env.AppId, req.Env.Name)
	if err != nil {
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}
	if ok {
		t.FailJson(errs.ERR_ENV_EXIST, req.Env.Name)
		return
	}

	// insert env
	req.Env.Uid = t.uid()
	id, err := DefaultEnvSvr.CreateAndInitDefaultCluster(&req.Env)
	if err != nil {
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}

	t.SuccJson(RespData{
		"id": id,
	})
}

// @Title Update app environment
// @router /:env_id [put]
func (t *EnvController) Update() {
	id, err := t.GetInt64(":env_id")
	if err != nil {
		t.FailJson(errs.ERR_PARAM, ":env_id")
		return
	}
	req := new(vo.ReqEnv)
	err = t.ReadRequestJson(req)
	if err != nil {
		t.FailJson(errs.ERR_PARAM, err.Error())
		return
	}

	// env detail
	env, err := DefaultEnvSvr.GetById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			t.FailJson(errs.ERR_NO_RECORD)
			return
		}
	}

	// check whether env exist
	ok, err := DefaultEnvSvr.IsExist(env.AppId, req.Env.Name, id)
	if err != nil {
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}
	if ok {
		t.FailJson(errs.ERR_ENV_EXIST, req.Env.Name)
		return
	}

	// update env
	_, err = daos.DefaultEnvDao.UpdateById(id, dbx.Params{
		"name":            req.Env.Name,
		"permit_branches": req.Env.PermitBranches,
		"is_auto_deploy":  req.Env.IsAutoDeploy,
		"is_need_audit":   req.Env.IsNeedAudit,
		"cmd_env":         req.Env.CmdEnv,
	})

	if err != nil {
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}

	t.SuccJson()
}

// @Title Get app environment detail
// @router /:env_id [get]
func (t *EnvController) Detail() {
	id, err := t.GetInt64(":env_id")
	if err != nil {
		t.FailJson(errs.ERR_PARAM, ":env_id")
		return
	}

	valid := validation.Validation{}
	valid.Min(id, 1, "env_id").Message(lang.I("env.id.invalid"))
	if valid.HasErrors() {
		t.FailJson(errs.ERR_PARAM, valid.Errors[0].Error())
		return
	}

	entity, err := DefaultEnvSvr.GetById(id)
	if err != nil {
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}

	t.SuccJson(RespData{
		"env": entity,
	})
}

// @Title Get app environment list
// @router /app/:app_id [get]
func (t *EnvController) List() {
	appId, err := t.GetInt64(":app_id")
	if err != nil {
		t.FailJson(errs.ERR_PARAM, ":env_id")
		return
	}

	valid := validation.Validation{}
	valid.Min(appId, 1, "app_id").Message(lang.I("app.id.invalid"))
	if valid.HasErrors() {
		t.FailJson(errs.ERR_PARAM, valid.Errors[0].Error())
		return
	}

	list, err := DefaultEnvSvr.ListAllByAppId(appId)
	if err != nil {
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}

	t.SuccJson(RespData{
		"list": list,
	})
}

// @Title Delete app environment
// @router /:env_id [delete]
func (t *EnvController) Delete() {
	id, err := t.GetInt64(":env_id")
	if err != nil {
		t.FailJson(errs.ERR_PARAM, ":env_id")
		return
	}

	_, err = daos.DefaultEnvDao.DeleteById(id)
	if err != nil {
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}

	t.SuccJson()
}
