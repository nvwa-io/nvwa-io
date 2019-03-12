package daos

var DefaultBuildDao = NewBuildDao()

type BuildDao struct {
    BaseDao
}

func NewBuildDao() *BuildDao {
    v := new(BuildDao)
    v.Self = v
    return v
}

func (t *BuildDao) Table() string {
    return "build"
}
