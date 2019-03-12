package vo

import (
    "errors"
    "github.com/astaxie/beego/validation"
    "github.com/nvwa-io/nvwa-io/nvwa-server/entities"
    "github.com/nvwa-io/nvwa-io/nvwa-server/lang"
)

type ReqProjectRolePerm struct {
    Perm entities.ProjectRolePermEntity `json:"project_role_perm"`
}

func (t *ReqProjectRolePerm) Valid() error {
    valid := validation.Validation{}
    valid.Min(t.Perm.ProjectRoleId, 1, "project_role_id").Message(lang.I("project_role.id.invalid"))
    valid.Required(t.Perm.Perm, "perm").Message(lang.I("project_role_perm.perm.not.empty"))
    if valid.HasErrors() {
        return valid.Errors[0]
    }

    // check whether perm is valid key
    if _, ok := entities.PERM_LABELS[t.Perm.Perm]; !ok {
        return errors.New(lang.I("project_role_perm.perm.invalid") + ": " + t.Perm.Perm)
    }

    return nil
}
