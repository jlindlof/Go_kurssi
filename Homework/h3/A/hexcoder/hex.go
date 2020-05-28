//converts string to hex dump

package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	content := []byte("string to encode")
	fmt.Printf("%s", hex.Dump(content))

}
