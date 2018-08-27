package main

import "fmt"
import caeser "github.com/ziliwesley/serious-cryptography/classical/caesar"

func main() {
	plainText := "Hello from Caesar!"
	cipher := caeser.Encrypt(plainText)

	fmt.Printf("Source text: %s\n", plainText)
	fmt.Printf("Cipher: %s\n", cipher)

	decrypted := caeser.Decrypt(cipher)
	
	fmt.Printf("Decrypted: %s\n", decrypted)
}