package entities

type ClusterEntity struct {
    BaseEntity

    AppId int64  `json:"app_id"`
    EnvId int64  `json:"env_id"`
    Uid   int64  `json:"uid"`
    Name  string `json:"name"`
    Hosts string `json:"hosts"`
}
