package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"net/http"
)

func main() {
	cli.NewApp()

	http.Handle("/", http.FileServer(http.Dir(".")))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(fmt.Sprintf("http server 端口监听失败,err:%s", err.Error()))
		return
	}

}
