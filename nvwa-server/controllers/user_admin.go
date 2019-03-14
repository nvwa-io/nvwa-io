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
	"github.com/nvwa-io/nvwa-io/nvwa-server/entities/vo"
	"github.com/nvwa-io/nvwa-io/nvwa-server/libs/errs"
)

// @Title get user list
// @router /admin/ [get]
func (t *UserController) List() {
	page, _ := t.GetInt("page", 1)
	pagesize, _ := t.GetInt("pagesize", 20)

	list := make([]entities.UserEntity, 0)
	err := daos.DefaultUserDao.GetDefaultPageList(&list, dbx.HashExp{}, page, pagesize, true)

	if err != nil {
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}

	total, err :=
		daos.DefaultUserDao.GetDefaultTotal(dbx.HashExp{})

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

// @Title update user role
// @router /admin/role [put]
func (t *UserController) UpdateRole() {
	u := new(vo.ReqUser)
	err := t.ReadRequestJson(u)
	if err != nil {
		t.FailJson(errs.ERR_PARAM, err.Error())
		return
	}

	if u.User.Role != entities.ROLE_ADMIN &&
		u.User.Role != entities.ROLE_USER {
		t.FailJson(errs.ERR_PARAM, err.Error())
		return
	}
}
