package daos

var DefaultDeploymentDao = NewDeploymentDao()

type DeploymentDao struct {
    BaseDao
}

func NewDeploymentDao() *DeploymentDao {
    v := new(DeploymentDao)
    v.Self = v
    return v
}

func (t *DeploymentDao) Table() string {
    return "deployment"
}
