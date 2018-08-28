package caesar

// Encrypt cipher using a shift of 3
func EncryptClassic(cipher string) string {
	return Encrypt(cipher, 3)
}

// Encrypt method of caesar cipher
// It encrypts a message by shifting each of the letters down 
// three positions in the alphabet, wrapping back around to A
// if the shift reaches Z
func Encrypt(src string, shift int32) string {
	// Create another array of same length
	output := make([]rune, len(src))

	for index, charCode := range src {
		if (charCode >= 65 && charCode < 91 - shift) ||
			(charCode >= 97 && charCode < 123 -shift) {
			// A-W
			// or
			// a-w
			output[index] = charCode + shift
		} else if (charCode >= 91 - shift && charCode <= 90) ||
			(charCode >= 123 -shift && charCode <= 122) {
			// X-Z
			// or
			// x-z
			output[index] = charCode - 26 + shift
		} else {
			output[index] = charCode
		}
	}

	return string(output)
}