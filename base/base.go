package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	f15()
}

func f15() {
	fmt.Println(int(time.Monday))
	fmt.Println(int(time.Tuesday))
	fmt.Println(int(time.Wednesday))
	fmt.Println(int(time.Thursday))
	fmt.Println(int(time.Friday))
	fmt.Println(int(time.Saturday))
	fmt.Println(int(time.Sunday))
}

func f14() {
	m1 := map[int32]map[int32]map[int32]int{}
	fmt.Println(m1[1][2])
}

func f13() {
	t1 := time.Unix(1642234169, 0)
	fmt.Println(t1)
	t2 := t1.Format("2006-01-02 15:04:05")
	fmt.Println(t2)
}

func f12() {
	type data struct {
		GameNickName string `json:"game_nick_name"`
	}
	d1 := &data{GameNickName: "拉伊基普林"}
	marshal, err := json.Marshal(d1)
	if err != nil {
		return
	}
	fmt.Println(string(marshal))

}

func f11() {
	fmt.Println()
}

func f10() {
	s1 := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 = s1[5:]
	fmt.Println(s1)
}

func f9() {
	str := ""
	fmt.Println(strings.Replace(str, "http", "https", 1))
}

func f8() {
	type struct1 struct {
		Id int
	}
	s1 := &struct1{Id: 1}
	var s3 *struct1

	//temp := *s1
	//s3 = &temp

	s3 = &(*s1)

	s3.Id = 2
	fmt.Println(s1)
}

func f7() {
	s1 := []int{1, 2, 3, 4, 5}
	for i := len(s1) - 1; i >= 0; i-- {
		fmt.Println(len(s1) - i)
	}
}

func f6(s []int) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(s), func(i, j int) { s[i], s[j] = s[j], s[i] })
}

func f5() {
	var s1 []int
	s1 = append(s1, 1)
	s1 = append(s1, 2)
	s1 = append(s1, 3)
	s1 = append(s1, 4)
	s1 = append(s1, 5)
	f6(s1)
	fmt.Println(s1)

	s2 := []int{1, 2, 3, 4, 5}
	f6(s2)
	fmt.Println(s2)
}

func f4() {
	fmt.Println(time.Unix(1630063960, 0))
}

func f3() {
	var m1 map[int]int
	m1[1] = 1
	fmt.Println(m1)
}

func f2() {
	s1 := []int{1, 2, 3, 4, 5, 6}
	for i := len(s1) - 1; i >= 0; i-- {
		fmt.Println(s1[i])
	}
}

var (
	m map[int]int
)

func f1() {
	fmt.Println(m[0])
}
