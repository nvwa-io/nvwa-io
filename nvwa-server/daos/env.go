package daos

var DefaultEnvDao = NewEnvDao()

type EnvDao struct {
    BaseDao
}

func NewEnvDao() *EnvDao {
    v := new(EnvDao)
    v.Self = v
    return v
}

func (t *EnvDao) Table() string {
    return "env"
}
