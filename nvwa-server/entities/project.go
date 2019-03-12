package entities

type ProjectEntity struct {
    BaseEntity

    Name        string `json:"name"`
    Description string `json:"description"`
    Uid         int64  `json:"uid"`
}
