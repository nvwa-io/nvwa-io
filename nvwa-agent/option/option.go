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

type Option struct {
	// configuration fields
	Cfg *Cfg

	JenkinsEnv *JenkinsEnv

	// config file path
	ConfFile string

	// app name, same as app created in nvwa-server
	App string

	Tar struct {
		// nvwa-server build id
		BuildId int64

		// assign files or dir to pack
		// package name e.g: {project}.{buildId}.{branch}.{commitId}.{time}.tar.gz
		Files []string

		// exclude files
		Excludes []string
	}

	BuildInfo struct {
		// nvwa-server build id
		BuildId int64

		// status of build
		Status string
	}

	Notify struct {
		// nvwa-server build id
		BuildId int64

		// notify message content
		Message string

		// notify users, username register in nvwa-server
		NotifyUser []string
	}

	Push struct {
		// nvwa-server build id
		BuildId int64

		// package to push
		VersionPkg string
	}

	Pull struct {
		// package to pull
		VersionPkg string
	}
}

func NewOption() *Option {
	return &Option{}
}

// init config from config file
func (t *Option) Config() {
	// configurations from config file
	t.Cfg = initCfg(t.ConfFile)

	// jenkins env variables
	t.JenkinsEnv = GetJenkinsEnv()
}
