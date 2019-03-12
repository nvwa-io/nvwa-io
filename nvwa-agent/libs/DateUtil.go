package libs

import (
    "time"
    "strings"
    "errors"
)

const (
    TIME_FORMAT          = "2006-01-02 15:04:05"
    TIME_FORMAT_YYYYMMDD = "20060102"
)

var TimeFormatMap = map[string]string{
    "Y": "2006",
    "m": "01",
    "d": "02",
    "H": "15",
    "i": "04",
    "s": "05",
}

/**
 * 将日期(字符串)转为时间戳
 */
func DateToTimestamp(dateStr string) int64 {
    loc, _ := time.LoadLocation("Local")
    theTime, _ := time.ParseInLocation(TIME_FORMAT, dateStr, loc)
    return theTime.Unix()
}

func CurDatetime() string {
    return time.Now().Format(TIME_FORMAT)
}

/**
 * 类似php的date()函数
 *
 * @param
 * @return
 *
 */
func Date(format string, timestamp ...int64) string {
    newFormat := format
    for k, v := range TimeFormatMap {
        newFormat = strings.Replace(newFormat, k, v, 1)
    }
    var tm time.Time
    if len(timestamp) > 0 {
        tm = time.Unix(timestamp[0], 0)
    } else {
        tm = time.Now()
    }
    return tm.Format(newFormat)
}

func GetNow() string {
    return time.Now().Format(TIME_FORMAT)
}

// UTC 时间转本地时间
func Utc2LocalDate(d string) (string, error) {
    t, err := time.Parse(time.RFC3339, d)
    if err != nil {
        return "", errors.New("格式化时间失败: " + err.Error())
    }

    local := t.In(time.Local)
    return local.Format(TIME_FORMAT), nil
}
