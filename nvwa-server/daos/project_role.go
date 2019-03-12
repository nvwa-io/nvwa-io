package daos

var DefaultProjectRoleDao = NewProjectRoleDao()

type ProjectRoleDao struct {
    BaseDao
}

func NewProjectRoleDao() *ProjectRoleDao {
    v := new(ProjectRoleDao)
    v.Self = v
    return v
}

func (t *ProjectRoleDao) Table() string {
    return "project_role"
}

