package daos

var DefaultPkgDao = NewPkgDao()

type PkgDao struct {
    BaseDao
}

func NewPkgDao() *PkgDao {
    v := new(PkgDao)
    v.Self = v
    return v
}

func (t *PkgDao) Table() string {
    return "pkg"
}
