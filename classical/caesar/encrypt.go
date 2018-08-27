package caesar

// Encrypt method of caesar cipher
// It encrypts a message by shifting each of the letters down 
// three positions in the alphabet, wrapping back around to A
// if the shift reaches Z
func Encrypt(src string) string {
	// Create another array of same length
	output := make([]rune, len(src))

	for index, charCode := range src {
		if (charCode >= 65 && charCode <= 87) ||
			(charCode >= 97 && charCode <= 119) {
			// A-W
			// or
			// a-w
			output[index] = charCode + 3
		} else if (charCode >= 88 && charCode <= 90) ||
			(charCode >= 120 && charCode <= 122) {
			// X-Z
			// or
			// x-z
			output[index] = charCode - 23
		} else {
			output[index] = charCode
		}
	}

	return string(output)
}