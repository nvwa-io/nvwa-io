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
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/nvwa-io/nvwa-io/nvwa-server/lang"
	"github.com/nvwa-io/nvwa-io/nvwa-server/libs/errs"
	"os"
	"strings"
)

type RespData map[string]interface{}

type BaseController struct {
	beego.Controller

	// 返回信息
	Msg map[string]interface{}
}

// action 的前钩子
func (this *BaseController) Prepare() {
	defer this.recoverBizPanic()

	// 初始化Msg的值
	this.initMsg()

	// 设置是否开启debug
	this.initDebug()

	// 签名验证
	if !this.checkSign() {
		return
	}
}

//恢复业务异常
func (this *BaseController) recoverBizPanic() {
	if err := recover(); err != nil {
		beego.Error(err)
		if v, ok := err.(string); ok {
			if strings.HasPrefix(v, "biz:") {
				this.FailJson(errs.ErrStr(v[4:]))
			}
		}
		this.FailJson(errs.ERR_SYSTEM)
	}
	return
}

// 动态开启debug
func (this *BaseController) initDebug() {
	if this.GetString("__debug") == "nvwa" {
		beego.AppConfig.Set("debug", "true")
	} else if this.GetString("__debug") == "no" {
		beego.AppConfig.Set("debug", "false")
	}
}

func (this *BaseController) initMsg() {
	this.Msg = make(map[string]interface{})
	this.Msg["code"] = errs.ERR_NONE
	this.Msg["msg"] = lang.I("success")
	this.Msg["data"] = struct{}{}
	this.Msg["node"], _ = os.Hostname()
	this.Data["json"] = this.Msg
}

// @TODO 验证签名
func (this *BaseController) checkSign() bool {
	return true
}

//设置错误信息
func (this *BaseController) FailJson(err errs.ErrStr, extMsg ...string) {
	// @TODO refactor error
	//if errs.IsErrStr(err) {
	//} else {
	//}
	this.Msg["code"] = err.GetErrno()
	if len(extMsg) > 0 {
		this.Msg["msg"] = fmt.Sprintf("%s:%s", err.GetErrmsg(), extMsg[0])
	} else {
		this.Msg["msg"] = err.GetErrmsg()
	}
	this.Data["json"] = this.Msg
	this.ServeJSON()
}

//设置返回值内容
func (this *BaseController) SuccJson(data ...interface{}) {
	if len(data) > 0 {
		this.Msg["data"] = data[0]
	}
	this.Data["json"] = this.Msg
	this.ServeJSON()
}

//返回 actionId
func (this *BaseController) GetActionId() string {
	_, actionId := this.GetControllerAndAction()
	return strings.ToLower(actionId)
}

// 获取 HTTP 请求体
func (t *BaseController) RequestBody() []byte {
	return t.Ctx.Input.RequestBody
}

// 获取 HTTP 请求体 JSON
// entity 是期望的结构的实例指针
func (t *BaseController) ReadRequestJson(entity interface{}) error {
	err := json.Unmarshal(t.RequestBody(), entity)
	return err
}
