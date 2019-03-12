package logger

import (
    "github.com/astaxie/beego/logs"
    "encoding/json"
    "fmt"
    "time"
    "strings"
    "github.com/astaxie/beego"
    "sync"
    "github.com/nvwa-io/nvwa-io/nvwa-server/libs"
)

var (
    DefaultLogMan = &LogMan{
        loggers: make(map[string]*logs.BeeLogger),
    }

    // 错误级别的文件后缀
    LevelExt = map[int]string{
        logs.LevelEmergency:     "emergency",
        logs.LevelAlert:         "alert",
        logs.LevelCritical:      "critical",
        logs.LevelError:         "error",
        logs.LevelWarning:       "warning", // = LevelWarn
        logs.LevelNotice:        "notice",
        logs.LevelInformational: "info",  // = LevelInfo
        logs.LevelDebug:         "debug", // = LevelTrace
    }

    iniLevel = logs.LevelDebug
    depth    = 3
)

type (
    LogMan struct {
        mu      sync.RWMutex
        loggers map[string]*logs.BeeLogger
    }

    logManR struct {
        category      string
        logMan        *LogMan
        funcCallDepth int
    }

    logFileParams struct {
        FileName string `json:"filename"`
        Level    int    `json:"level"`
    }
)

func init() {
    var err error
    iniLevel, err = beego.AppConfig.Int("logger::LogFileLevel")
    if err != nil {
        iniLevel = logs.LevelDebug
    }

    if _, ok := LevelExt[iniLevel]; !ok {
        iniLevel = logs.LevelDebug
    }

    depth, err = beego.AppConfig.Int("logger::DefaultLogFuncCallDepth")
    if err != nil {
        depth = 3
    }
}

func (t *LogMan) R(category ...string) *logManR {
    lmr := new(logManR)
    if len(category) > 0 {
        lmr.category = category[0]
    } else {
        lmr.category = beego.AppConfig.DefaultString("appname", "app")
    }

    lmr.logMan = t
    return lmr
}

func (t *LogMan) initLogger(category string, level int, funcCallDepth int) *logs.BeeLogger {
    key := libs.Md5Str(fmt.Sprintf("%s%d", category, level))

    //{
    //t.mu.RLock()
    //defer t.mu.RUnlock()
    if _, ok := t.loggers[key]; ok {
        return t.loggers[key]
    }
    //}
    //
    //t.mu.Lock()
    //defer t.mu.Unlock()

    // 日志缓冲区间长度
    chLen, err := beego.AppConfig.Int64("logger::ChannelLen")
    if err != nil || chLen <= 0 {
        t.loggers[key] = logs.NewLogger(1000)
    } else {
        t.loggers[key] = logs.NewLogger(chLen)
    }

    // 日志调用
    enableCallDepth, err := beego.AppConfig.Bool("logger::DefaultLogFuncCallDepth")
    if err != nil {
        enableCallDepth = true
    }

    //  enable console
    enableConsole, err := beego.AppConfig.Bool("logger::EnableConsole")
    if err != nil {
        enableConsole = true
    }
    if enableConsole {
        t.loggers[key].SetLogger("console", "")
    }

    //  enable file log
    enableFile, err := beego.AppConfig.Bool("logger::EnableFile")
    if err != nil {
        enableFile = true
    }
    if enableFile {
        // 文件名格式
        curDate := time.Now().Format("20060102")
        logPath := strings.TrimRight(beego.AppConfig.String("logger::LogFilePath"), "/") + "/"
        filename := logPath + category + "." + LevelExt[level] + "." + curDate
        pb, _ := json.Marshal(&logFileParams{
            FileName: filename,
            Level:    level,
        })
        t.loggers[key].SetLogger("file", string(pb))
    }

    // log level
    t.loggers[key].SetLevel(level)
    t.loggers[key].EnableFuncCallDepth(enableCallDepth)
    t.loggers[key].SetLogFuncCallDepth(funcCallDepth)
    t.loggers[key].Async()

    return t.loggers[key]
}

func R(category ...string) *logManR {
    return DefaultLogMan.R(category...)
}

func Debugf(format string, v ...interface{}) {
    DefaultLogMan.Debugf(format, v...)
}

func Infof(format string, v ...interface{}) {
    DefaultLogMan.Infof(format, v...)
}

func Errorf(format string, v ...interface{}) {
    DefaultLogMan.Errorf(format, v...)
}

func (t *logManR) Debugf(format string, v ...interface{}) {
    if iniLevel < logs.LevelDebug {
        return
    }

    t.logMan.initLogger(t.category, logs.LevelDebug, depth).Debug(format, v...)
}

func (t *logManR) Infof(format string, v ...interface{}) {
    if iniLevel < logs.LevelInfo {
        return
    }

    t.logMan.initLogger(t.category, logs.LevelInfo, depth).Info(format, v...)
}

func (t *logManR) Errorf(format string, v ...interface{}) {
    if iniLevel < logs.LevelError {
        return
    }

    t.logMan.initLogger(t.category, logs.LevelError, depth).Error(format, v...)
}

func (t *LogMan) Debugf(format string, v ...interface{}) {
    if iniLevel < logs.LevelDebug {
        return
    }

    t.initLogger(beego.AppConfig.DefaultString("appname", "app"), logs.LevelDebug, depth+1).Debug(format, v...)
}

func (t *LogMan) Infof(format string, v ...interface{}) {
    if iniLevel < logs.LevelInfo {
        return
    }

    t.initLogger(beego.AppConfig.DefaultString("appname", "app"), logs.LevelInfo, depth+1).Info(format, v...)
}

func (t *LogMan) Errorf(format string, v ...interface{}) {
    if iniLevel < logs.LevelError {
        return
    }

    t.initLogger(beego.AppConfig.DefaultString("appname", "app"), logs.LevelError, depth+1).Error(format, v...)
}
