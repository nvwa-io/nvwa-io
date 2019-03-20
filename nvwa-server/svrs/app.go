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
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-ozzo/ozzo-dbx"
	"github.com/nvwa-io/nvwa-io/nvwa-server/clients/git"
	"github.com/nvwa-io/nvwa-io/nvwa-server/clients/jenkins"
	. "github.com/nvwa-io/nvwa-io/nvwa-server/daos"
	. "github.com/nvwa-io/nvwa-io/nvwa-server/entities"
	"github.com/nvwa-io/nvwa-io/nvwa-server/lang"
	"github.com/nvwa-io/nvwa-io/nvwa-server/libs"
	"github.com/nvwa-io/nvwa-io/nvwa-server/libs/logger"
	"os"
	"strings"
)

var DefaultAppSvr = new(AppSvr)

type AppSvr struct{}

// create app and init env and cluster for app
func (t *AppSvr) CreateAndInitEnvCluster(entity *AppEntity) (int64, error) {
	tx, err := GetDb().Begin()

	// 1. insert app
	// trans entity to dbx.Params{}
	byteV, err := json.Marshal(entity)
	if err != nil {
		return 0, err
	}

	p := dbx.Params{}
	err = json.Unmarshal(byteV, &p)
	if err != nil {
		return 0, err
	}
	p["enabled"] = ENABLED
	p["ctime"] = libs.GetNow()
	p["utime"] = p["ctime"]

	res, err := tx.Insert(DefaultAppDao.Table(), p).Execute()
	if err != nil {
		logger.Errorf("Failed to insert app, err=%s", err.Error())
		tx.Rollback()
		return 0, err
	}

	appId, err := res.LastInsertId()
	if err != nil {
		logger.Errorf("Failed to get app id, err=%s", err.Error())
		tx.Rollback()
		return 0, err
	}

	// 1.1 if use jenkins, create Jenkins project
	sys := DefaultSystemSvr.Get()
	if sys.UseJenkins {
		app := new(AppEntity)
		err = tx.Select("*").From(DefaultAppDao.Table()).Where(dbx.HashExp{"id": appId}).One(app)
		if err != nil {
			logger.Errorf("Failed to get app, err=%s", err.Error())
			tx.Rollback()
			return 0, err
		}

		appJenkinsXml, err := sys.FormatJenkinsTemplate(app)
		if err != nil {
			logger.Errorf("Failed to format jenkins template for app %s, err=%s", app.Name, err.Error())
			tx.Rollback()
			return 0, err
		}
		err = jenkins.C().Config(sys.JenkinsUrl, sys.JenkinsUser, sys.JenkinsPassword).
			Job().Create(app.Name, appJenkinsXml)
		if err != nil {
			logger.Errorf("Failed to create jenkins project for app %s, err=%s", app.Name, err.Error())
			tx.Rollback()
			return 0, err
		}
	}

	// 2. init basic environments for app
	for _, envName := range InitEnvs {
		res, err = tx.Insert(DefaultEnvDao.Table(), dbx.Params{
			"uid":             entity.Uid,
			"app_id":          appId,
			"name":            envName,
			"permit_branches": "*",
			"is_auto_deploy":  false,
			"is_need_audit":   false,
			"ctime":           libs.GetNow(),
		}).Execute()
		if err != nil {
			logger.Errorf("Failed to init app env, err=%s", err.Error())
			tx.Rollback()
			return 0, err
		}

		// 3. create default cluster for env
		envId, err := res.LastInsertId()
		if err != nil {
			logger.Errorf("Failed to get env id, err=%s", err.Error())
			tx.Rollback()
			return 0, err
		}

		_, err = tx.Insert(DefaultClusterDao.Table(), dbx.Params{
			"app_id": appId,
			"env_id": envId,
			"uid":    entity.Uid,
			"name":   lang.I("cluster.default"),
			"hosts":  "",
			"ctime":  libs.GetNow(),
		}).Execute()
		if err != nil {
			logger.Errorf("Failed to init default cluster for env, err=%s", err.Error())
			tx.Rollback()
			return 0, err
		}
	}

	// 4. commit
	err = tx.Commit()
	if err != nil {
		logger.Errorf("Failed to commit CreateAppAndInitEnvCluster, err=%s", err.Error())
		return 0, err
	}

	return appId, nil
}

