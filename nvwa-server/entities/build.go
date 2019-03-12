package entities

const (
    BUILD_STATUS_CREATED         = 10
    BUILD_STATUS_BUIDING         = 20
    BUILD_STATUS_BUILD_SUCC      = 30
    BUILD_STATUS_BUILD_FAILED    = 40
    BUILD_STATUS_PACK_SUCC       = 50
    BUILD_STATUS_PACK_FAILED     = 60
    BUILD_STATUS_PKG_PUSH_SUCC   = 70
    BUILD_STATUS_PKG_PUSH_FAILED = 80
)

type BuildEntity struct {
    BaseEntity

    AppId           int64  `json:"app_id"`
    Uid             int64  `json:"uid"`
    Branch          string `json:"branch"`
    Tag             string `json:"tag"`
    CommitId        string `json:"commit_id"`
    PackageName     string `json:"package_name"`
    JenkinsBuildNum int    `json:"jenkins_build_num"`
    Log             string `json:"log"`
    Notified        bool   `json:"notified"`
    Status          int    `json:"status"`
}
