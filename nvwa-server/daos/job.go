package daos

var DefaultJobDao = NewJobDao()

type JobDao struct {
    BaseDao
}

func NewJobDao() *JobDao {
    v := new(JobDao)
    v.Self = v
    return v
}

func (t *JobDao) Table() string {
    return "job"
}
