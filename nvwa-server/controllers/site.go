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
	"github.com/nvwa-io/nvwa-io/nvwa-server/lang"
	"github.com/nvwa-io/nvwa-io/nvwa-server/libs/auth"
	"github.com/nvwa-io/nvwa-io/nvwa-server/libs/errs"
	. "github.com/nvwa-io/nvwa-io/nvwa-server/svrs"
)

type SiteController struct {
	BaseController
}

// @Title TODO Site index
// @router / [get]
func (t *SiteController) Index() {

}

// @Title User login
// @router /login [post]
func (t *SiteController) Login() {
	// json decode request
	var req map[string]string
	err := t.ReadRequestJson(&req)
	if err != nil {
		t.FailJson(errs.ERR_PARAM, err.Error())
		return
	}

	// validate request params
	valid := validation.Validation{}
	valid.AlphaDash(req["username"], "username").Message(lang.I("user.username.invalid"))
	valid.Required(req["password"], "password").Message(lang.I("user.password.not.empty"))
	valid.MinSize(req["password"], 6, "password").Message(lang.I("user.password.length.error"))
	if valid.HasErrors() {
		t.FailJson(errs.ERR_PARAM, valid.Errors[0].Error())
		return
	}

	// get user
	user, err := DefaultUserSvr.GetByUsernamePassword(req["username"], req["password"])
	if err != nil {
		if err != sql.ErrNoRows {
			t.FailJson(errs.ERR_OPERATE, err.Error())
			return
		}
	}

	if user == nil {
		t.FailJson(errs.ERR_LOGIN, req["username"])
		return
	}

	// response token info
	token, expire, err := auth.NewUserAuth().UCInfoEncryptByUid(user.Id)
	if err != nil {
		t.FailJson(errs.ERR_OPERATE, user.Username)
		return
	}
	t.SuccJson(RespData{
		"token":  token,
		"expire": expire,
	})
}

// @Title User register
// @router /register [post]
func (t *SiteController) Register() {
	// json decode request
	var req map[string]string
	err := t.ReadRequestJson(&req)
	if err != nil {
		t.FailJson(errs.ERR_PARAM, err.Error())
		return
	}

	// validate request params
	valid := validation.Validation{}
	valid.AlphaDash(req["username"], "username").Message(lang.I("user.username.invalid"))
	valid.Email(req["email"], "email").Message(lang.I("user.email.format.error"))
	valid.Required(req["password"], "password").Message(lang.I("user.password.not.empty"))
	valid.MinSize(req["password"], 6, "password").Message(lang.I("user.password.length.error"))
	if valid.HasErrors() {
		t.FailJson(errs.ERR_PARAM, valid.Errors[0].Error())
		return
	}

	// check whether username exist
	user, err := DefaultUserSvr.GetByUsername(req["username"])
	if err != nil {
		if err != sql.ErrNoRows {
			t.FailJson(errs.ERR_OPERATE, err.Error())
			return
		}
	}

	if user != nil {
		t.FailJson(errs.ERR_USER_EXIST, user.Username)
		return
	}

	// insert user
	uid, err := DefaultUserSvr.Register(req["username"], req["email"], req["password"])
	if err != nil {
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}

	// response token info
	token, expire, err := auth.NewUserAuth().UCInfoEncryptByUid(uid)
	if err != nil {
		t.FailJson(errs.ERR_OPERATE, user.Username)
		return
	}
	t.SuccJson(RespData{
		"token":  token,
		"expire": expire,
	})
}

// @Title Get user by token
// @router /token/user [get]
func (t *SiteController) GetUserByToken() {
	ucInfo, err := auth.NewUserAuth().ParseUserTokenFromHeader(t.Ctx)
	if err != nil {
		t.FailJson(errs.ERR_INVALID_TOKEN, err.Error())
		return
	}

	user, err := DefaultUserSvr.GetById(ucInfo.Uid)
	if err != nil {
		if err == sql.ErrNoRows {
			t.FailJson(errs.ERR_NO_RECORD, err.Error())
			return
		}

		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}

	t.SuccJson(RespData{
		"user": user,
	})
}
