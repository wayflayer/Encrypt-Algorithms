package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

func main() {
	key := []byte("aeskey123aeskey1")
	message := []byte("Hello, World!")

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	ciphertext := make([]byte, aes.BlockSize+len(message))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], message)

	fmt.Printf("AES Encrypted: %xn", ciphertext)

	decrypted := make([]byte, len(message))
	stream = cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(decrypted, ciphertext[aes.BlockSize:])

	fmt.Printf("AES Decrypted: %sn", decrypted)
}
