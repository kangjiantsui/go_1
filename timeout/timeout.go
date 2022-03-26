package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	done := make(chan error)
	go func() {
		done <- f1()
	}()

	select {
	case err := <-done:
		panic(err)
	case <-ctx.Done():
		fmt.Printf("%v,成功", time.Now())
	}
}

func f1() error {
	time.Sleep(time.Second)
	return nil
}
