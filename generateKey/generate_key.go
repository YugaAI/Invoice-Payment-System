package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func main() {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		panic("failed to generate random key: " + err.Error())
	}

	hexKey := hex.EncodeToString(key)
	fmt.Println("PASETO_SECRET_HEX=" + hexKey)
}
