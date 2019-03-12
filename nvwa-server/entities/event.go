package entities

type Event struct {
    BaseEntity

    ProjectId int64 `json:"project_id"`
    AppId int64 `json:"app_id"`
    Uid int64 `json:"uid"`
    Content int `json:"content"`
}
