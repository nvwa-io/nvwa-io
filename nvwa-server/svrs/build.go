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

package svrs

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-ozzo/ozzo-dbx"
	"github.com/nvwa-io/nvwa-io/nvwa-server/clients/jenkins"
	"github.com/nvwa-io/nvwa-io/nvwa-server/clients/oss-client"
	. "github.com/nvwa-io/nvwa-io/nvwa-server/daos"
	. "github.com/nvwa-io/nvwa-io/nvwa-server/entities"
	"github.com/nvwa-io/nvwa-io/nvwa-server/libs"
	"github.com/nvwa-io/nvwa-io/nvwa-server/libs/logger"
	"strings"
	"time"
)

var DefaultBuildSvr = new(BuildSvr)

type BuildSvr struct {
}

func (t *BuildSvr) Create(uid, appId int64, branch string) (int64, error) {
	id, err := DefaultBuildDao.CreateByMap(dbx.Params{
		"uid":    uid,
		"app_id": appId,
		"branch": branch,
	})
	if err != nil {
		logger.Errorf("Failed to Create, appId=%d, branch=%s, err=%s", appId, branch, err.Error())
		return 0, err
	}

	return id, nil
}

func (t *BuildSvr) GetById(id int64) (*BuildEntity, error) {
	entity := new(BuildEntity)
	err := DefaultBuildDao.GetById(id, entity)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (t *BuildSvr) ListDetailPageByAppIds(appIds []int64, page, pagesize int) ([]map[string]interface{}, error) {
	list, err := t.ListPageByAppId(appIds, page, pagesize)
	if err != nil {
		return nil, err
	}

	uids := make([]int64, 0)
	for _, v := range list {
		uids = append(uids, v.Uid)
	}

	users, err := DefaultUserSvr.GetByUids(uids)
	if err != nil {
		return nil, err
	}

	apps, err := DefaultAppSvr.GetByIds(appIds)
	if err != nil {
		return nil, err
	}

	res := make([]map[string]interface{}, 0)
	for _, v := range list {
		tmp := map[string]interface{}{
			"build": v,
			"user":  users[v.Uid],
			"app":   apps[v.AppId],
		}

		res = append(res, tmp)
	}

	return res, nil
}

func (t *BuildSvr) ListPageByAppId(appIds []int64, page, pagesize int) ([]BuildEntity, error) {
	list := make([]BuildEntity, 0)
	err := DefaultBuildDao.GetDefaultPageList(&list, dbx.HashExp{
		"app_id": libs.SliceInt64ToSliceIntf(appIds),
	}, page, pagesize, true)
	if err != nil {
		logger.Errorf("Failed to ListPageByAppId, err=%s", err.Error())
		return nil, err
	}

	return list, nil
}

func (t *BuildSvr) CountByAppIds(appIds []int64) (int, error) {
	return DefaultBuildDao.GetDefaultTotal(dbx.HashExp{
		"app_id": libs.SliceInt64ToSliceIntf(appIds),
	})
}

// get build tasks created nearly n seconds and status = BUILD_STATUS_CREATED
func (t *BuildSvr) GetWaitToBuilds(seconds int64) ([]BuildEntity, error) {
	list := make([]BuildEntity, 0)

	limitDatetime := libs.Date("Y-m-d H:i:s", time.Now().Unix()-seconds)
	err := GetDb().Select("*").From(DefaultBuildDao.Table()).
		Where(dbx.NewExp(fmt.Sprintf("ctime > '%s'", limitDatetime))).
		AndWhere(dbx.HashExp{"status": BUILD_STATUS_CREATED}).
		OrderBy("id ASC").
		All(&list)
	if err != nil {
		logger.Errorf("Failed to GetWaiToBuilds, err=%s", err.Error())
		return nil, err
	}

	return list, nil
}

// encapsulate different build type (jenkins or local)
func (t *BuildSvr) DealBuild(entity *BuildEntity) error {
	system := DefaultSystemSvr.Get()
	if system.UseJenkins {
		return t.dealJenkinsBuild(entity)
	} else {
		return t.dealLocalBuild(entity)
	}
}

// deal jenkins build
func (t *BuildSvr) dealJenkinsBuild(build *BuildEntity) error {
	// 1.1 get app configs
	app, err := DefaultAppSvr.GetById(build.AppId)
	if err != nil {
		logger.Errorf("Failed to get App, build=%s, err=%s", libs.JsonStr(build), err.Error())
		return err
	}

	// 2.1 create jenkins build job
	sys := DefaultSystemSvr.Get()
	job := jenkins.C().Config(sys.JenkinsUrl, sys.JenkinsUser, sys.JenkinsPassword).Job()

	isAppExist, err := job.IsExist(app.Name)
	if err != nil {
		logger.Errorf("Failed to get jenkins app=%s, err=%s", app.Name, err.Error())
		return err
	}

	if !isAppExist {
		info := fmt.Sprintf("Jenkins app=%s not exist.", app.Name)
		logger.Errorf(info)
		return errors.New(info)
	}

	err = job.BuildWithParameters(app.Name, map[string]string{
		"BUILD_ID":     fmt.Sprintf("%d", build.Id),
		"BUILD_BRANCH": build.Branch,
	})

	if err != nil {
		logger.Errorf("Failed to start jenkins job of %s, err=%s", app.Name, err.Error())
		return err
	}

	return nil
}

// deal local build, steps:
// 1.1 get app configs
// 2.1 fetch repository and init temporary workspace to build
// 3.1 execute build commands
// 4.1 tar version package
// 5.1 check package storage type to judge
// whether need to push package to common package repository (e.g: aliyun oss).
func (t *BuildSvr) dealLocalBuild(build *BuildEntity) error {
	logs := make([]string, 0)
	defer func() {
		// update execute log
		DefaultBuildDao.UpdateById(build.Id, dbx.Params{
			"log": strings.Join(logs, "\n"),
		})
	}()

	// 1.1 get app configs
	logs = append(logs, fmt.Sprintf("[app] app id: %d", build.AppId))
	app, err := DefaultAppSvr.GetById(build.AppId)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Errorf("App not exist, app id=%d", build.AppId)
		} else {
			logger.Errorf("Failed to GetById, err=%s", err.Error())
		}

		logs = append(logs, err.Error())
		return err
	}

	// 2.1 fetch repository and init temporary workspace to build
	checkout := ""
	if build.Branch != "" {
		checkout = build.Branch
	} else {
		checkout = build.Tag
	}
	commit, err := DefaultAppSvr.CloneOrFetchRepository(app, checkout)
	if err != nil {
		logger.Errorf("Failed to CloneOrFetchRepository, err=%s", err.Error())
		logs = append(logs, err.Error())
		return err
	}
	logs = append(logs, fmt.Sprintf("App %s repository updated ... ok", app.Name))

	// init temporary workspace for build
	output, err := DefaultAppSvr.InitTemporaryWorkspaceForBuild(app.LocalRepoWorkspace, app.FormatTemporaryBuildPath(build.Id))
	if err != nil {
		logger.Errorf("Failed to InitTemporaryWorkspaceForBuild, err=%s", err.Error())
		logs = append(logs, err.Error(), string(output))
		return err
	}
	logs = append(logs, string(output))
	defer func() { // clear temporary build workspace
		// @TODO  clear
	}()

	// 3.1 execute build commands
	logs = append(logs, app.CmdBuild)
	output, err = DefaultAppSvr.RunBuild(app.FormatTemporaryBuildPath(build.Id), app.CmdBuild, app.CmdTimeout)
	if err != nil {
		logger.Errorf("Failed to exec build commands, err=%s", err.Error())
		logs = append(logs, err.Error(), string(output))
		DefaultBuildDao.UpdateById(build.Id, dbx.Params{
			"status": BUILD_STATUS_BUILD_FAILED,
		})
		return err
	}
	logs = append(logs, string(output))
	DefaultBuildDao.UpdateById(build.Id, dbx.Params{
		"status": BUILD_STATUS_BUILD_SUCC,
	})

	// 4.1 pack version package
	output, cmd, packageName, err := DefaultAppSvr.PackVersionPackage(build.Id, checkout, commit, app)
	logs = append(logs, "[pack version package]", cmd)
	if err != nil {
		logger.Errorf("Failed to PackVersionPackage, cmd: %s, err=%s", cmd, err.Error())
		logs = append(logs, err.Error())
		DefaultBuildDao.UpdateById(build.Id, dbx.Params{
			"status": BUILD_STATUS_PACK_FAILED,
		})
		return err
	}

	// update build result
	_, err = DefaultBuildDao.UpdateById(build.Id, dbx.Params{
		"package_name": packageName,
		"status":       BUILD_STATUS_PACK_SUCC,
	})
	if err != nil {
		logger.Errorf("Failed to update package, err=%s", err.Error())
		return err
	}

	// @TODO 5.1 check package storage type to judge
	// whether need to push package to common package repository (e.g: aliyun oss).
	sys := DefaultSystemSvr.Get()
	fileKey := app.FormatKey(app.Name, packageName)
	localPkgPath := app.FormatLocalVersionPackagePath(packageName)
	storageConfig, err := sys.DecodePkgStorageConfig()
	if err != nil {
		logger.Errorf(err.Error())
		return err
	}

	switch sys.PkgStorageType {
	case PKG_STORAGE_TYPE_LOCAL:
		err = t.UpdatePushSuccAndInsertPkg(build, packageName)
		if err != nil {
			logger.Errorf("Failed to UpdatePushSuccAndInsertPkg, err=%s", err.Error())
			return err
		}

		return nil
	case PKG_STORAGE_TYPE_OSS:
		if storageConfig.Oss == nil {
			info := fmt.Sprintf("Invalid oss config, config=%s", sys.PkgStorageConfig)
			logger.Errorf(info)
			return errors.New(info)
		}

		err := oss_client.C().Config(&oss_client.OssClientConfig{
			Endpoint:     storageConfig.Oss.Endpoint,
			AccessKey:    storageConfig.Oss.AccessKey,
			AccessSecret: storageConfig.Oss.AccessSecret,
			Bucket:       storageConfig.Oss.Bucket,
		}).Put(localPkgPath, fileKey)
		if err != nil {
			logger.Errorf("Failed to put version package to oss, err=%s", err.Error())
			return err
		}

		return nil

	default:
		err = errors.New(fmt.Sprintf("Invalid package storage type: %d", DefaultSystemSvr.Get().PkgStorageType))
		logger.Errorf(err.Error())
		return err
	}
}

