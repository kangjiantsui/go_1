package main

import (
	"encoding/json"
	"go_01/pb"
	"io"
	"log"
	"os"
)

func main() {
	jsonFile, err := os.Open("C:\\workspace\\go_1\\json\\demo1.json")
	if err != nil {
		return
	}
	if err != nil {
		log.Fatal(err.Error())
	}
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err.Error())
	}
	shopInfo := &pb.ShopNewInfo{}
	err = json.Unmarshal(byteValue, shopInfo)
	if err != nil {
		log.Fatal(err.Error())
	}

}
