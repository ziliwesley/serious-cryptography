package caesar

import letterutil "github.com/ziliwesley/serious-cryptography/classical/helpers/letterutil"

// EncryptClassic provides a shortcut to encrypt text by using
// a static shift of 3
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
        output[index] = letterutil.ShiftLetter(charCode, shift)
    }

    return string(output)
}