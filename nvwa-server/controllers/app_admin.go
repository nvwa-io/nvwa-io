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
	"github.com/go-ozzo/ozzo-dbx"
	"github.com/nvwa-io/nvwa-io/nvwa-server/daos"
	"github.com/nvwa-io/nvwa-io/nvwa-server/entities"
	"github.com/nvwa-io/nvwa-io/nvwa-server/libs/errs"
	"github.com/nvwa-io/nvwa-io/nvwa-server/svrs"
)

// @Title Get app list
// @router /admin/ [get]
func (t *AppController) AdminList() {
	page, _ := t.GetInt("page", 1)
	pagesize, _ := t.GetInt("pagesize", 20)

	// list all apps of project
	list := make([]entities.AppEntity, 0)
	err := daos.DefaultAppDao.GetDefaultPageList(&list, dbx.HashExp{}, page, pagesize, true)
	if err != nil {
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}
	total, err := daos.DefaultAppDao.GetDefaultTotal(dbx.HashExp{})
	if err != nil {
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}

	// get app users who create the apps
	uids := make([]int64, 0)
	for _, a := range list {
		uids = append(uids, a.Uid)
	}
	users, err := svrs.DefaultUserSvr.GetByUids(uids)
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
		"list":     res,
		"total":    total,
		"page":     page,
		"pagesize": pagesize,
	})
}
