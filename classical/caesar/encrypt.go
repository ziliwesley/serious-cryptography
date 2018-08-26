package caesar

// import "bytes"

// Encrypt plain text which contains:
// letters: a to z
// and
// letters: A to Z
func Encrypt(src string) string {
	input := []byte(src)
	// Create another array of same length
	output := make([]byte, len(input))

	for index := range input {
		charCode := input[index]

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