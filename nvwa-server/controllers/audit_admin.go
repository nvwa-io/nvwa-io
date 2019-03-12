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

// @Title get audit list waiting to deal
// @router /admin/status/:status [get]
func (t *AuditController) ListByStatus() {
	page, _ := t.GetInt("page", 1)
	pagesize, _ := t.GetInt("pagesize", 20)
	status, _ := t.GetInt(":status", entities.AUDIT_STATUS_WAITING)

	// get list
	list := make([]entities.AuditEntity, 0)
	err := daos.DefaultAuditDao.GetDefaultPageList(&list, dbx.HashExp{
		"status": status,
	}, page, pagesize, true)
	if err != nil {
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}

	fullList, err := svrs.DefaultAuditSvr.FullFillList(list, t.UcInfo().Uid)
	if err != nil {
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}

	// get total count
	total, err := daos.DefaultAuditDao.GetDefaultTotal(dbx.HashExp{"status": status})
	if err != nil {
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}

	t.SuccJson(RespData{
		"list":     fullList,
		"total":    total,
		"page":     page,
		"pagesize": pagesize,
	})
}
