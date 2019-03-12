package entities

type MemberEntity struct {
    BaseEntity

    ProjectId     int64 `json:"project_id"`
    Uid           int64 `json:"uid"`
    ProjectRoleId int64   `json:"project_role_id"`
}
