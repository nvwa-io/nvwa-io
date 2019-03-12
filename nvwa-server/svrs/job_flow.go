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
	"errors"
	"fmt"
	"github.com/go-ozzo/ozzo-dbx"
	"github.com/nvwa-io/nvwa-io/nvwa-server/clients/ansible"
	. "github.com/nvwa-io/nvwa-io/nvwa-server/entities"
	"github.com/nvwa-io/nvwa-io/nvwa-server/libs/logger"
	"strings"
	"time"
)

func NewJobFlow(deploy *DeploymentEntity, job *JobEntity, app *AppEntity) *JobFlow {
	return &JobFlow{
		Deploy: deploy,
		Job:    job,
		App:    app,
	}
}

// Deploy job flow
type JobFlow struct {
	Deploy *DeploymentEntity
	Job    *JobEntity
	App    *AppEntity
}

type ConsumeTimer struct {
	jobFlow *JobFlow
	step    int

	beginTime int64
	params    dbx.Params
}

func NewConsumeTimer(jobFlow *JobFlow, step int) *ConsumeTimer {
	ct := new(ConsumeTimer)
	ct.jobFlow = jobFlow
	ct.params = dbx.Params{
		"cmd":    "",
		"log":    "",
		"status": 0,
	}
	ct.step = step

	return ct
}
func (t *ConsumeTimer) begin() *ConsumeTimer {
	t.beginTime = time.Now().UnixNano() / 1e6
	_, err := DefaultJobStepSvr.CreateByJob(t.jobFlow.Job, t.step)
	if err != nil {
		logger.Errorf("Failed to create step, err=%s", err.Error())
	}

	return t
}

func (t *ConsumeTimer) update(cmd, log string, status int) error {
	t.params["cmd"] = cmd
	t.params["log"] = log
	t.params["status"] = status
	t.params["consume"] = time.Now().UnixNano()/1e6 - t.beginTime
	return DefaultJobStepSvr.UpdateByJobIdStep(t.jobFlow.Job.Id, t.step, t.params)
}

// do deploy works
func (t *JobFlow) Do() error {
	defer func() {
		err := t.clean()
		if err != nil {
			logger.Errorf("Failed to clean JobFlow works, err=%s", err.Error())
		}
	}()

	// 1. init deploy workspace for remote hosts
	err := t.initRemoteWorkspace()
	if err != nil {
		return err
	}

	// 2. push/(or pull package from common storage repository) version package to remote hosts
	err = t.syncVersionPkg()
	if err != nil {
		return err
	}

	// 3. unpack remote version package
	err = t.unpackRemoteVersionPkg()
	if err != nil {
		return err
	}

	// 4. execute before deploy commands
	err = t.execCmdBeforeDeploy()
	if err != nil {
		return err
	}

	// 5. do remote deploy
	err = t.doRemoteDeploy()
	if err != nil {
		return err
	}

	// 6. execute after deploy commands
	err = t.execCmdAfterDeploy()
	if err != nil {
		return err
	}

	// 7. execute healthy check commands
	err = t.execCmdHealthCheck()
	if err != nil {
		return err
	}

	// 8. execute online commands
	err = t.execCmdOnline()
	if err != nil {
		return err
	}

	return nil
}

func (t *JobFlow) initRemoteWorkspace() error {
	// init step stat log
	ct := NewConsumeTimer(t, JOB_STEP_INIT_WORKSPACE).begin()

	// format command and execute
	cmd := fmt.Sprintf("mkdir -p %s", t.App.FormatRemoteVersionPackageWorkspace(t.Deploy.Pkg))
	output, cmd, err := ansible.C().ExecShell(t.App.DeployUser,
		cmd,
		strings.Split(t.Job.DeployHosts, ","),
		t.App.CmdTimeout)
	if err != nil {
		logger.Errorf("Failed to init remote workspaces, err=%s, output=%s", err.Error(), string(output))
		ct.update(cmd, fmt.Sprintf("%s \n %s", string(output), err.Error()), JOB_STEP_STATUS_FAILED)
		return err
	}
	ct.update(cmd, string(output), JOB_STEP_STATUS_SUCC)
	return nil
}

