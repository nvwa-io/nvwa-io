package daos

var DefaultEventDao = NewEventDao()

type EventDao struct {
    BaseDao
}

func NewEventDao() *EventDao {
    v := new(EventDao)
    v.Self = v
    return v
}

func (t *EventDao) Table() string {
    return "event"
}
