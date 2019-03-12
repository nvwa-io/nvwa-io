package daos

import (
    "database/sql"
    "errors"
    "github.com/go-ozzo/ozzo-dbx"
    "reflect"
)

// 反射获取子类的 Table() 表名
// 要求"子类"要: 1: 继承 BaseDao, 2: 赋值子类的实例指针给 Self
func (t *BaseDao) reflectTable() (string, error) {
    v := reflect.ValueOf(t.Self)
    if !v.IsValid() {
        return "", errors.New("invalid BaseDao.self")
    }

    method := v.MethodByName("Table")
    if !method.IsValid() {
        return "", errors.New("invalid table func")
    }

    ret := method.Call(nil)
    return ret[0].String(), nil
}

func (t *BaseDao) UpdateById(id int64, params dbx.Params) (sql.Result, error) {
    table, err := t.reflectTable()
    if err != nil {
        return nil, err
    }

    res, err := t.updateById(table, id, params)
    if err != nil {
        return res, err
    }

    return res, nil
}

func (t *BaseDao) DeleteById(id int64) (sql.Result, error) {
    table, err := t.reflectTable()
    if err != nil {
        return nil, err
    }

    res, err := t.logicDelById(table, id)
    if err != nil {
        return res, err
    }

    return res, nil
}

func (t *BaseDao) GetById(id int64, objPtr interface{}) error {
    table, err := t.reflectTable()
    if err != nil {
        return err
    }

    err = t.getById(table, id, objPtr)
    if err != nil {
        return err
    }

    return nil
}

func (t *BaseDao) GetOneByExp(exp dbx.HashExp, objPtr interface{}) error {
    table, err := t.reflectTable()
    if err != nil {
        return err
    }

    err = t.getOneByHashExp(table, exp, objPtr)
    if err != nil {
        return err
    }

    return nil
}

func (t *BaseDao) GetAllByExp(exp dbx.HashExp, listPtr interface{}) error {
    table, err := t.reflectTable()
    if err != nil {
        return err
    }

    err = t.getAllByHashExp(table, exp, listPtr)
    if err != nil {
        return err
    }

    return nil
}

func (t *BaseDao) GetAllByIdsInt64(ids []int64, o interface{}, col ...string) error {
    return t.GetAllByFieldInt64("id", ids, o, col...)
}

func (t *BaseDao) GetAllByFieldInt64(field string, ids []int64, o interface{}, col ...string) error {
    table, err := t.reflectTable()
    if err != nil {
        return err
    }

    // because not support HashExp []int64
    c := make([]interface{}, 0)
    columns := make([]string, 0)
    if len(col) > 0 {
        columns = t._parseSelectStr(col[0])
    }
    for _, v := range ids {
        c = append(c, v)
    }

    return GetDb().Select(columns...).From(table).Where(dbx.HashExp{field: c, "enabled": ENABLED}).All(o)
}

func (t *BaseDao) Create(entityPtr interface{}) (int64, error) {
    table, err := t.reflectTable()
    if err != nil {
        return 0, err
    }
    res, err := t.insertEntity(table, entityPtr)
    if err != nil {
        return 0, err
    }

    return res.LastInsertId()
}

func (t *BaseDao) CreateByMap(params dbx.Params) (int64, error) {
    table, err := t.reflectTable()
    if err != nil {
        return 0, err
    }

    res, err := t.insert(table, params)
    if err != nil {
        return 0, err
    }

    return res.LastInsertId()
}

// oList: 必须是 &[]type，也就是 slice 的指针
func (t *BaseDao) GetDefaultPageList(oList interface{}, exp dbx.HashExp, page, pagesize int, isDesc bool) error {
    table, err := t.reflectTable()
    if err != nil {
        return err
    }

    offset := pagesize * (page - 1)
    if _, ok := exp["enabled"]; !ok {
        exp["enabled"] = ENABLED
    }
    query := GetDb().Select().From(table).Where(exp)
    if isDesc {
        query.OrderBy("id DESC")
    } else {
        query.OrderBy("id DESC")
    }

    return query.Offset(int64(offset)).Limit(int64(pagesize)).All(oList)
}

func (t *BaseDao) GetDefaultTotal(cond dbx.HashExp) (int, error) {
    table, err := t.reflectTable()
    if err != nil {
        return 0, err
    }

    total, err := t.count(table, cond)
    if err != nil {
        return 0, err
    }

    return total, nil
}
