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
	"github.com/go-ozzo/ozzo-dbx"
	"github.com/nvwa-io/nvwa-io/nvwa-server/entities/vo"
	"github.com/nvwa-io/nvwa-io/nvwa-server/lang"
	"github.com/nvwa-io/nvwa-io/nvwa-server/libs/errs"
	. "github.com/nvwa-io/nvwa-io/nvwa-server/svrs"
)

type MemberController struct {
	BaseAuthController
}

// @Title Add member to project
// @router / [post]
func (t *MemberController) Add() {
	// json decode request
	req := new(vo.ReqMember)
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

	// @TODO check whether has permission to add member

	// check whether member exist
	ok, err := DefaultMemberSvr.IsExist(req.Member.ProjectId, req.Member.Uid)
	if err != nil {
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}
	if ok {
		t.FailJson(errs.ERR_MEMBER_EXIST, fmt.Sprintf("UID: %d", req.Member.Uid))
		return
	}

	// add member
	id, err := DefaultMemberSvr.Add(req.Member.ProjectId, req.Member.Uid, req.Member.ProjectRoleId)
	if err != nil {
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}

	t.SuccJson(RespData{
		"id": id,
	})
}

// @Title Modify member role
// @router /:member_id/role [put]
func (t *MemberController) UpdateRole() {
	// @TODO check permission

	memberId, err := t.GetInt64(":member_id")
	if err != nil {
		t.FailJson(errs.ERR_PARAM, err.Error())
		return
	}

	req := new(vo.ReqMember)
	err = t.ReadRequestJson(&req)
	if err != nil {
		t.FailJson(errs.ERR_PARAM, err.Error())
		return
	}

	roles, err := DefaultProjectRoleSvr.ListMapAll()
	if err != nil {
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}

	if _, ok := roles[req.Member.ProjectRoleId]; !ok {
		t.FailJson(errs.ERR_OPERATE, lang.I("project_role.id.invalid"))
		return
	}

	err = DefaultMemberSvr.UpdateById(memberId, dbx.Params{
		"project_role_id": req.Member.ProjectRoleId,
	})
	if err != nil {
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}

	t.SuccJson()
}

// @Title Get project member list
// @router /project/:project_id [get]
func (t *MemberController) List() {
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

	// @TODO check whether has permission to get all members

	list, err := DefaultMemberSvr.GetAllDetailByProjectId(projectId)
	if err != nil {
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}

	t.SuccJson(RespData{
		"list": list,
	})
}

// @Title Remove project member
// @router /:member_id [delete]
func (t *MemberController) Delete() {
	memberId, err := t.GetInt64(":member_id")
	if err != nil {
		t.FailJson(errs.ERR_PARAM, ":member_id")
		return
	}

	valid := validation.Validation{}
	valid.Min(memberId, 1, "member_id").Message(lang.I("member.id.invalid"))
	if valid.HasErrors() {
		t.FailJson(errs.ERR_PARAM, valid.Errors[0].Error())
		return
	}

	// @TODO check whether has permission to get all members

	err = DefaultMemberSvr.DeleteById(memberId)
	if err != nil {
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}

	t.SuccJson()
}
