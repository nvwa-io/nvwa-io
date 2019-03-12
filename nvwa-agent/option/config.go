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

package option

import (
	"encoding/json"
	"github.com/BurntSushi/toml"
	"github.com/nvwa-io/nvwa-io/nvwa-agent/libs/logger"
	"log"
)

const (
	// store in host
	STORAGE_TYPE_LOCAL = "local"
	// aliyun oss
	STORAGE_TYPE_OSS = "oss"
	// tencent cloud cos
	STORAGE_TYPE_COS = "cos"
	// aws s3
	STORAGE_TYPE_AWS_S3 = "aws-s3"
)

type Cfg struct {
	Version string
	Locale  string

	JenkinsPkgWorkspace string
	StorageType         string
	PkgRootPath         string

	Oss struct {
		Endpoint     string
		AccessKey    string
		AccessSecret string
		Bucket       string
	}

	// interaction api with nvwa server
	NvwaApi struct {
		BuildInfo     string
		Notify        string
		UploadPackage string
	}
}

func initCfg(confFile string) *Cfg {
	logger.Debugf("Init config file: %s", confFile)
	cfg := new(Cfg)
	if _, err := toml.DecodeFile(confFile, cfg); err != nil {
		log.Fatalln("Failed to init config file: ", err.Error())
	}

	o, _ := json.Marshal(*cfg)
	logger.Debugf("config=%v", string(o))
	return cfg
}
