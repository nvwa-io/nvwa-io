package daos

var DefaultUserDao = NewUserDao()

type UserDao struct {
    BaseDao
}

func NewUserDao() *UserDao {
    v := new(UserDao)
    v.Self = v
    return v
}

func (t *UserDao) Table() string {
    return "user"
}
