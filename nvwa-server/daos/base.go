package daos

import (
    "github.com/go-ozzo/ozzo-dbx"
    "database/sql"
    "strings"
    "strconv"
    "errors"
    "encoding/json"
    "github.com/nvwa-io/nvwa-io/nvwa-server/libs"
)

const (
    ENABLED = 1
    DELETED = 2
)

type BaseDao struct {
    Self interface{}
}

func (t *BaseDao) insert(table string, p dbx.Params) (sql.Result, error) {
    if _, ok := p["enabled"]; !ok {
        p["enabled"] = ENABLED
    }
    p["ctime"] = libs.GetNow()
    p["utime"] = p["ctime"]
    return GetDb().Insert(table, p).Execute()
}

// v 是简单可以构成 key:value Entity
func (t *BaseDao) insertEntity(table string, v interface{}) (sql.Result, error) {
    byteV, err := json.Marshal(v)
    if err != nil {
        return nil, err
    }

    p := dbx.Params{}
    err = json.Unmarshal(byteV, &p)
    if err != nil {
        return nil, err
    }
    p["enabled"] = ENABLED
    p["ctime"] = libs.GetNow()
    p["utime"] = p["ctime"]

    return t.insert(table, p)
}

func (t *BaseDao) updateById(table string, id int64, p dbx.Params) (sql.Result, error) {
    return GetDb().Update(table, p, dbx.HashExp{"id": id}).Execute()
}

func (t *BaseDao) updateByHashExp(table string, p dbx.Params, c dbx.HashExp) (sql.Result, error) {
    return GetDb().Update(table, p, c).Execute()
}

func (t *BaseDao) logicDelById(table string, id int64) (sql.Result, error) {
    return t.updateById(table, id, dbx.Params{"enabled": DELETED})
}

func (t *BaseDao) getById(table string, id int64, o interface{}, col ...string) error {
    // select
    columns := make([]string, 0)
    if len(col) > 0 {
        columns = t._parseSelectStr(col[0])
    }
    return GetDb().Select(columns...).From(table).Where(dbx.HashExp{"id": id, "enabled": ENABLED}).One(o)
}

func (t *BaseDao) getOneByHashExp(table string, p dbx.HashExp, o interface{}, col ...string) error {
    p["enabled"] = ENABLED

    // select
    columns := make([]string, 0)
    if len(col) > 0 {
        columns = t._parseSelectStr(col[0])
    }
    return GetDb().Select(columns...).From(table).Where(p).One(o)
}

func (t *BaseDao) getAllByHashExp(table string, p dbx.HashExp, s interface{}) error {
    p["enabled"] = ENABLED
    return GetDb().Select("*").From(table).Where(p).All(s)
}

func (t *BaseDao) getAllByIdsInt64(table string, ids []int64, o interface{}, col ...string) error {
    // coz: not support HashExp []int64
    c := make([]interface{}, 0)

    columns := make([]string, 0)
    if len(col) > 0 {
        columns = t._parseSelectStr(col[0])
    }
    for _, v := range ids {
        c = append(c, v)
    }

    return GetDb().Select(columns...).From(table).Where(dbx.HashExp{"id": c}).All(o)
}

// 获取总数
func (t *BaseDao) count(table string, exp dbx.HashExp) (int, error) {
    if _, ok := exp["enabled"]; !ok {
        exp["enabled"] = ENABLED
    }

    var ret dbx.NullStringMap
    err := GetDb().Select("count(*) as total").From(table).Where(exp).One(&ret)
    if err != nil {
        return 0, err
    }

    if !ret["total"].Valid {
        return 0, errors.New("查询总数失败")
    }

    return strconv.Atoi(ret["total"].String)
}


func (t *BaseDao) _parseSelectStr(str string) []string {
    columns := strings.Split(str, ",")
    for i, v := range columns {
        columns[i] = strings.Trim(v, " ")
    }

    return columns
}

