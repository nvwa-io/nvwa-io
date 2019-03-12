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
	"strings"
)

type PkgController struct {
	BaseAuthController
}

// @Title get latest {limit} package records
// @router /env/:env_id [get]
func (t *PkgController) List() {
	envId, err := t.GetInt64(":env_id")
	if err != nil {
		t.FailJson(errs.ERR_PARAM, err.Error())
		return
	}
	limit, _ := t.GetInt("limit", 20)

	env, err := DefaultEnvSvr.GetById(envId)
	if err != nil {
		if err == sql.ErrNoRows {
			t.FailJson(errs.ERR_NO_RECORD, err.Error())
			return
		}

		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}

	branch := make([]string, 0)
	if env.PermitBranches != "*" {
		branch = strings.Split(env.PermitBranches, ",")
	}
	storageType := DefaultSystemSvr.Get().PkgStorageType
	list, err := DefaultPkgSvr.GetLatestListByEnvId(env.AppId, branch, storageType, limit)
	if err != nil {
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}

	t.SuccJson(RespData{
		"list": list,
	})
}
