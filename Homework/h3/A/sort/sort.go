// Program that sorts strings
package main

import (
	"fmt"
	"sort"
)

func main() {

	lista := []string{"f", "b", "a", "c", "e", "d"}

	fmt.Println("unsorted: ", lista)

	sort.Strings(lista)

	fmt.Println("sorted\t: ", lista)
}
