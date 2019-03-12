package libs

import (
    "bytes"
    "crypto/md5"
    "encoding/hex"
    "fmt"
    "k8s.io/kubernetes/staging/src/k8s.io/apimachinery/pkg/util/json"
    "strings"
    "text/template"
    "time"
)

/**
 * Md5加密字符串
 */
func Md5Str(s string) string {
    h := md5.New()
    h.Write([]byte(s))
    return hex.EncodeToString(h.Sum(nil))
}

// parse golang template
func ParseTemplate(templateStr string, data interface{}) (string, error) {
    name := fmt.Sprintf("%d", time.Now().Nanosecond())
    temp, err := template.New(name).Parse(templateStr)
    if err != nil {
        return "", err
    }

    content := make([]byte, 0)
    buf := bytes.NewBuffer(content)
    err = temp.Execute(buf, data)
    if err != nil {
        return "", err
    }

    return buf.String(), nil
}

/**
 * 截取字符串
 */
func Substr(str string, start, length int) string {
    rs := []rune(str)
    rl := len(rs)
    end := 0

    if start < 0 {
        start = rl - 1 + start
    }
    end = start + length

    if start > end {
        start, end = end, start
    }

    if start < 0 {
        start = 0
    }
    if start > rl {
        start = rl
    }
    if end < 0 {
        end = 0
    }
    if end > rl {
        end = rl
    }

    return string(rs[start:end])
}

// 字符串：下划线转驼峰
func TransToCamel(s string) string {
    data := make([]byte, 0, len(s))
    j := false
    k := false
    num := len(s) - 1
    for i := 0; i <= num; i++ {
        d := s[i]
        if k == false && d >= 'A' && d <= 'Z' {
            k = true
        }
        if d >= 'a' && d <= 'z' && (j || k == false) {
            d = d - 32
            j = false
            k = true
        }
        if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
            j = true
            continue
        }
        data = append(data, d)
    }
    return string(data[:])
}

// snake string, XxYy to xx_yy , XxYY to xx_yy
func TransToUnderline(s string) string {
    data := make([]byte, 0, len(s)*2)
    j := false
    num := len(s)
    for i := 0; i < num; i++ {
        d := s[i]
        if i > 0 && d >= 'A' && d <= 'Z' && j {
            data = append(data, '_')
        }
        if d != '_' {
            j = true
        }
        data = append(data, d)
    }
    return strings.ToLower(string(data[:]))
}

func JsonStr(v interface{}) string {
    b, _ := json.Marshal(v)
    return string(b)
}
