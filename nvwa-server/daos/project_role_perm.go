package daos

var DefaultProjectRolePermDao = NewProjectRolePermDao()

type ProjectRolePermDao struct {
    BaseDao
}

func NewProjectRolePermDao() *ProjectRolePermDao {
    v := new(ProjectRolePermDao)
    v.Self = v
    return v
}

func (t *ProjectRolePermDao) Table() string {
    return "project_role_perm"
}