// update package push success and insert package record
func (t *BuildSvr) UpdatePushSuccAndInsertPkg(build *BuildEntity, packageName string) error {
	tx, err := GetDb().Begin()
	if err != nil {
		logger.Errorf("Failed to update package, err=%s", err.Error())
		return err
	}

	_, err = tx.Update(DefaultBuildDao.Table(),
		dbx.Params{"status": BUILD_STATUS_PKG_PUSH_SUCC},
		dbx.HashExp{"id": build.Id}).
		Execute()
	if err != nil {
		logger.Errorf("Failed to Update build's status to BUILD_STATUS_PKG_PUSH_SUCC, err=%s", err.Error())
		tx.Rollback()
		return err
	}

	pkg := new(PkgEntity)
	tx.Select("*").From(DefaultPkgDao.Table()).Where(dbx.HashExp{
		"build_id": build.Id,
		"name":     packageName,
		"enabled":  ENABLED,
	}).One(pkg)
	if pkg.Id == 0 { // pkg not exist, insert new package record
		_, err = tx.Insert(DefaultPkgDao.Table(), dbx.Params{
			"app_id":       build.AppId,
			"build_id":     build.Id,
			"name":         packageName,
			"branch":       build.Branch,
			"tag":          build.Tag,
			"commit_id":    build.CommitId,
			"storage_type": DefaultSystemSvr.Get().PkgStorageType,
			"ctime":        libs.GetNow(),
		}).Execute()
		if err != nil {
			logger.Errorf("Failed to insert package record, err=%s", err.Error())
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		logger.Errorf("Failed to commit update build and insert package record, err=%s", err.Error())
		return err
	}

	return nil
}

// get builds which are waiting to notify
func (t *BuildSvr) GetWaitToNotifyList() ([]BuildEntity, error) {
	list := make([]BuildEntity, 0)

	// notify when build task ends
	status := []interface{}{
		BUILD_STATUS_BUILD_FAILED,
		BUILD_STATUS_PKG_PUSH_FAILED,
		BUILD_STATUS_PKG_PUSH_SUCC,
	}
	err := GetDb().Select("*").From(DefaultBuildDao.Table()).
		Where(dbx.NewExp(fmt.Sprintf("ctime > '%s'", libs.Date("Y-m-d H:i:s", time.Now().Unix()-3600)))).
		AndWhere(dbx.HashExp{"status": status, "notified": false}).
		OrderBy("id ASC").
		All(&list)
	if err != nil {
		logger.Errorf("Failed to GetWaiToNotifyList, err=%s", err.Error())
		return nil, err
	}

	return list, nil
}
