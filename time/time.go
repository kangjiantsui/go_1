package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func main() {
	f6()
}

func f6() {
	_, err := json.Marshal(nil)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func f5() {
	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	fmt.Println(s1[21:len(s1)])
}

func f4() {
	now := time.Now()
	day := 0
	if now.Weekday() == time.Sunday {
		day = time.Now().AddDate(0, 0, int(time.Monday-now.Weekday())).Day() - 7
	} else {
		if int(now.Weekday()) > now.Day() {
			day = now.Day() - int(now.Weekday()) + 1
		} else {
			day = time.Now().AddDate(0, 0, int(time.Monday-now.Weekday())).Day()
		}
	}
	fmt.Println(time.Date(now.Year(), now.Month(), day, 5, 0, 0, 0, time.Local))
}

func f3() {
	fmt.Println(int(time.Monday - time.Now().Weekday()))
	fmt.Println(int(time.Monday))
	fmt.Println(int(time.Sunday))
}

func f2() {
	location, err := time.LoadLocation("Local")
	if err != nil {
		return
	}
	fmt.Println(float64(time.Now().In(location).Unix()%86400) / 3600)
}

func f1() {
	fmt.Println(time.Now().Day())
}
