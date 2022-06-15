package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unsafe"
)

// StringToInt : string -> int
func StringToInt(s string, optionalNum ...int) int {
	if s == "" {
		if len(optionalNum) != 0 {
			return optionalNum[0]
		}
		return 0
	}

	temp, _ := strconv.Atoi(s)
	if temp == 0 && len(optionalNum) != 0 {
		return optionalNum[0]
	}

	return temp
}

// StringToInt16 : string -> int16
func StringToInt16(s string) int16 {
	if s == "" {
		return 0
	}

	TmpInt, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}

	return int16(TmpInt)
}

// StringToInt32 : string -> int32
func StringToInt32(s string) int32 {
	if s == "" {
		return 0
	}

	TmpInt, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}

	return int32(TmpInt)
}

// StringToInt64 : string -> int64
func StringToInt64(s string) int64 {
	if s == "" {
		return 0
	}

	tmp, _ := strconv.ParseInt(s, 10, 64)

	return tmp
}

// StringToUint : string -> uint
func StringToUint(s string) uint {
	if s == "" {
		return 0
	}

	TmpInt, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}

	return uint(TmpInt)
}

// StringToUint16 : string -> uint16
func StringToUint16(s string) uint16 {
	if s == "" {
		return 0
	}

	TmpInt, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}

	return uint16(TmpInt)
}

// StringToUint32 : string -> uint32
func StringToUint32(s string) uint32 {
	if s == "" {
		return 0
	}

	TmpInt, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}

	return uint32(TmpInt)
}

// StringToUint64 : string -> uint64
func StringToUint64(s string) uint64 {
	if s == "" {
		return 0
	}

	TmpInt, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}

	return uint64(TmpInt)
}

// StringToFloat64 : string -> float64
func StringToFloat64(s string) float64 {
	if s == "" {
		return 0
	}

	tmp, _ := strconv.ParseFloat(s, 64)

	return tmp
}

// Int64ToString : int64 -> string
func Int64ToString(a int64) string {
	return strconv.FormatInt(a, 10)
}

// Int32ToString : int32 -> string
func Int32ToString(a int32) string {
	return strconv.FormatInt(int64(a), 10)
}

// Int16ToString : int16 -> string
func Int16ToString(a int16) string {
	return strconv.FormatInt(int64(a), 10)
}

// IntToString : int -> string
func IntToString(a int) string {
	return strconv.Itoa(a)
}

// Uint64ToString : uint64 -> string
func Uint64ToString(a uint64) string {
	return strconv.FormatUint(a, 10)
}

// Uint32ToString : uint32 -> string
func Uint32ToString(a uint32) string {
	return strconv.FormatUint(uint64(a), 10)
}

// Uint16ToString : uint16 -> string
func Uint16ToString(a uint16) string {
	return strconv.FormatUint(uint64(a), 10)
}

// Float64ToString : float64 -> string
func Float64ToString(a float64) string {
	return strconv.FormatFloat(a, 'f', -1, 64)
}

func Float64SliceToStringSlice(arr []float64) []string {
	ret := make([]string, 0)

	for _, item := range arr {
		ret = append(ret, Float64ToString(item))
	}
	return ret
}

// StringToInt32Slice :
func StringToInt32Slice(s string, sep string) (ret []int32) {
	tokens := strings.Split(s, sep)
	for _, k := range tokens {
		i, err := strconv.ParseInt(k, 10, 64)
		if err != nil {
			return nil
		}
		ret = append(ret, int32(i))
	}
	return
}

// BytesToString convert bytes to string
func BytesToString(b *[]byte) *string {
	s := bytes.NewBuffer(*b)
	r := s.String()
	return &r
}

func ToJsonString(v interface{}) string {
	ret, err := json.Marshal(v)
	if err != nil {
		return fmt.Sprintf(`{"err":"%s"}`, err.Error())
	}
	return String(ret)
}

// ToString convert some type to string
// []string{"a","b"} => "a,b"
// []int{1,2} => "1,2"
// []string{} => ""
func ToString(v interface{}) string {
	switch x := v.(type) {
	case string:
		return x
	case int:
		return IntToString(x)
	case int64:
		return Int64ToString(x)
	case float64:
		return Float64ToString(x)
	case json.Number:
		return x.String()
	case []string:
		return strings.Join(x, ",")
	case []int, []int64, []float64:
		return ToString(ToStringSlice(x))
	case map[string]interface{}:
		data, _ := json.Marshal(x)
		return String(data)
	default:
		return fmt.Sprint(v)
	}
}

// ToStringSlice ToStringSlice
func ToStringSlice(v interface{}) []string {
	switch x := v.(type) {
	case string:
		return strings.Split(x, ",")
	case []string:
		return x
	case []int:
		s := make([]string, len(x))
		for i := range x {
			s[i] = ToString(x[i])
		}
		return s
	case []int64:
		s := make([]string, len(x))
		for i := range x {
			s[i] = ToString(x[i])
		}
		return s
	case []float64:
		s := make([]string, len(x))
		for i := range x {
			s[i] = ToString(x[i])
		}
		return s
	case []interface{}:
		s := make([]string, len(x))
		for i := range x {
			s[i] = ToString(x[i])
		}
		return s
	default:
		return nil
	}
}

func String(b []byte) (s string) {
	pbytes := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	pstring := (*reflect.StringHeader)(unsafe.Pointer(&s))
	pstring.Data = pbytes.Data
	pstring.Len = pbytes.Len
	return
}

func Slice(s string) (b []byte) {
	pbytes := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	pstring := (*reflect.StringHeader)(unsafe.Pointer(&s))
	pbytes.Data = pstring.Data
	pbytes.Len = pstring.Len
	pbytes.Cap = pstring.Len
	return
}

// 数字转字符串 数组
func Int64Slice2StringSlice(arr []int64) []string {
	strs := make([]string, 0)
	for _, v := range arr {
		strs = append(strs, strconv.FormatInt(v, 10))
	}
	return strs
}

// 数字转字符串 数组
func Uint64Slice2StringSlice(arr []uint64) []string {
	strs := make([]string, 0)
	for _, v := range arr {
		strs = append(strs, strconv.FormatUint(v, 10))
	}
	return strs
}

func StringSlice2Uint64Slice(arr []string) []uint64 {
	ret := make([]uint64, 0)
	for _, v := range arr {
		val, _ := strconv.ParseUint(v, 10, 64)
		ret = append(ret, val)
	}
	return ret
}
