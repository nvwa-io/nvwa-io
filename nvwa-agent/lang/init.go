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

var info localeInfo

type localeInfo map[string]map[string]string

func init() {
	info = make(localeInfo)
	info["zh_CN"] = zhCN
	info["en_US"] = enUs
}

func locale() map[string]string {
	//locale := beego.AppConfig.DefaultString("locale", "zh_CN")
	// @TODO Read from config file
	locale := "zh_CN"
	if _, ok := info[locale]; ok {
		return info[locale]
	}

	for _, v := range info {
		return v
	}

	return make(map[string]string)
}

func I(key string) string {
	tmp := locale()
	if v, ok := tmp[key]; ok {
		return v
	}

	return key
}
