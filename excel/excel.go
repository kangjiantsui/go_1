package main

import (
	"github.com/xuri/excelize/v2"
)

func main() {
	m1 := getShuShuIdfa()
	m2 := getQimaiIdfa()
	for e := range m2 {
		if _, ok := m2[e]; !ok {
			println(m1)
		} else {
			println("激活,", e)
		}
	}
}

func getShuShuIdfa() map[string]interface{} {
	excel, err := excelize.OpenFile("事件分析_20220328-20220328.xlsx")
	if err != nil {
		panic(err)
	}
	rows, err := excel.GetRows("事件分析_20220328-20220328")
	if err != nil {
		panic(err)
	}
	ret := map[string]interface{}{}
	index := 0
	for _, e := range rows {
		for _, f := range e {
			index++
			if index%2 == 0 || index == 1 || index%3 == 0 || index%4 == 0 {
				continue
			}
			ret[f] = nil
		}
	}
	return ret
}

func getQimaiIdfa() map[string]interface{} {
	qimaiExcel, err := excelize.OpenFile("(2022-03-28_2022-03-28)_单挑篮球_300_IDFA.xlsx")
	if err != nil {
		panic(err)
	}
	rows, err := qimaiExcel.GetRows("单挑篮球_300")
	if err != nil {
		panic(err)
	}
	qimaiIdfaMap := map[string]interface{}{}
	index := 0
	for _, e := range rows {
		for _, f := range e {
			index++
			if index%2 == 0 || index == 1 {
				continue
			}
			qimaiIdfaMap[f] = nil
		}
	}
	return qimaiIdfaMap
}
