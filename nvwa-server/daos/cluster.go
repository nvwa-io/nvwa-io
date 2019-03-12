package daos

var DefaultClusterDao = NewClusterDao()

type ClusterDao struct {
    BaseDao
}

func NewClusterDao() *ClusterDao {
    v := new(ClusterDao)
    v.Self = v
    return v
}

func (t *ClusterDao) Table() string {
    return "cluster"
}
