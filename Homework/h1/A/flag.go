package main

import (
	"flag"
	"fmt"
)

func main() {
	var onnenluku int
	flag.IntVar(&onnenluku, "onnenluku", 9999, "Anna luku")
	flag.Parse()

	if onnenluku == 7 {
		fmt.Println("Annoit onnenluvun")
	} else if onnenluku == 9999 {
		fmt.Println("Et antanut lukua")
	} else {
		fmt.Println("Et antanut onnenlukua")
	}

}
