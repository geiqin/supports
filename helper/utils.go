package helper

import (
	"errors"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"reflect"
)

//截取字符串 start 起点下标 end 终点下标(不包括)
func Substr(str string, start int, end int) (string,error) {
	rs := []rune(str)
	length := len(rs)

	if start < 0 || start > length {
		return "", errors.New("start is wrong:" + string(start))
	}

	if end < 0 || end > length {
		return "", errors.New("end is wrong:" + string(start))
	}
	return string(rs[start:end]), nil
}

//Map类型转换为Struct
func MapToStruct(fromMap interface{},toStruct interface{}) interface{} {
	mapstructure.Decode(fromMap,toStruct)
	return toStruct
}

//判断字符是否在数组中
func InArray(s [] string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}


//把任意数字类型转换为int64
func ToInt64(value interface{}) (d int64, err error) {
	val := reflect.ValueOf(value)
	switch value.(type) {
	case int, int8, int16, int32, int64:
		d = val.Int()
	case uint, uint8, uint16, uint32, uint64:
		d = int64(val.Uint())
	default:
		err = fmt.Errorf("ToInt64 need numeric not `%T`", value)
	}
	return
}