package main

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func main() {
	pub, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		panic(err)
	}
	fmt.Println("PRIVATE_KEY_HEX=" + hex.EncodeToString(priv))
	fmt.Println("PUBLIC_KEY_HEX=" + hex.EncodeToString(pub))
}
