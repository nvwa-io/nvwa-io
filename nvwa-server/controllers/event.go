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

type EventController struct {
	BaseAuthController
}

// @Title 新建事件
// @router / [post]
func (t *EventController) Create() {

}

// @Title 修改事件
// @router /:event_id [put]
func (t *EventController) Update() {

}

// @Title 查看事件详情
// @router /:event_id [put]
func (t *EventController) Detail() {

}

// @Title 查看事件列表
// @router /user/:uid [get]
func (t *EventController) List() {

}

// @Title 删除事件
// @router /:event_id [delete]
func (t *EventController) Delete() {

}
