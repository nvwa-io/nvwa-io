package daos

var DefaultMemberDao = NewMemberDao()

type MemberDao struct {
    BaseDao
}

func NewMemberDao() *MemberDao {
    v := new(MemberDao)
    v.Self = v
    return v
}

func (t *MemberDao) Table() string {
    return "member"
}
