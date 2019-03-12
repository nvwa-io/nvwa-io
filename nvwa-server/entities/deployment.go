package entities

import "strings"

const (
    FILE_DEPLOY_MODE_FULL  = 1
    FILE_DEPLOY_MODE_DELTA = 2

    DEPLOYMENT_STATUS_CREATED      = 10
    DEPLOYMENT_STATUS_NO_AUDIT     = 20
    DEPLOYMENT_STATUS_WAIT_AUDIT   = 30
    DEPLOYMENT_STATUS_AUDIT_PASS   = 40
    DEPLOYMENT_STATUS_AUDIT_REJECT = 50
    DEPLOYMENT_STATUS_CANCELED     = 60
    DEPLOYMENT_STATUS_DEALING      = 70
    DEPLOYMENT_STATUS_SUCC         = 80
    DEPLOYMENT_STATUS_FAILED       = 90
)

type DeploymentEntity struct {
    BaseEntity

    ProjectId      int64  `json:"project_id"`
    AppId          int64  `json:"app_id"`
    Uid            int64  `json:"uid"`
    EnvId          int64  `json:"env_id"`
    PkgId          int64  `json:"pkg_id"`
    Pkg            string `json:"pkg"`
    ClusterIds     string `json:"cluster_ids"`
    ClusterHosts   string `json:"cluster_hosts"`
    IsAutoDeploy   bool   `json:"is_auto_deploy"`
    IsAllCluster   bool   `json:"is_all_cluster"`
    LinkId         string `json:"link_id"`
    Branch         string `json:"branch"`
    CommitId       string `json:"commit_id"`
    FileList       string `json:"file_list"`
    FileDeployMode int    `json:"file_deploy_mode"`
    LatestLinkId   string `json:"latest_link_id"`
    IsNeedAudit    bool   `json:"is_need_audit"`
    Status         int    `json:"status"`
}

func (t *DeploymentEntity) GetVersion() string {
    return strings.TrimRight(t.Pkg, ".tar.gz")
}
