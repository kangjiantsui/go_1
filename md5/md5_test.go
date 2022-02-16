package main

import (
	"crypto/md5"
	"fmt"
	"reflect"
	"sort"
	"testing"
	"unsafe"
)

func getMd5String2(b []byte) string {
	return fmt.Sprintf("%x", md5.Sum(b))
}

func slice(s string) (b []byte) {
	pbytes := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	pstring := (*reflect.StringHeader)(unsafe.Pointer(&s))
	pbytes.Data = pstring.Data
	pbytes.Len = pstring.Len
	pbytes.Cap = pstring.Len
	return
}

func TestCreateMd5Str(t *testing.T) {
	params := map[string]string{
		"channel_id": "100",
		"appid":      "1504860311",
		"device_id":  "AB821758-4D6D-457A-98C7-B0D33FE898B2",
		"date_start": "1642411012",
		"date_end":   "1644940792",
		"period_id":  "1",
	}
	key := "82eaf5b418c00c3a475b0bd1f5ccd462"

	type param struct {
		Key string
		Val string
	}
	var paramsList []*param
	for k, v := range params {
		paramsList = append(paramsList, &param{Key: k, Val: v})
	}
	sort.Slice(paramsList, func(i, j int) bool { return paramsList[i].Key < paramsList[j].Key })

	result := ""
	for _, e := range paramsList {
		if result == "" {
			result += e.Key + "=" + e.Val
			continue
		}
		result += "&" + e.Key + "=" + e.Val
	}

	sign := getMd5String2(slice(key + result))

	t.Logf(sign)
}
