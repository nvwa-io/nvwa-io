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

package apis

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-ozzo/ozzo-dbx"
	"github.com/nvwa-io/nvwa-io/nvwa-server/controllers"
	"github.com/nvwa-io/nvwa-io/nvwa-server/daos"
	"github.com/nvwa-io/nvwa-io/nvwa-server/libs/errs"
	"github.com/nvwa-io/nvwa-io/nvwa-server/libs/logger"
	"github.com/nvwa-io/nvwa-io/nvwa-server/svrs"
	"strings"
)

type BuildApi struct {
	BaseApi
}

// e.g: /update-info?build_id=3&hostname=hostname01
// @Title Update build
// @router /update-info [put]
func (t *BuildApi) UpdateInfo() {
	buildId, _ := t.GetInt64("build_id", 0)
	if buildId <= 0 {
		t.FailJson(errs.ERR_PARAM, fmt.Sprintf("invalid build_id: %d", buildId))
		return
	}
	params := dbx.Params{}
	err := json.Unmarshal(t.RequestBody(), &params)
	if err != nil {
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}

	_, err = daos.DefaultBuildDao.UpdateById(buildId, params)
	if err != nil {
		t.FailJson(errs.ERR_OPERATE, "failed to upload build info, "+err.Error())
		return
	}

	t.SuccJson()
}

// e.g: /upload-package?build_id=3&hostname=hostname01
// @Title upload-package
// @router /upload-package [post]
func (t *BuildApi) UploadPackage() {
	// 1.1 get parameters
	formFileName := "pkgFile"
	buildId, _ := t.GetInt64("build_id", 0)
	if buildId <= 0 {
		t.FailJson(errs.ERR_PARAM, fmt.Sprintf("invalid build_id: %d", buildId))
		return
	}
	f, h, err := t.GetFile(formFileName)
	if err != nil {
		t.FailJson(errs.ERR_PARAM, err.Error())
		return
	}
	defer f.Close()

	logger.Debugf("package file: %s", h.Filename)
	if !strings.HasSuffix(h.Filename, ".tar.gz") {
		t.FailJson(errs.ERR_PARAM, "invalid file, only ext .tar.gz allowed.")
		return
	}

	// 2.1 get build record
	build, err := svrs.DefaultBuildSvr.GetById(buildId)
	if err != nil {
		if err == sql.ErrNoRows {
			t.FailJson(errs.ERR_NO_RECORD, err.Error())
			return
		}
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}

	// 2.2 get app record
	app, err := svrs.DefaultAppSvr.GetById(build.AppId)
	if err != nil {
		if err == sql.ErrNoRows {
			t.FailJson(errs.ERR_NO_RECORD, err.Error())
			return
		}
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}

	// 3.1 save file to local package path
	localPkgPath := app.FormatLocalVersionPackagePath(h.Filename)
	err = t.SaveToFile(formFileName, localPkgPath)
	if err != nil {
		info := fmt.Sprintf("Failed to save file to %s, err=%s", localPkgPath, err.Error())
		logger.Errorf(info)
		t.FailJson(errs.ERR_OPERATE, info)
		return
	}

	// 4.1 update build record's package_name
	err = svrs.DefaultBuildSvr.UpdatePushSuccAndInsertPkg(build, h.Filename)
	if err != nil {
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}

	t.SuccJson(controllers.RespData{
		"pkg_path": localPkgPath,
	})
}
