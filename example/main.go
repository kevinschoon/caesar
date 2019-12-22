package main

import (
	"fmt"

	"github.com/kevinschoon/caesar"
)

func main() {
	// create a cipher with the default character set
	// and letters shifted 10 characters forward.
	cipher := caesar.NewCaesarCipher(10, caesar.DefaultCharacterSet)
	cipherText := cipher.MustEncode("Hello World!")
	fmt.Println(cipherText) // Rovvy*gy!vn+
	decodedText := cipher.MustDecode(cipherText)
	fmt.Println(decodedText) // Hello World!
}
