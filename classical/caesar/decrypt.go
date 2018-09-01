package caesar

import letterutil "github.com/ziliwesley/serious-cryptography/classical/helpers/letterutil"

// DecryptClassic provides a shortcut to decrypt cipher by using
// a static shift of 3
func DecryptClassic(cipher string) string {
    return Decrypt(cipher, 3)
}

// Decrypt method of caesar cipher
func Decrypt(cipher string, shift int32) string {
    // Create another array of same length
    output := make([]rune, len(cipher))

    for index, charCode := range cipher {
        output[index] = letterutil.ShiftLetter(charCode, -shift)
    }

    return string(output)
}