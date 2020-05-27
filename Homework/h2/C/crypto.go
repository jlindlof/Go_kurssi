package main

import (
	"crypto/sha512"
	"encoding/base64"
	"fmt"
)

func main() {

	hash := sha512.New()
	hash.Write([]byte("salainen"))
	salattu := base64.URLEncoding.EncodeToString(hash.Sum(nil))

	fmt.Println("Sha512 hash: (Base64 encoded): ", salattu)
}
