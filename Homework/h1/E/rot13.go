package main

import (
	"flag"
	"fmt"
	"strings"
)

func rot13(merkki rune) rune {
	if merkki >= 'a' && merkki <= 'z' {
		if merkki >= 'm' {
			return merkki - 13
		} else {
			return merkki + 13
		}
	} else if merkki >= 'A' && merkki <= 'Z' {
		if merkki >= 'M' {
			return merkki - 13
		} else {
			return merkki + 13
		}
	}
	return merkki
}

func main() {
	var input string
	flag.StringVar(&input, "input", "Test", "Anna merkkijono (a-zA-Z)")
	flag.Parse()
	mapped := strings.Map(rot13, input)
	fmt.Println(mapped)

}
