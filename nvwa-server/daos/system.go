package daos

var DefaultSystemDao = NewSystemDao()

type SystemDao struct {
    BaseDao
}

func NewSystemDao() *SystemDao {
    v := new(SystemDao)
    v.Self = v
    return v
}

func (t *SystemDao) Table() string {
    return "system"
}
