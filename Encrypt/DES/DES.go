package main

import (
	"crypto/cipher"
	"crypto/des"
	"crypto/rand"
	"fmt"
	"io"
)

func main() {
	key := []byte("deskey12")
	message := []byte("Hello, World!")

	block, err := des.NewCipher(key)
	if err != nil {
		panic(err)
	}

	ciphertext := make([]byte, des.BlockSize+len(message))
	iv := ciphertext[:des.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[des.BlockSize:], message)

	fmt.Printf("DES Encrypted: %xn", ciphertext)

	decrypted := make([]byte, len(message))
	stream = cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(decrypted, ciphertext[des.BlockSize:])

	fmt.Printf("DES Decrypted: %sn", decrypted)
}
