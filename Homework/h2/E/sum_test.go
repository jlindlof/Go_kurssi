package main

import (
	"fmt"
	"testing"
)

func main() {

	fmt.Println("Sum of 1 + 2 =", Sum(1, 2))

}

func Sum(a int, b int) int {
	return a + b
}

func TestSum(t *testing.T) {
	yhteensa := Sum(1, 2)
	if yhteensa != 3 {
		t.Errorf("Summa oli väärin, tulos: %d, piti olla: %d.", yhteensa, 3)
	}
}
