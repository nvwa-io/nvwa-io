package vo

import (
    "github.com/astaxie/beego/validation"
    "github.com/nvwa-io/nvwa-io/nvwa-server/entities"
    "github.com/nvwa-io/nvwa-io/nvwa-server/lang"
)

type ReqBuild struct {
    Build entities.BuildEntity `json:"build"`
}

func (t *ReqBuild) Valid() error {
    valid := validation.Validation{}
    valid.Min(t.Build.AppId, 0, "app_id").Message(lang.I("app.id.invalid"))
    valid.Required(t.Build.Branch, "branch").Message(lang.I("branch.not.empty"))
    if valid.HasErrors() {
        return valid.Errors[0]
    }

    return nil
}
