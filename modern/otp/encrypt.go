// [One-time pad](https://en.wikipedia.org/wiki/One-time_pad)
// In cryptography, the one-time pad (OTP) is an encryption technique that 
// cannot be cracked, but requires the use of a one-time pre-shared key the 
// same size as, or longer than, the message being sent.

package otp

import "encoding/hex"
import "fmt"

// Encrypt method of OTP using XOR
func Encrypt(src string, key string) string {
    output := make([]byte, len(src))
    keyBytes, err := hex.DecodeString(key)

    if (err != nil) {
        fmt.Println("Invalid key provided")
        keyBytes = make([]byte, len(src))
    }

    for index := 0; index < len(output); index++ {
        srcCharCode := src[index]
        key := keyBytes[index]

        // P XOR K
        output[index] = srcCharCode ^ key
    }

    return hex.EncodeToString(output)
}