package entities

type PkgEntity struct {
    BaseEntity

    AppId       int64  `json:"app_id"`
    BuildId     int64  `json:"build_id"`
    Name        string `json:"name"`
    Branch      string `json:"branch"`
    Tag         string `json:"tag"`
    CommitId    string `json:"commit_id"`
    StorageType string `json:"storage_type"`
}
