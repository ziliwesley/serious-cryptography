package main

import "fmt"
import caeser "github.com/ziliwesley/serious-cryptography/classical/caesar"

func main() {
	plainText := "CAESAR"

	t := []byte{1, 2, 3, 4}
	at := make([]byte, len(t))
	// slice := t[1:]
	for i := range t {
		at[i] = t[i]
	}

	fmt.Println(at)
	
	fmt.Println(caeser.Encrypt(plainText))
	
	fmt.Println("Hello, 世界")
}