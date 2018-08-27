package caesar

// Decrypt method of caesar cipher
func Decrypt(cipher string) string {
	// Create another array of same length
	output := make([]rune, len(cipher))

	for index, charCode := range cipher {
		if (charCode >= 65 && charCode <= 67) ||
			(charCode >= 97 && charCode <= 99) {
			// A-C
			// or
			// a-c
			output[index] = charCode + 23
		} else if (charCode >= 68 && charCode <= 90) ||
			(charCode >= 100 && charCode <= 122) {
			// D-Z
			// or
			// d-z
			output[index] = charCode - 3
		} else {
			output[index] = charCode
		}
	}

	return string(output)
}