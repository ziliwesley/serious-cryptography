// [One-time pad](https://en.wikipedia.org/wiki/One-time_pad)
// In cryptography, the one-time pad (OTP) is an encryption technique that 
// cannot be cracked, but requires the use of a one-time pre-shared key the 
// same size as, or longer than, the message being sent.

package otp

import "encoding/hex"
import "fmt"

// Decrypt method of OTP using XOR
func Decrypt(cipher string, key string) string {
    cipherBytes, err := hex.DecodeString(cipher)

    if (err != nil) {
        fmt.Println("Invalid cipher provided")
        return ""
    }

    output := make([]byte, len(cipherBytes))
    keyBytes, err := hex.DecodeString(key)

    if (err != nil) {
        fmt.Println("Invalid key provided")
        return ""
    }

    for index := 0; index < len(output); index++ {
        srcCharCode := cipherBytes[index]
        key := keyBytes[index]

        // P XOR K
        output[index] = srcCharCode ^ key
    }

    return string(output)
}