func (t *AppSvr) UpdateById(id int64, params dbx.Params) error {
	_, err := DefaultAppDao.UpdateById(id, params)
	return err
}

func (t *AppSvr) CreateByEntity(entity *AppEntity) (int64, error) {
	id, err := DefaultAppDao.Create(entity)
	if err != nil {
		logger.Errorf("Failed to CreateByEntity, entity=%v, err=%s", entity, err.Error())
		return 0, err
	}

	return id, nil
}

func (t *AppSvr) IsExist(name string, excludeId ...int64) (bool, error) {
	p, err := t.GetByName(name)
	if err != nil {
		if err != sql.ErrNoRows {
			return false, err
		}

		return false, nil
	}

	if len(excludeId) > 0 && p.Id == excludeId[0] {
		return false, nil
	}

	return true, nil
}

func (t *AppSvr) GetByName(name string) (*AppEntity, error) {
	a := new(AppEntity)
	err := DefaultAppDao.GetOneByExp(dbx.HashExp{
		"name": name,
	}, a)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (t *AppSvr) ListAllByProjectId(projectId int64) ([]AppEntity, error) {
	list := make([]AppEntity, 0)
	err := DefaultAppDao.GetAllByExp(dbx.HashExp{
		"project_id": projectId,
	}, &list)
	if err != nil {
		logger.Errorf("Failed to ListAllByProjectId, err=%s", err.Error())
		return nil, err
	}

	return list, nil
}

func (t *AppSvr) GetAppIdsByProjectId(projectId int64) ([]int64, error) {
	apps, err := t.ListAllByProjectId(projectId)
	if err != nil {
		return nil, err
	}

	appIds := make([]int64, 0)
	for _, v := range apps {
		appIds = append(appIds, v.Id)
	}

	return appIds, nil
}

func (t *AppSvr) DeleteById(id int64) error {
	app, err := t.GetById(id)
	if err != nil {
		return err
	}

	tx, err := GetDb().Begin()
	if err != nil {
		return err
	}

	// 1.1 delete record
	_, err = tx.Delete(DefaultAppDao.Table(), dbx.HashExp{"id": id}).Execute()
	if err != nil {
		logger.Errorf("Failed to delete appId=%d, err=%s", id, err.Error())
		tx.Rollback()
		return err
	}

	// 1.2 delete jenkins job
	sys := DefaultSystemSvr.Get()
	if sys.UseJenkins {
		job := jenkins.C().Config(sys.JenkinsUrl, sys.JenkinsUser, sys.JenkinsPassword).Job()
		job.Delete(app.Name)

		isExist, err := job.IsExist(app.Name)
		if err != nil {
			logger.Errorf("Failed to check jenkins job [%s] whether exist, err=%s", app.Name, err.Error())
			tx.Rollback()
			return err
		}

		if isExist {
			tx.Rollback()
			return errors.New("Delete jenkins job failed.")
		}
	}

	err = tx.Commit()
	if err != nil {
		logger.Errorf("Failed to commit delete app [%s] transaction, err=%s", app.Name, err.Error())
		tx.Rollback()
		return err
	}

	return err
}

func (t *AppSvr) GetById(id int64) (*AppEntity, error) {
	entity := new(AppEntity)
	err := DefaultAppDao.GetById(id, entity)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (t *AppSvr) GetByIds(ids []int64) (map[int64]AppEntity, error) {
	apps := make([]AppEntity, 0)
	err := DefaultAppDao.GetAllByIdsInt64(ids, &apps)
	if err != nil {
		return nil, err
	}

	res := make(map[int64]AppEntity)
	for _, v := range apps {
		res[v.Id] = v
	}

	return res, nil
}

func (t *AppSvr) GetGitClientByApp(app *AppEntity) (*git.Client, error) {
	gitClient := git.C()
	system := DefaultSystemSvr.Get()
	if strings.HasPrefix(app.RepoUrl, "git") {
		// config non-thing
		// because ssh repository url uses user's ssh public key (the system user who boot nvwa-server)
		return gitClient, nil
	} else if strings.HasPrefix(app.RepoUrl, "http") {
		switch system.GitCIAuthType {
		case GIT_CI_AUTH_TYPE_BASIC:
			gitClient.BasicAuth(system.GitCIUser, system.GitCIPassword)
			return gitClient, nil
		case GIT_CI_AUTH_TYPE_TOKEN:
			gitClient.TokenAuth(system.GitCIToken)
			return gitClient, nil
		default:
			return nil, errors.New(fmt.Sprintf("Failed to recognize git ci auth type, type=%d", system.GitCIAuthType))
		}
	} else {
		return gitClient, errors.New("Can not recognize repo url type (ssh or http supported), repo url: " + app.RepoUrl)
	}
}

// clone or fetch repository
// if repository hasn't been clone, clone it from repo_url into repo_workspace.
// if it already has been cloned, fetch (update) repository.
func (t *AppSvr) CloneOrFetchRepository(app *AppEntity, checkout string, isBranch ...bool) (string, error) {
	if app == nil {
		return "", errors.New("App can't be nil.")
	}

	if app.RepoType != REPO_TYPE_GIT {
		return "", errors.New("Non-git repository type not supported.")
	}

	// 1.1 check git url and auth type
	gitClient, err := t.GetGitClientByApp(app)
	if err != nil {
		return "", err
	}

	// 2.1 check whether director exists or not.
	fi, err := os.Stat(app.LocalRepoWorkspace)
	if err != nil && !os.IsNotExist(err) {
		logger.Errorf("Failed to get file info of %s,err=%s", app.LocalRepoWorkspace, err.Error())
		return "", err
	}

	// local repository not exist, do git clone
	if os.IsNotExist(err) {
		// init local git repository
		_, err := libs.CmdExecShellDefault(fmt.Sprintf("mkdir -p %s", DefaultSystemSvr.Get().RepoRootPath))
		if err != nil {
			logger.Errorf("Failed to mkdir -p %s, err=%s", app.LocalRepoWorkspace, err.Error())
			return "", err
		}

		// clone repository
		err = gitClient.Clone(app.RepoUrl, app.LocalRepoWorkspace)
		if err != nil {
			logger.Errorf("Failed to git clone %s into %s,err=%s", app.RepoUrl, app.LocalRepoWorkspace, err.Error())
			return "", err
		}

		return "", nil
	}

	// 2.2 if path exist, check whether it's dir type
	if !fi.IsDir() {
		logger.Errorf("app=%s's repository workspaces is not dir (dir expected)", app.Name)
		return "", errors.New(fmt.Sprintf("Invalid repository workspaces of app=%s", app.Name))
	}

	// repository workspaces exist, git fetch
	err = gitClient.FetchAll(app.LocalRepoWorkspace)
	if err != nil {
		logger.Errorf("Failed to git fetch app=%s, err=%s", app.Name, err.Error())
		return "", err
	}

	// checkout to branch/tag
	if len(isBranch) > 0 && !isBranch[0] {
		err = gitClient.CheckoutTag(app.LocalRepoWorkspace, checkout)
	} else {
		err = gitClient.CheckoutBranch(app.LocalRepoWorkspace, checkout)
	}
	if err != nil {
		logger.Errorf("Failed to git checout %s, err=%s", checkout, err.Error())
		return "", err
	}

	// get branch/tag commit
	commit := ""
	checkouts := make([]string, 0)
	commits := make([]string, 0)
	if len(isBranch) > 0 && !isBranch[0] {
		checkouts, commits, err = gitClient.AllTags(app.LocalRepoWorkspace)
	} else {
		checkouts, commits, err = gitClient.AllBranches(app.LocalRepoWorkspace)
	}
	if err != nil {
		logger.Errorf("Failed to get commit of %s, err=%s", checkout, err.Error())
		return "", err
	}

	for i, v := range checkouts {
		if v == checkout {
			commit = commits[i]
		}
	}
	if commit == "" {
		err = errors.New(fmt.Sprintf("Failed to get commit, match no commit: %s", checkout))
		logger.Errorf(err.Error())
		return "", err
	}

	return commit, nil
}

func (t *AppSvr) InitTemporaryWorkspaceForBuild(repoWorkspaces, tempBuildPath string) ([]byte, error) {
	// init temporary workspace
	output, err := libs.CmdExecShellDefault(fmt.Sprintf("mkdir -p %s", tempBuildPath))
	if err != nil {
		logger.Errorf("Failed to mkdir -p %s, err=%s", tempBuildPath, err.Error())
		return output, err
	}

	// cp repo workspace to temporary build workspace
	output, err = libs.CmdExecShellDefault(fmt.Sprintf("cp -rf %s/. %s", repoWorkspaces, tempBuildPath))
	if err != nil {
		logger.Errorf("Failed to InitTemporaryWorkspaceForBuild, err=%s", err.Error())
		return output, err
	}

	return output, nil
}

func (t *AppSvr) RunBuild(buildWorkspace string, cmdBuild string, timeout int) ([]byte, error) {
	output, err := libs.CmdExecShellDefault(fmt.Sprintf("cd %s", buildWorkspace))
	if err != nil {
		logger.Errorf("Failed to cd %s, err=%s", buildWorkspace, err.Error())
		return output, err
	}

	output, err = libs.CmdExecShell(cmdBuild, timeout)
	if err != nil {
		logger.Errorf("Failed to build, err=%s", err.Error())
		return output, err
	}

	return output, nil
}

// pack version package and save to local package path
func (t *AppSvr) PackVersionPackage(buildId int64, branch, commit string, app *AppEntity) ([]byte, string, string, error) {
	// mkdir app's package root path
	cmdMkdir := fmt.Sprintf("mkdir -p %s", app.LocalPkgWorkspace)
	output, err := libs.CmdExecShellDefault(cmdMkdir)
	if err != nil {
		logger.Errorf("Failed to %s, err=%s", cmdMkdir, err.Error())
		return output, cmdMkdir, "", err
	}

	// pack package
	buildWorkspace := app.FormatTemporaryBuildPath(buildId)
	packagePath := app.GenLocalVersionPackagePath(buildId, branch, commit)
	arr := strings.Split(packagePath, "/")
	packageName := arr[len(arr)-1]
	output, cmd, err := libs.TarPackage(buildWorkspace,
		packagePath,
		strings.Split(app.Files, "\n"),
		strings.Split(app.Excludes, "\n"))
	if err != nil {
		logger.Errorf("Failed to pack version package, cmd: %s err=%s", cmd, err.Error())
		return output, cmd, packageName, err
	}

	return output, cmd, packageName, nil
}

func (t *AppSvr) GetBranchesByAppId(appId int64) ([]string, error) {
	app, err := t.GetById(appId)
	if err != nil {
		return nil, err
	}

	// clone or fetch repository
	_, err = t.CloneOrFetchRepository(app, "master")
	if err != nil {
		logger.Errorf("Failed to CloneOrFetchRepository, err=%s", err.Error())
		return nil, err
	}

	// list all branches
	gitClient, err := t.GetGitClientByApp(app)
	if err != nil {
		logger.Errorf("Failed to get git client, err=%s", err.Error())
		return nil, err
	}

	branches, _, err := gitClient.AllBranches(app.LocalRepoWorkspace)
	if err != nil {
		logger.Errorf("Failed to get branches from %s, local path=%s, err=%s", app.RepoUrl, app.LocalRepoWorkspace, err.Error())
		return nil, err
	}

	return branches, nil
}
