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

package logic

import (
	"errors"
	"fmt"
	"github.com/nvwa-io/nvwa-io/nvwa-agent/clients/oss-client"
	"github.com/nvwa-io/nvwa-io/nvwa-agent/libs"
	"github.com/nvwa-io/nvwa-io/nvwa-agent/libs/logger"
	"github.com/nvwa-io/nvwa-io/nvwa-agent/option"
	"io/ioutil"
	"os"
	"strings"
)

func NewPkgLogic() *PkgLogic {
	return new(PkgLogic)
}

type PkgLogic struct {
	opt *option.Option
}

func (t *PkgLogic) Option(opt *option.Option) *PkgLogic {
	t.opt = opt
	return t
}

func (t *PkgLogic) Pull() error {
	fileKey := t.FormatKey(t.opt.App, t.opt.Pull.VersionPkg)
	pkgPath := t.FormatPkgPath(t.opt.App, t.opt.Pull.VersionPkg)
	appPkgDir := t.FormatPkgDir(t.opt.App)

	err := os.MkdirAll(appPkgDir, 0755)
	if err != nil {
		info := fmt.Sprintf("Failed to mkdir %s, err=%s", appPkgDir, err.Error())
		logger.Errorf(info)
		return errors.New(fmt.Sprintf("%s %s", info, err.Error()))
	}
	logger.Debugf("%s created.", appPkgDir)

	data := make([]byte, 0)
	switch t.opt.Cfg.StorageType {
	case option.STORAGE_TYPE_LOCAL:
		logger.Infof("package is stored in nvwa-server host, do nothing, package transport by ssh")
	case option.STORAGE_TYPE_OSS:
		// pull file content
		data, err = oss_client.C().Config(&oss_client.OssClientConfig{
			Endpoint:     t.opt.Cfg.Oss.Endpoint,
			AccessKey:    t.opt.Cfg.Oss.AccessKey,
			AccessSecret: t.opt.Cfg.Oss.AccessSecret,
			Bucket:       t.opt.Cfg.Oss.Bucket,
		}).Get(fileKey)
		if err != nil {
			logger.Errorf("Failed to pull version package from oss, err=%s", err.Error())
			return err
		}

		// write to local file
		err = ioutil.WriteFile(pkgPath, data, 0644)
		if err != nil {
			logger.Errorf("Failed to save package bytes to file %s, err=%s", pkgPath, err.Error())
			return err
		}

	case option.STORAGE_TYPE_COS:
		// @TODO  store in tencent cos
	case option.STORAGE_TYPE_AWS_S3:
		// @TODO  store in aws s3

	default:
		return errors.New(fmt.Sprintf("Invalid storage type: %s", t.opt.Cfg.StorageType))
	}

	logger.Debugf("%s saved.", pkgPath)
	return nil
}

func (t *PkgLogic) Push() error {
	localFile := fmt.Sprintf("%s/%s/%s", strings.TrimRight(t.opt.Cfg.JenkinsPkgWorkspace, "/"),
		strings.Trim(t.opt.App, "/"),
		strings.Trim(t.opt.Push.VersionPkg, "/"))
	fileKey := t.FormatKey(t.opt.App, t.opt.Push.VersionPkg)
	switch t.opt.Cfg.StorageType {
	case option.STORAGE_TYPE_LOCAL:
		err := NewApiLogic().Option(t.opt).UploadPackage(t.opt.Push.BuildId, t.opt.Push.VersionPkg, localFile)
		if err != nil {
			logger.Errorf("Failed to UploadPackage, err=%s", err.Error())
			return err
		}

		return nil
	case option.STORAGE_TYPE_OSS:
		err := oss_client.C().Config(&oss_client.OssClientConfig{
			Endpoint:     t.opt.Cfg.Oss.Endpoint,
			AccessKey:    t.opt.Cfg.Oss.AccessKey,
			AccessSecret: t.opt.Cfg.Oss.AccessSecret,
			Bucket:       t.opt.Cfg.Oss.Bucket,
		}).Put(localFile, fileKey)
		if err != nil {
			logger.Errorf("Failed to put version package to oss, err=%s", err.Error())
			return err
		}

		return nil
	case option.STORAGE_TYPE_COS:
		// @TODO  store in tencent cos
		return nil
	case option.STORAGE_TYPE_AWS_S3:
		// @TODO  store in aws s3

		return nil
	default:
		return errors.New(fmt.Sprintf("Invalid storage type: %s", t.opt.Cfg.StorageType))
	}
}

func (t *PkgLogic) Tar(buildId int64, files []string, excludes []string) (string, error) {
	// format jenkins packing package name and path
	pkgName := t.FormatVersionPackageName()
	appPkgWorkspace := fmt.Sprintf("%s/%s", strings.TrimRight(t.opt.Cfg.JenkinsPkgWorkspace, "/"), t.opt.App)
	pkgPath := fmt.Sprintf("%s/%s", appPkgWorkspace, pkgName)
	logger.Infof("version package path: %s", pkgPath)

	// create jenkins package workspace
	err := os.MkdirAll(appPkgWorkspace, 0644)
	if err != nil {
		info := fmt.Sprintf("Failed to mkdir -p %s, err=%s", appPkgWorkspace, err.Error())
		logger.Errorf(info)
		return "", errors.New(info)
	}

	// pack version package using linux tar command
	output, cmd, err := libs.TarPackage(t.opt.JenkinsEnv.Workspace, pkgPath, files, excludes)
	if err != nil {
		info := fmt.Sprintf("Failed to pack version package, err=%s, cmd=%s, output=%s", err.Error(), cmd, string(output))
		logger.Errorf(info)
		return "", errors.New(info)
	}

	return pkgName, nil
}

// format tar version package name
func (t *PkgLogic) FormatVersionPackageName() string {
	// substr short commit
	shortCommit := "-"
	if len(t.opt.JenkinsEnv.GitCommit) > 8 {
		shortCommit = t.opt.JenkinsEnv.GitCommit[0:8]
	} else if len(t.opt.JenkinsEnv.GitCommit) > 0 {
		shortCommit = t.opt.JenkinsEnv.GitCommit
	} else {
		shortCommit = "unknown"
	}

	// appName.buildId.branch.commit.datetime.tar.gz
	return fmt.Sprintf("%s.%d.%s.%s.%s.tar.gz",
		t.opt.App,
		t.opt.Tar.BuildId,
		t.opt.JenkinsEnv.GitBranch,
		shortCommit,
		libs.Date("YmdHis"))
}

// format package file name
func (t *PkgLogic) FormatKey(app, pkgName string) string {
	return fmt.Sprintf("packages/%s/%s", app, pkgName)
}

// format target host's package path
func (t *PkgLogic) FormatPkgPath(app, pkgName string) string {
	return fmt.Sprintf("%s/%s/%s", strings.TrimRight(t.opt.Cfg.PkgRootPath, "/"), app, pkgName)
}

// format target host's package dir
func (t *PkgLogic) FormatPkgDir(app string) string {
	return fmt.Sprintf("%s/%s", strings.TrimRight(t.opt.Cfg.PkgRootPath, "/"), app)
}
