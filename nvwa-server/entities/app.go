package entities

import (
    "fmt"
    "github.com/nvwa-io/nvwa-io/nvwa-server/libs"
    "strings"
)

const (
    REPO_TYPE_GIT = "git"
    REPO_TYPE_SVN = "svn"
)

type AppEntity struct {
    BaseEntity

    Uid                 int64  `json:"uid"`
    ProjectId           int64  `json:"project_id"`
    Name                string `json:"name"`
    Description         string `json:"description"`
    DeployType          int    `json:"deploy_type"`
    AppType             int    `json:"app_type"`
    AppConfig           string `json:"app_config"`
    CmdBuild            string `json:"cmd_build"`
    RepoUrl             string `json:"repo_url"`
    RepoUsername        string `json:"repo_username"`
    RepoPassword        string `json:"repo_password"`
    PublicKey           string `json:"public_key"`
    RepoType            string `json:"repo_type"`
    LocalRepoWorkspace  string `json:"local_repo_workspace"`
    LocalBuildWorkspace string `json:"local_build_workspace"`
    LocalPkgWorkspace   string `json:"local_pkg_workspace"`
    Excludes            string `json:"excludes"`
    Files               string `json:"files"`
    CmdBeforeDeploy     string `json:"cmd_before_deploy"`
    CmdAfterDeploy      string `json:"cmd_after_deploy"`
    CmdHealthCheck      string `json:"cmd_health_check"`
    CmdOnline           string `json:"cmd_online"`
    CmdTimeout          int    `json:"cmd_timeout"`
    CmdGetPid           string `json:"cmd_get_pid"`
    DeployUser          string `json:"deploy_user"`
    DeployPath          string `json:"deploy_path"`
    RemotePkgWorkspace  string `json:"remote_pkg_workspace"`
}

// format temporary build path
func (t *AppEntity) FormatTemporaryBuildPath(buildId int64) string {
    return fmt.Sprintf(strings.TrimRight(t.LocalBuildWorkspace, "/")+"/build_id_%d", buildId)
}

// format version package name
func (t *AppEntity) FormatVersionPackageName(buildId int64, branch, commit string) string {
    // substr short commit
    shortCommit := "-"
    if len(commit) > 8 {
        shortCommit = commit[0:8]
    } else if len(commit) > 0 {
        shortCommit = commit
    }

    // appName.buildId.branch.commit.datetime.tar.gz
    return fmt.Sprintf("%s.%d.%s.%s.%s.tar.gz",
        t.Name,
        buildId,
        branch,
        shortCommit,
        libs.Date("YmdHis"))
}

// format package file name
// key for package repository, e.g: oss/cos/aws s3
func (t *AppEntity) FormatKey(app, pkgName string) string {
    return fmt.Sprintf("packages/%s/%s", app, pkgName)
}

// generate local version package path
func (t *AppEntity) GenLocalVersionPackagePath(buildId int64, branch, commit string) string {
    return strings.TrimRight(t.LocalPkgWorkspace, "/") + "/" + t.FormatVersionPackageName(buildId, branch, commit)
}

// format local version package path, e.g: /data/nvwa/demo-01/....tar.gz
func (t *AppEntity) FormatLocalVersionPackagePath(pkgName string) string {
    return strings.TrimRight(t.LocalPkgWorkspace, "/") + "/" + pkgName
}

// format remote version package path, e.g:/data/nvwa/demo-01/{pkgName}
func (t *AppEntity) FormatRemoteVersionPackagePath(pkgName string) string {
    return strings.TrimRight(t.RemotePkgWorkspace, "/") + "/" + pkgName
}

// format remote version package workspace, e.g:/data/nvwa/demo-01/{version}/
func (t *AppEntity) FormatRemoteVersionPackageWorkspace(pkgName string) string {
    return strings.TrimRight(t.RemotePkgWorkspace, "/") + "/" + strings.TrimRight(pkgName, ".tar.gz")
}
