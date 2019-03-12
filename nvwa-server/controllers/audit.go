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
	. "github.com/nvwa-io/nvwa-io/nvwa-server/entities"
	"github.com/nvwa-io/nvwa-io/nvwa-server/libs/errs"
	. "github.com/nvwa-io/nvwa-io/nvwa-server/svrs"
)

type AuditController struct {
	BaseAuthController
}

// @Title get audit list waiting to deal
// @router /wait [get]
func (t *AuditController) ListWait() {
	projectIds, err := DefaultMemberSvr.GetProjectIdsByUid(t.uid())
	if err != nil {
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}

	page, _ := t.GetInt("page", 1)
	pagesize, _ := t.GetInt("pagesize", 20)
	status := []int{AUDIT_STATUS_WAITING}

	list, err := DefaultAuditSvr.GetByProjectIdsAndStatus(projectIds, status, page, pagesize)
	if err != nil {
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}
	fullList, err := DefaultAuditSvr.FullFillList(list, t.uid())
	if err != nil {
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}

	total, err := DefaultAuditSvr.GetCountByProjectIdsAndStatus(projectIds, status)
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

// @Title get audit list waiting num
// @router /wait-num [get]
func (t *AuditController) GetWaitCount() {
	projectIds, err := DefaultMemberSvr.GetProjectIdsByUid(t.uid())
	if err != nil {
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}
	status := []int{AUDIT_STATUS_WAITING}
	total, err := DefaultAuditSvr.GetCountByProjectIdsAndStatus(projectIds, status)
	if err != nil {
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}
	t.SuccJson(RespData{
		"num": total,
	})
}

// @Title get audit list of mine
// @router /mine [get]
func (t *AuditController) ListMine() {
	page, _ := t.GetInt("page", 1)
	pagesize, _ := t.GetInt("pagesize", 20)
	list, err := DefaultAuditSvr.GetByUid(t.uid(), page, pagesize)
	if err != nil {
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}
	fullList, err := DefaultAuditSvr.FullFillList(list, t.uid())
	if err != nil {
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}

	total, err := DefaultAuditSvr.GetCountByUid(t.uid())
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

// @Title get audit list of audited
// @router /audited [get]
func (t *AuditController) ListAudited() {
	projectIds, err := DefaultMemberSvr.GetProjectIdsByUid(t.uid())
	if err != nil {
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}

	page, _ := t.GetInt("page", 1)
	pagesize, _ := t.GetInt("pagesize", 20)

	status := []int{AUDIT_STATUS_PASS, AUDIT_STATUS_REJECT}

	list, err := DefaultAuditSvr.GetByProjectIdsAndStatus(projectIds, status, page, pagesize)
	if err != nil {
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}
	fullList, err := DefaultAuditSvr.FullFillList(list, t.uid())
	if err != nil {
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}

	total, err := DefaultAuditSvr.GetCountByProjectIdsAndStatus(projectIds, status)
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

// @Title pass audit
// @router /:audit_id/pass [put]
func (t *AuditController) Pass() {
	t.updateStatus(AUDIT_STATUS_PASS)
}

// @Title reject audit
// @router /:audit_id/reject [put]
func (t *AuditController) Reject() {
	t.updateStatus(AUDIT_STATUS_REJECT)
}

// @Title cancel audit
// @router /:audit_id/cancel [put]
func (t *AuditController) Cancel() {
	t.updateStatus(AUDIT_STATUS_CANCELED)
}

func (t *AuditController) updateStatus(status int) {
	auditId, err := t.GetInt64(":audit_id")
	if err != nil {
		t.FailJson(errs.ERR_PARAM, err.Error())
		return
	}

	// @TODO check permission

	err = DefaultAuditSvr.UpdateStatusById(auditId, status, t.uid())
	if err != nil {
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}

	t.SuccJson()
}
