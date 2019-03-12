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
	"github.com/nvwa-io/nvwa-io/nvwa-server/libs/auth"
	"github.com/nvwa-io/nvwa-io/nvwa-server/libs/errs"
	"github.com/nvwa-io/nvwa-io/nvwa-server/libs/logger"
)

// 需要登录后才能访问的 Controller 只需要组合该 BaseAuthController
type BaseAuthController struct {
	BaseController

	ucInfo *auth.UCInfo
}

func (t *BaseAuthController) Prepare() {
	t.BaseController.Prepare()

	ucInfo, err := auth.NewUserAuth().ParseUserTokenFromHeader(t.Ctx)
	if err != nil {
		logger.Errorf("Failed to ParseUserTokenFromCookie: %s", err.Error())
		t.FailJson(errs.ERR_INVALID_TOKEN, err.Error())
		t.StopRun()
	}

	logger.Debugf("UCInfo %v", ucInfo)
	t.ucInfo = ucInfo
}

// 当前用户 uid
func (t *BaseAuthController) uid() int64 {
	return t.ucInfo.Uid
}

func (t *BaseAuthController) UcInfo() *auth.UCInfo {
	return t.ucInfo
}
