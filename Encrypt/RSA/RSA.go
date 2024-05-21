package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println(err)
		return
	}
	publicKey := &privateKey.PublicKey

	message := []byte("Hello, World!")

	ciphertext, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, publicKey, message, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Encrypted message: %xn", ciphertext)

	plaintext, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, ciphertext, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Decrypted message: %sn", plaintext)

	fmt.Println("Encrypted (hex):", hex.EncodeToString(ciphertext))

	if string(plaintext) == "Hello, World!" {
		fmt.Println("Decryption successful, message matches original")
	}
}
