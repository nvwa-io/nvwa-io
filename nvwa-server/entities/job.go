package entities

const (
    JOB_STATUS_CREATED = 10
    JOB_STATUS_READY   = 20
    JOB_STATUS_DEALING = 30
    JOB_STATUS_SUCC    = 40
    JOB_STATUS_FAILED  = 50
)

type JobEntity struct {
    BaseEntity

    DeploymentId int64  `json:"deployment_id"`
    AppId        int64  `json:"app_id"`
    EnvId        int64  `json:"env_id"`
    ClusterId    int64  `json:"cluster_id"`
    AllHosts     string `json:"all_hosts"`
    DeployHosts  string `json:"deploy_hosts"`
    ExcludeHosts string `json:"exclude_hosts"`
    Status       int    `json:"status"`
}
