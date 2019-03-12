package daos

var DefaultProjectDao = NewProjectDao()

type ProjectDao struct {
    BaseDao
}

func NewProjectDao() *ProjectDao {
    v := new(ProjectDao)
    v.Self = v
    return v
}

func (t *ProjectDao) Table() string {
    return "project"
}
