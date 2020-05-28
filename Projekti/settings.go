package main

import (
	"fmt"
	"io/ioutil"
)

func readAPIKeyFromFile() string {
	data, err := ioutil.ReadFile("apikey")
	if err != nil {
		fmt.Println("File reading error", err)
	}
	return string(data)
}
