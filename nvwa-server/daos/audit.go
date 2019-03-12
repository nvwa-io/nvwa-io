package daos

var DefaultAuditDao = NewAuditDao()

type AuditDao struct {
    BaseDao
}

func NewAuditDao() *AuditDao{
    v := new(AuditDao)
    v.Self = v
    return v
}

func (t *AuditDao) Table() string {
    return "audit"
}