func (t *JobFlow) syncVersionPkg() error {
	// init step stat log
	ct := NewConsumeTimer(t, JOB_STEP_SYNC_VERISON_PACKAGE).begin()

	// sync version package according to package storage type
	var output []byte
	var cmd string
	var err error
	switch DefaultSystemSvr.Get().PkgStorageType {
	case PKG_STORAGE_TYPE_LOCAL:
		srcPkg := t.App.FormatLocalVersionPackagePath(t.Deploy.Pkg)
		destPkg := t.App.FormatRemoteVersionPackagePath(t.Deploy.Pkg)
		output, cmd, err = ansible.C().CopyFile(t.App.DeployUser,
			srcPkg,
			destPkg,
			strings.Split(t.Job.DeployHosts, ","),
			t.App.CmdTimeout)
		if err != nil {
			logger.Errorf("Failed to init remote workspaces,cmd=%s err=%s, output=%s", cmd, err.Error(), string(output))
			ct.update(cmd, fmt.Sprintf("%s \n %s", string(output), err.Error()), JOB_STEP_STATUS_FAILED)
			return err
		}
	case PKG_STORAGE_TYPE_OSS:
		// @TODO OSS version package storage
	default:
		return errors.New(fmt.Sprintf("Failed to recognize PkgStorageType: %d", DefaultSystemSvr.Get().PkgStorageType))
	}

	ct.update(cmd, string(output), JOB_STEP_STATUS_SUCC)
	return nil
}

func (t *JobFlow) unpackRemoteVersionPkg() error {
	// init step stat log
	ct := NewConsumeTimer(t, JOB_STEP_UNPACK_VERISON_PACKAGE).begin()

	// unpack remote version package
	versionPkgWorkspace := t.App.FormatRemoteVersionPackageWorkspace(t.Deploy.Pkg)
	targetPkg := t.App.FormatRemoteVersionPackagePath(t.Deploy.Pkg)
	cmd := fmt.Sprintf("cd %s && tar --no-same-owner -pm -C %s -xz -f %s",
		versionPkgWorkspace,
		versionPkgWorkspace,
		targetPkg,
	)

	output, cmd, err := ansible.C().ExecShell(t.App.DeployUser, cmd, strings.Split(t.Job.DeployHosts, ","), t.App.CmdTimeout)
	if err != nil {
		logger.Errorf("Failed to init remote workspaces, cmd=%s err=%s, output=%s", cmd, err.Error(), string(output))
		ct.update(cmd, fmt.Sprintf("%s \n %s", string(output), err.Error()), JOB_STEP_STATUS_FAILED)
		return err
	}

	ct.update(cmd, string(output), JOB_STEP_STATUS_SUCC)
	return nil
}

func (t *JobFlow) execCmdBeforeDeploy() error {
	// init step stat log
	ct := NewConsumeTimer(t, JOB_STEP_CMD_BEFORE_DEPLOY).begin()

	// @TODO 替换环境变量

	versionPkgWorkspace := t.App.FormatRemoteVersionPackageWorkspace(t.Deploy.Pkg)
	cmd := fmt.Sprintf("cd %s; %s", versionPkgWorkspace, t.App.CmdBeforeDeploy)
	output, cmd, err := ansible.C().ExecShell(t.App.DeployUser, cmd, strings.Split(t.Job.DeployHosts, ","), t.App.CmdTimeout)
	if err != nil {
		logger.Errorf("Failed to execute commands before deploy, cmd=%s err=%s, output=%s", cmd, err.Error(), string(output))
		ct.update(cmd, fmt.Sprintf("%s \n %s", string(output), err.Error()), JOB_STEP_STATUS_FAILED)
		return err
	}

	ct.update(cmd, string(output), JOB_STEP_STATUS_SUCC)
	return nil
}

func (t *JobFlow) doRemoteDeploy() error {
	// execute soft link
	err := t.softLink()
	if err != nil {
		return err
	}

	return nil
}

func (t *JobFlow) softLink() error {
	// init step stat log
	ct := NewConsumeTimer(t, JOB_STEP_DO_DEPLOY).begin()

	// soft link package version to deploy path
	tmpDeployPath := fmt.Sprintf("%s.%s", strings.TrimRight(t.App.DeployPath, "/"), "tmp")
	versionPkgWorkspace := t.App.FormatRemoteVersionPackageWorkspace(t.Deploy.Pkg)
	parentDeployPath := t.App.DeployPath[0:strings.LastIndex(t.App.DeployPath, "/")]

	cmds := []string{
		fmt.Sprintf("mkdir -p %s", strings.TrimRight(parentDeployPath, "/")),
		fmt.Sprintf("rm -rf %s", tmpDeployPath),
		fmt.Sprintf("ln -sfn %s %s", versionPkgWorkspace, tmpDeployPath),
		fmt.Sprintf("chown -h %s %s", t.App.DeployUser, tmpDeployPath),
		fmt.Sprintf("mv -fT %s %s", tmpDeployPath, strings.TrimRight(t.App.DeployPath, "/")),
	}
	cmd := strings.Join(cmds, " && ")
	output, cmd, err := ansible.C().ExecShell(t.App.DeployUser, cmd, strings.Split(t.Job.DeployHosts, ","), t.App.CmdTimeout)
	if err != nil {
		logger.Errorf("Failed to execute deploy soft link, cmd=%s err=%s, output=%s", cmd, err.Error(), string(output))
		ct.update(cmd, fmt.Sprintf("%s \n %s", string(output), err.Error()), JOB_STEP_STATUS_FAILED)
		return err
	}

	ct.update(cmd, string(output), JOB_STEP_STATUS_SUCC)
	return err
}

