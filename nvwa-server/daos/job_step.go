package daos

var DefaultJobStepDao = NewJobStepDao()

type JobStepDao struct {
    BaseDao
}

func NewJobStepDao() *JobStepDao {
    v := new(JobStepDao)
    v.Self = v
    return v
}

func (t *JobStepDao) Table() string {
    return "job_step"
}
