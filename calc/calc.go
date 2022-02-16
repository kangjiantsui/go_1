package main

import (
	"fmt"
	"sort"
	"time"
)

func main() {
	f4()
}

func f4() {
	for i:=1;i>0;i-- {
		fmt.Println(1)
	}
}

func f3() {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	for i := 0; i < len(a); i++ {
		if a[i] == 3 {
			a = append(a[:i+1], a[i+1+1:]...)
		}
		if a[i] == 5 {
			a = append(a[:i+1], a[i+1+1:]...)
		}
		fmt.Println(a[i])
	}

}

func f2() {
	a := []int{1, 2, 3}
	sort.Slice(a, func(i, j int) bool { return a[i] > a[j] })
	fmt.Println(a)
}

func f1() {
	a := 4 + 1 - float64(time.Now().Unix())/10000000000
	fmt.Println(a)
	fmt.Println(int32(a))
}
