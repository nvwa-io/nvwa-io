package lang

import "github.com/astaxie/beego"

var info localeInfo

type localeInfo map[string]map[string]string

func init() {
    info = make(localeInfo)
    info["zh_CN"] = zhCN
    info["en_US"] = enUs
}

func locale() map[string]string {
    locale := beego.AppConfig.DefaultString("locale", "zh_CN")
    if _, ok := info[locale]; ok {
        return info[locale]
    }

    for _, v := range info {
        return v
    }

    return make(map[string]string)
}

func I(key string) string {
    tmp := locale()
    if v, ok := tmp[key]; ok {
        return v
    }

    return key
}
