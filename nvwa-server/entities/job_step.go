package entities

const (
    JOB_STEP_STATUS_DEALING = 10
    JOB_STEP_STATUS_SUCC    = 20
    JOB_STEP_STATUS_FAILED  = 30

    // JOB_STEPS
    JOB_STEP_INIT_WORKSPACE         = 10
    JOB_STEP_SYNC_VERISON_PACKAGE   = 20
    JOB_STEP_UNPACK_VERISON_PACKAGE = 30
    JOB_STEP_CMD_BEFORE_DEPLOY      = 40
    JOB_STEP_DO_DEPLOY              = 50
    JOB_STEP_CMD_AFTER_DEPLOY       = 60
    JOB_STEP_CMD_HEALTH_CHECK       = 70
    JOB_STEP_CMD_ONLINE             = 80
    JOB_STEP_END_CLEAN              = 90
)

type JobStepEntity struct {
    BaseEntity

    JobId        int64  `json:"job_id"`
    AppId        int64  `json:"app_id"`
    DeploymentId int64  `json:"deployment_id"`
    Cmd          string `json:"cmd"`
    Log          string `json:"log"`
    Consume      int    `json:"consume"`
    Step         int    `json:"step"`
    Status       int    `json:"status"`
}
