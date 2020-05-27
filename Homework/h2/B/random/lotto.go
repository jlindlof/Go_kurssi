package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 40
	fmt.Println("Kierroksen lottonumerot: ")

	for i := 0; i < 7; i++ {
		fmt.Println(rand.Intn(max-min+1) + min)
	}

}
