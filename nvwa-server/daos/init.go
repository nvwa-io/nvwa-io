package daos

import (
    "fmt"
    "github.com/nvwa-io/nvwa-io/nvwa-server/libs/logger"
    "github.com/astaxie/beego"
    "github.com/go-ozzo/ozzo-dbx"
    _ "github.com/go-sql-driver/mysql"
)

var (
    db *dbx.DB
)

// 初始化数据库配置
func initConfig() {
    dbconf, _ := beego.AppConfig.GetSection("db")

    var err error
    charset := fmt.Sprintf("charset=%s", dbconf["charset"])
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s",
        dbconf["username"],
        dbconf["password"],
        dbconf["host"],
        dbconf["port"],
        dbconf["database"],
        charset)

    db, err = dbx.Open("mysql", dsn)
    if err != nil {
        panic(err.Error())
    }

    db.LogFunc = func(format string, sql ...interface{}) {
        // operations, e.g: log to file
        logger.R("sql").Debugf(format, sql...)
    }

    db.DB().SetMaxOpenConns(1000)
    db.DB().SetMaxIdleConns(10)
}

// 获取 DB
func GetDb() *dbx.DB {
    if db == nil {
        initConfig()
    }
    return db
}
