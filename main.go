package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Key struct {
	Key    string `json:"Key"`
	Secret string `json:"Secret"`
}

func main() {

	var key Key

	file, err := os.Open(".key.json")
	if err !=nil{
		fmt.Println("Unable to Open File")
		return
	}
	defer file.Close()

	decoder:= json.NewDecoder(file)

	decoder.Decode(&key)

	fmt.Println(key)

}