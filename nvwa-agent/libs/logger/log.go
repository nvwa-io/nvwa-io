package logger

import (
    . "log"
    "runtime"
    "fmt"
    "strings"
)

const (
    LevelDebug = 1
    LevelInfo  = 2
    LevelError = 3
)

type LogFunc func(level int, format string, a ...interface{})

var log LogFunc

func SetLogger(logFunc LogFunc) {
    log = logFunc
}

func (t LogFunc) Debugf(format string, a ...interface{}) {
    t(LevelDebug, format, a...)
}

func (t LogFunc) Infof(format string, a ...interface{}) {
    t(LevelInfo, format, a...)
}

func (t LogFunc) Errorf(format string, a ...interface{}) {
    t(LevelError, format, a...)
}

func init() {
    if log == nil {
        log = func(level int, format string, a ...interface{}) {
            switch level {
            case LevelDebug:
                format = "[debug] " + format
            case LevelInfo:
                format = "[info] " + format
            case LevelError:
                format = "[error] " + format
            }

            funcName, file, line, _ := runtime.Caller(3)
            fullFuncName := runtime.FuncForPC(funcName).Name()
            arr := strings.Split(fullFuncName, "/")
            arrFile := strings.Split(file, "/")
            Printf(fmt.Sprintf("%s %s:%d ", arr[len(arr)-1], arrFile[len(arrFile)-1], line)+format, a...)
        }
    }
}

func Debugf(format string, a ...interface{}) {
    log.Debugf(format, a...)
}

func Infof(format string, a ...interface{}) {
    log.Infof(format, a...)
}

func Errorf(format string, a ...interface{}) {
    log.Errorf(format, a...)
}
