package libs

import (
	"reflect"
)

// for example :
//     type A struct {
//	      str string `mapKey: "MAP_STRING"`
//     }
//     a := &A{
//       str: "val"
//     }
//     m := Struct2MapFromMapKeyTag(a)
//     m["MAP_STRING"] == "val"
func Struct2MapFromMapKeyTag(obj interface{}) map[string]string {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]string)
	for i := 0; i < t.NumField(); i++ {
		key := t.Field(i).Tag.Get("mapKey")
		data[key] = v.Field(i).String()
	}
	return data
}