func (t *JobFlow) execCmdAfterDeploy() error {
	// init step stat log
	ct := NewConsumeTimer(t, JOB_STEP_CMD_AFTER_DEPLOY).begin()

	// @TODO 替换环境变量

	versionPkgWorkspace := t.App.FormatRemoteVersionPackageWorkspace(t.Deploy.Pkg)
	cmd := fmt.Sprintf("cd %s; %s", versionPkgWorkspace, t.App.CmdBeforeDeploy)
	output, cmd, err := ansible.C().ExecShell(t.App.DeployUser, cmd, strings.Split(t.Job.DeployHosts, ","), t.App.CmdTimeout)
	if err != nil {
		logger.Errorf("Failed to execute commands after deploy, cmd=%s err=%s, output=%s", cmd, err.Error(), string(output))
		ct.update(cmd, fmt.Sprintf("%s \n %s", string(output), err.Error()), JOB_STEP_STATUS_FAILED)
		return err
	}

	ct.update(cmd, string(output), JOB_STEP_STATUS_SUCC)
	return nil
}

func (t *JobFlow) execCmdHealthCheck() error {
	// init step stat log
	ct := NewConsumeTimer(t, JOB_STEP_CMD_HEALTH_CHECK).begin()

	// @TODO 替换环境变量

	versionPkgWorkspace := t.App.FormatRemoteVersionPackageWorkspace(t.Deploy.Pkg)
	cmd := fmt.Sprintf("cd %s; %s", versionPkgWorkspace, t.App.CmdHealthCheck)
	output, cmd, err := ansible.C().ExecShell(t.App.DeployUser, cmd, strings.Split(t.Job.DeployHosts, ","), t.App.CmdTimeout)
	if err != nil {
		logger.Errorf("Failed to execute healthy check commands, cmd=%s err=%s, output=%s", cmd, err.Error(), string(output))
		ct.update(cmd, fmt.Sprintf("%s \n %s", string(output), err.Error()), JOB_STEP_STATUS_FAILED)
		return err
	}

	ct.update(cmd, string(output), JOB_STEP_STATUS_SUCC)
	return nil
}

func (t *JobFlow) execCmdOnline() error {
	// init step stat log
	ct := NewConsumeTimer(t, JOB_STEP_CMD_ONLINE).begin()

	// @TODO 替换环境变量

	versionPkgWorkspace := t.App.FormatRemoteVersionPackageWorkspace(t.Deploy.Pkg)
	cmd := fmt.Sprintf("cd %s; %s", versionPkgWorkspace, t.App.CmdOnline)
	output, cmd, err := ansible.C().ExecShell(t.App.DeployUser, cmd, strings.Split(t.Job.DeployHosts, ","), t.App.CmdTimeout)
	if err != nil {
		logger.Errorf("Failed to execute online commands, cmd=%s err=%s, output=%s", cmd, err.Error(), string(output))
		ct.update(cmd, fmt.Sprintf("%s \n %s", string(output), err.Error()), JOB_STEP_STATUS_FAILED)
		return err
	}

	ct.update(cmd, string(output), JOB_STEP_STATUS_SUCC)
	return nil
}

// @TODO clean
func (t *JobFlow) clean() error {
	// init step stat log
	ct := NewConsumeTimer(t, JOB_STEP_END_CLEAN).begin()

	//$this->cleanRemoteWork();
	//$this->cleanLocalWork();

	cmd := "clean"
	var output = []byte{}
	//ct.update(cmd, fmt.Sprintf("%s \n %s", string(output), err.Error()), JOB_STEP_STATUS_FAILED)
	ct.update(cmd, string(output), JOB_STEP_STATUS_SUCC)
	return nil
}
