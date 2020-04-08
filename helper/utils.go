package helper

import (
	"errors"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/shomali11/util/xhashes"
	"reflect"
	"strconv"
)

func GetVal(name string, mps map[string]interface{}) interface{} {
	v, ok := mps[name]
	if !ok {
		return nil
	}
	return v
}

func GenerateSn(prefix ...string) string {
	sn := xhashes.FNV64(UniqueId())
	snStr := strconv.FormatUint(sn, 10)
	if prefix != nil {
		snStr = prefix[0] + snStr
	}
	return snStr
}

/*
//把汉字转换未Pinyin
func ConvertPinyin(chinese string, mode ...int) string {
	strArr := pinyin.LazyConvert(chinese, nil)
	if strArr != nil {
		return strings.Join(strArr, "")
	}
	return ""
}


*/
//截取字符串 start 起点下标 end 终点下标(不包括)
func Substr(str string, start int, end int) (string, error) {
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
func MapToStruct(fromMap interface{}, toStruct interface{}) interface{} {
	mapstructure.Decode(fromMap, toStruct)
	return toStruct
}

//判断字符是否在数组中
func InArray(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

//string到int64
func StringToInt(val string) int {
	// string到int
	ret, _ := strconv.Atoi(val)
	return ret
}

func StringToInt64(val string) int64 {
	ret, _ := strconv.ParseInt(val, 10, 64)
	return ret
}
func IntToString(val int) string {
	// int到string
	ret := strconv.Itoa(val)
	return ret
}
func Int64ToString(val int64) string {
	// int64到string
	ret := strconv.FormatInt(val, 10)
	return ret
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

func DisplaySize(raw float64) string {
	if raw < 1024 {
		return fmt.Sprintf("%.2fB", raw)
	}

	if raw < 1024*1024 {
		return fmt.Sprintf("%.2fK", raw/1024.0)
	}

	if raw < 1024*1024*1024 {
		return fmt.Sprintf("%.2fM", raw/1024.0/1024.0)
	}

	if raw < 1024*1024*1024*1024 {
		return fmt.Sprintf("%.2fG", raw/1024.0/1024.0/1024.0)
	}

	if raw < 1024*1024*1024*1024*1024 {
		return fmt.Sprintf("%.2fT", raw/1024.0/1024.0/1024.0/1024.0)
	}

	if raw < 1024*1024*1024*1024*1024*1024 {
		return fmt.Sprintf("%.2fP", raw/1024.0/1024.0/1024.0/1024.0/1024.0)
	}

	return "TooLarge"
}

// addslashes() 函数返回在预定义字符之前添加反斜杠的字符串。
// 预定义字符是：
// 单引号（'）
// 双引号（"）
// 反斜杠（\）
func AddSlashes(str string) string {
	tmpRune := []rune{}
	strRune := []rune(str)
	for _, ch := range strRune {
		switch ch {
		case []rune{'\\'}[0], []rune{'"'}[0], []rune{'\''}[0]:
			tmpRune = append(tmpRune, []rune{'\\'}[0])
			tmpRune = append(tmpRune, ch)
		default:
			tmpRune = append(tmpRune, ch)
		}
	}
	return string(tmpRune)
}

// stripslashes() 函数删除由 addslashes() 函数添加的反斜杠。
func StripSlashes(str string) string {
	dstRune := []rune{}
	strRune := []rune(str)
	strLenth := len(strRune)
	for i := 0; i < strLenth; i++ {
		if strRune[i] == []rune{'\\'}[0] {
			i++
		}
		dstRune = append(dstRune, strRune[i])
	}
	return string(dstRune)
}
