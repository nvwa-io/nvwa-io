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
	"os"
	"strings"
)

//  all jenkins environment variables
//  BUILD_URL=http://192.168.20.235:8080/job/svr-infra-demo/14/
//	POM_GROUPID=com.xf
//	HUDSON_SERVER_COOKIE=50d02f8ad16a60e9
//	SHELL=/bin/bash
//	BUILD_TAG=jenkins-svr-infra-demo-14
//	POM_DISPLAYNAME=service-user
//	GIT_PREVIOUS_COMMIT=a6e0006b3ce077511e8ac4b12960413fd3ed9c41
//	WORKSPACE=/var/lib/jenkins/workspace/svr-infra-demo
//	JOB_URL=http://192.168.20.235:8080/job/svr-infra-demo/
//	RUN_CHANGES_DISPLAY_URL=http://192.168.20.235:8080/job/svr-infra-demo/14/display/redirect?page=changes
//	USER=jenkins
//	POM_ARTIFACTID=service-user
//	NLSPATH=/usr/dt/lib/nls/msg/%L/%N.cat
//	JENKINS_HOME=/var/lib/jenkins
//	GIT_COMMIT=a6e0006b3ce077511e8ac4b12960413fd3ed9c41
//	MAVEN_HOME=/usr/local/apache-maven
//	PATH=/usr/local/apache-maven/bin:/sbin:/usr/sbin:/bin:/usr/bin
//	RUN_DISPLAY_URL=http://192.168.20.235:8080/job/svr-infra-demo/14/display/redirect
//	_=/bin/env
//	PWD=/var/lib/jenkins/workspace/svr-infra-demo
//	HUDSON_URL=http://192.168.20.235:8080/
//	LANG=zh_CN.UTF-8
//	POM_VERSION=0.0.1-SNAPSHOT
//	JOB_NAME=svr-infra-demo
//	BUILD_DISPLAY_NAME=#14
//	XFILESEARCHPATH=/usr/dt/app-defaults/%L/Dt
//	JENKINS_URL=http://192.168.20.235:8080/
//	BUILD_ID=14
//	JOB_BASE_NAME=svr-infra-demo
//	GIT_PREVIOUS_SUCCESSFUL_COMMIT=a6e0006b3ce077511e8ac4b12960413fd3ed9c41
//	POM_PACKAGING=jar
//	SHLVL=3
//	HOME=/var/lib/jenkins
//	M2_HOME=/usr/local/apache-maven
//	GIT_BRANCH=origin/master
//	EXECUTOR_NUMBER=0
//	JENKINS_SERVER_COOKIE=50d02f8ad16a60e9
//	GIT_URL=http://git.xf.io/infra/svr-infra-demo.git
//	NODE_LABELS=master
//	LOGNAME=jenkins
//	HUDSON_HOME=/var/lib/jenkins
//	NODE_NAME=master
//	JOB_DISPLAY_URL=http://192.168.20.235:8080/job/svr-infra-demo/display/redirect
//	BUILD_NUMBER=14
//	HUDSON_COOKIE=2ca6f3d0-885f-4b58-9847-567b55fdd940

// jenkins environment variables
type JenkinsEnv struct {
	// the app's workspace path
	Workspace string

	// jenkins home
	JenkinsHome string

	// jenkins build number
	BuildNumber string

	// job name, e.g: app-demo-01
	JobName string

	// jenkins-${JOB_NAME}-${BUILD_NUMBER}
	BuildTag string

	// git branch
	GitBranch string

	// git Commit ID
	GitCommit string

	// SVN_REVISION
	SvnRevision string

	// user who launch this build
	BuildUserId string
}

func GetJenkinsEnv() *JenkinsEnv {
	env := new(JenkinsEnv)

	env.Workspace = os.Getenv("WORKSPACE")
	env.JenkinsHome = os.Getenv("JENKINS_HOME")
	env.BuildNumber = os.Getenv("BUILD_NUMBER")
	env.JobName = os.Getenv("JOB_NAME")
	env.BuildTag = os.Getenv("BUILD_TAG")
	env.GitBranch = os.Getenv("GIT_BRANCH")
	env.GitCommit = os.Getenv("GIT_COMMIT")
	env.SvnRevision = os.Getenv("SVN_REVISION")
	env.BuildUserId = os.Getenv("BUILD_USER_ID")
	arr := strings.Split(env.GitBranch, "/")
	if len(arr) > 0 {
		env.GitBranch = arr[len(arr)-1]
	}

	if env.GitCommit != "" {
		env.GitCommit = env.GitCommit[0:8]
	}

	return env
}
