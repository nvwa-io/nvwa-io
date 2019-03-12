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

package lang

var zhCN = map[string]string{
	"agent.name":         "nvwa-agent",
	"agent.long":         "查看 nvwa-agent 使用文档",
	"agent.no.operation": "[提示] 未指定具体操作, 使用请参考: ",

	"agent.cmd.config": "配置文件路径，默认：/usr/local/nvwa-agent/conf/app.ini",
	"agent.cmd.app":    "应用名称，必须与 nvwa 新建的应用名称一致",

	"agent.cmd.notify.use":         "notify --app=nvwa-demo --message=\"Hello world\" --users=username1,username2",
	"agent.cmd.notify.short":       "通知用户",
	"agent.cmd.notify.arg.message": "通知的内容",
	"agent.cmd.notify.arg.users":   "指定要通知的用户（邮箱前缀），多个用英文半角逗号分隔",

	"agent.cmd.pull.use":             "pull --app=nvwa-demo --version-pkg=nvwa-demo.11.master.abcdef.20190101010101.tar.gz",
	"agent.cmd.pull.short":           "拉包到本地",
	"agent.cmd.pull.arg.version-pkg": "版本包名称，如：nvwa-demo.11.master.abcdef.20190101010101.tar.gz",
	"agent.cmd.pull.arg.timeout":     "超时时间(单位：秒)，默认：10分钟",

	"agent.cmd.push.use":             "push --app=nvwa-demo --version-pkg=nvwa-demo.11.master.abcdef.20190101010101.tar.gz",
	"agent.cmd.push.short":           "推包到仓库",
	"agent.cmd.push.arg.build-id":    "女娲上的构建 ID",
	"agent.cmd.push.arg.version-pkg": "版本包名称，如：nvwa-demo.11.master.abcdef.20190101010101.tar.gz",
	"agent.cmd.push.arg.timeout":     "超时时间(单位：秒)，默认：10分钟",

	"agent.cmd.tar.use":          "tar --app=nvwa-demo --build-id=1 --files=file1,file2,file-*.md --excludes=file1,file2,file-*.md",
	"agent.cmd.tar.short":        "打版本包",
	"agent.cmd.tar.arg.build-id": "女娲上的构建 ID",
	"agent.cmd.tar.arg.files":    "指定要进行打包(tar.gz)的文件列表(英文半角逗号分隔, 如: -files=file1,file2,file3), 支持通配符 *, 如：-f=target/*.jar",
	"agent.cmd.tar.arg.excludes": "指定不要进行打包(tar.gz)的文件列表(英文半角逗号分隔, 如: -excludes=file1,file2,file3), 支持通配符 *, 如：-f=target/*.jar",

	"agent.cmd.build-info.use":          "build-info --app=nvwa-demo --build-id=1 --status=building",
	"agent.cmd.build-info.short":        "构建任务的信息同步",
	"agent.cmd.build-info.arg.build-id": "女娲上的构建 ID",
	"agent.cmd.build-info.arg.status":   "同步该构建的进度状态, status={building|build-success|build-failed|pack-success|pack-failed|push-success|push-failed}",
}
