package entities

const (
    AUDIT_STATUS_WAITING = 10
    AUDIT_STATUS_PASS    = DEPLOYMENT_STATUS_AUDIT_PASS
    AUDIT_STATUS_REJECT  = DEPLOYMENT_STATUS_AUDIT_REJECT
    AUDIT_STATUS_CANCELED  = DEPLOYMENT_STATUS_CANCELED
)

type AuditEntity struct {
    BaseEntity

    DeploymentId int64 `json:"deployment_id"`
    EnvId        int64 `json:"env_id"`
    ProjectId    int64 `json:"project_id"`
    AppId        int64 `json:"app_id"`
    Uid          int64 `json:"uid"`
    AuditUid     int64 `json:"audit_uid"`
    Status       int   `json:"status"`
}
