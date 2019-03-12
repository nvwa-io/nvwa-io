package libs

import (
    "reflect"
)

/**
判断obj是否存在于target中,类似于php的in_array方法
 */
func InArray(obj interface{}, target interface{}) (bool) {
    targetValue := reflect.ValueOf(target)
    switch reflect.TypeOf(target).Kind() {
    case reflect.Slice, reflect.Array:
        for i := 0; i < targetValue.Len(); i++ {
            if targetValue.Index(i).Interface() == obj {
                return true
            }
        }
    case reflect.Map:
        if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
            return true
        }
    }
    return false
}

func SliceInt64ToSliceIntf(s []int64) []interface{} {
    ret := make([]interface{}, 0)
    for _, v := range s {
        ret = append(ret, v)
    }
    return ret
}

/**
 * 过滤空字符串数组
 * eg: ["aaa", "", "bbb"]
 */
func FilterStrArr(data []string) []string {
    if len(data) == 0 {
        return data
    }

    tmpArr := make([]string, 0)
    for _, v := range data {
        if v == "" {
            continue
        }

        tmpArr = append(tmpArr, v)
    }

    return tmpArr
}

/**
 * 过滤重复的字符串元素
 * eg: ["a", "b", "a"]
 */
func UniqueStrArr(data []string) []string {
    if len(data) == 0 {
        return data
    }

    tmpArr := make([]string, 0)
    for _, v := range data {
        if InArray(v, tmpArr) {
            continue
        }

        tmpArr = append(tmpArr, v)
    }

    return tmpArr
}

/**
 * 获取在arr1但不在arr2中的元素(作用跟php的array_diff一样)
 *
 */
func DiffStrArr(arr1 []string, arr2 []string) []string {
    if len(arr2) == 0 {
        return arr1
    }

    var diffArr []string
    excludes := make(map[string]int)
    for i, s := range arr2 {
        excludes[s] = i
    }

    for _, s := range arr1 {
        if _, ok := excludes[s]; ok {
            continue
        }

        diffArr = append(diffArr, s)
    }

    return diffArr
}
