package vo

import (
	"errors"
	"github.com/astaxie/beego/validation"
	"github.com/nvwa-io/nvwa-io/nvwa-server/entities"
	"github.com/nvwa-io/nvwa-io/nvwa-server/lang"
)

type ReqBatchProjectRolePerm struct {
	ProjectRoleId   int64    `json:"project_role_id"`
	ProjectRoleName string   `json:"project_role_name"`
	Perms           []string `json:"perms"`
}

func (t *ReqBatchProjectRolePerm) Valid() error {
	valid := validation.Validation{}
	valid.Min(t.ProjectRoleId, 1, "project_role_id").Message(lang.I("project_role.id.invalid"))
	valid.Required(t.ProjectRoleName, "project_role_name").Message(lang.I("project_role.name.not.empty"))
	valid.MinSize(t.Perms, 1, "perms").Message(lang.I("project_role_perm.perm.not.empty"))
	if valid.HasErrors() {
		return valid.Errors[0]
	}

	// check whether perm is valid key
	for _, perm := range t.Perms {
		if _, ok := entities.PERM_LABELS[perm]; !ok {
			return errors.New(lang.I("project_role_perm.perm.invalid") + ": " + perm)
		}
	}

	return nil
}

type ReqBatchCreateProjectRolePerm struct {
	ProjectRoleName string   `json:"project_role_name"`
	Perms           []string `json:"perms"`
}

func (t *ReqBatchCreateProjectRolePerm) Valid() error {
	valid := validation.Validation{}
	valid.Required(t.ProjectRoleName, "project_role_id").Message(lang.I("project_role.name.not.empty"))
	valid.MinSize(t.Perms, 1, "perms").Message(lang.I("project_role_perm.perm.not.empty"))
	if valid.HasErrors() {
		return valid.Errors[0]
	}

	// check whether perm is valid key
	for _, perm := range t.Perms {
		if _, ok := entities.PERM_LABELS[perm]; !ok {
			return errors.New(lang.I("project_role_perm.perm.invalid") + ": " + perm)
		}
	}

	return nil
}
