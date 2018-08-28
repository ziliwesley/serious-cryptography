package caesar

// Decrypt cipher using a shift of 3
func DecryptClassic(cipher string) string {
	return Decrypt(cipher, 3)
}

// Decrypt method of caesar cipher
func Decrypt(cipher string, shift int32) string {
	// Create another array of same length
	output := make([]rune, len(cipher))

	for index, charCode := range cipher {
		if (charCode >= 65 && charCode < 65 + shift) ||
			(charCode >= 97 && charCode < 97 +shift) {
			// A-C
			// or
			// a-c
			output[index] = charCode + 26 - shift
		} else if (charCode >= 65 + shift && charCode <= 90) ||
			(charCode >= 97 + shift && charCode <= 122) {
			// D-Z
			// or
			// d-z
			output[index] = charCode - shift
		} else {
			output[index] = charCode
		}
	}

	return string(output)
}