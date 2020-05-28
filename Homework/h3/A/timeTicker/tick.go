//Program that Prints minutes and seconds every second

package main

import (
	"fmt"
	"time"
)

func main() {

	loop(1000*time.Millisecond, printSomething)

}

func loop(d time.Duration, f func(time.Time)) {
	for a := range time.Tick(d) {
		f(a)
	}
}

func printSomething(t time.Time) {
	fmt.Println(time.Now().Minute(), ":", time.Now().Second())
}
