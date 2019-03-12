package vo

import (
    "github.com/astaxie/beego/validation"
    "github.com/nvwa-io/nvwa-io/nvwa-server/entities"
    "github.com/nvwa-io/nvwa-io/nvwa-server/lang"
)

type ReqEnv struct {
    Env entities.EnvEntity `json:"env"`
}

func (t *ReqEnv) Valid() error {
    valid := validation.Validation{}
    valid.Min(t.Env.AppId, 1, "app_id").Message(lang.I("app.id.invalid"))
    valid.Required(t.Env.Name, "name").Message(lang.I("env.name.not.empty"))
    valid.Required(t.Env.PermitBranches, "permit_branches").Message(lang.I("env.permit_branches.not.empty"))
    valid.Required(t.Env.IsAutoDeploy, "is_auto_deploy").Message(lang.I("env.is_auto_deploy.not.empty"))
    valid.Required(t.Env.IsNeedAudit, "is_need_audit").Message(lang.I("env.is_need_audit.not.empty"))
    if valid.HasErrors() {
        return valid.Errors[0]
    }

    return nil
}
