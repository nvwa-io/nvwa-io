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
	. "github.com/nvwa-io/nvwa-io/nvwa-server/daos"
	. "github.com/nvwa-io/nvwa-io/nvwa-server/entities"
	"github.com/nvwa-io/nvwa-io/nvwa-server/libs/errs"
	. "github.com/nvwa-io/nvwa-io/nvwa-server/svrs"
)

// @Title get projects
// @router /admin/
func (t *ProjectController) AdminList() {
	page, _ := t.GetInt("page", 1)
	pagesize, _ := t.GetInt("pagesize", 20)

	// projects list
	projects := make([]ProjectEntity, 0)
	err := DefaultProjectDao.GetDefaultPageList(&projects, dbx.HashExp{}, page, pagesize, true)
	if err != nil {
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}
	total, err := DefaultProjectDao.GetDefaultTotal(dbx.HashExp{})
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
		"total":    total,
		"list":     data,
		"page":     page,
		"pagesize": pagesize,
	})
}
