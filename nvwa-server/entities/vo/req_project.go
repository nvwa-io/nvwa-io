package vo

import (
    "github.com/astaxie/beego/validation"
    "github.com/nvwa-io/nvwa-io/nvwa-server/entities"
    "github.com/nvwa-io/nvwa-io/nvwa-server/lang"
)

type ReqProject struct {
    Project entities.ProjectEntity `json:"project"`
}

func (t *ReqProject) Valid() error {
    valid := validation.Validation{}
    valid.Required(t.Project.Name, "name").Message(lang.I("project.name.not.empty"))
    if valid.HasErrors() {
        return valid.Errors[0]
    }

    return nil
}
