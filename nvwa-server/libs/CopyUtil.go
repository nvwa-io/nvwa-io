package libs

import (
    "reflect"
    "errors"
    "encoding/json"
)

// @param dst 要被赋值的结构体，必须是指针
// @param src 数据源，必须是一个结构体（非指针）
func Copy(dst interface{}, src interface{}) (err error) {
    dstValue := reflect.ValueOf(dst)
    if dstValue.Kind() != reflect.Ptr {
        err = errors.New("dst isn't a pointer to struct")
        return
    }
    dstElem := dstValue.Elem()
    if dstElem.Kind() != reflect.Struct {
        err = errors.New("pointer doesn't point to struct")
        return
    }

    srcValue := reflect.ValueOf(src)
    srcType := reflect.TypeOf(src)
    if srcType.Kind() != reflect.Struct {
        err = errors.New("src isn't struct")
        return
    }

    for i := 0; i < srcType.NumField(); i++ {
        sf := srcType.Field(i)
        sv := srcValue.FieldByName(sf.Name)
        // make sure the value which in dst is valid and can set
        if dv := dstElem.FieldByName(sf.Name); dv.IsValid() && dv.CanSet() {
            dv.Set(sv)
        }
    }

    return
}

// 结构体转为 map
// obj 必须是结构体实例，不可以是指针
// toUnderLine, 是否需要将驼峰转为 下划线
func Struct2Map(obj interface{}, toUnderline ...bool) map[string]interface{} {
    t := reflect.TypeOf(obj)
    v := reflect.ValueOf(obj)

    var data = make(map[string]interface{})
    for i := 0; i < t.NumField(); i++ {
        if len(toUnderline) > 0 && toUnderline[0] {
            data[TransToUnderline(t.Field(i).Name)] = v.Field(i).Interface()
        } else {
            data[t.Field(i).Name] = v.Field(i).Interface()
        }
    }
    return data
}

// mapObj - map 实例
// structObj - 结构体实例指针
func Map2Struct(mapObj interface{}, structObj interface{}) error {
    data, err := json.Marshal(mapObj)
    if err != nil {
        return err
    }
    return json.Unmarshal(data, structObj)
}
