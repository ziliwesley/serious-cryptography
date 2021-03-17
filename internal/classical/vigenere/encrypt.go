package vigenere

import (
	letterutil "serious-cryptography/internal/helpers/letterutil"
	"strings"
)

// Encrypt method of vigenere cipher
// The Vigenère cipher is similar to the Caesar cipher, except
// that letters aren’t shifted by three places but rather by
// values defined by a key, a collec- tion of letters that
// represent numbers based on their position in the alphabet.
func Encrypt(src string, key string) string {
	// Create another array of same length
	output := make([]rune, len(src))
	keyUpper := strings.ToUpper(key)
	keyLen := len(keyUpper)

	for index, charCode := range src {
		keyPtr := index % keyLen
		shift := int32(keyUpper[keyPtr]) - letterutil.UpperCaseA
		output[index] = letterutil.ShiftLetter(charCode, shift)
	}

	return string(output)
}
