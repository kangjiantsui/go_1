package main

import "fmt"

func main() {
	first := []string{"1", "2", "3", "4", "5"}
	var second []*string

	for _, item := range first {
		second = append(second, &item)
	}

	for _, item := range second {
		fmt.Println(*item)
	}
}
