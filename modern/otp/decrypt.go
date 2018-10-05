// [One-time pad](https://en.wikipedia.org/wiki/One-time_pad)
// In cryptography, the one-time pad (OTP) is an encryption technique that 
// cannot be cracked, but requires the use of a one-time pre-shared key the 
// same size as, or longer than, the message being sent.

package otp

import (
    "encoding/hex"
    "fmt"
    "github.com/ziliwesley/serious-cryptography/random/lcg"
)

func decryptWithBytes(cipher string, keyBytes []byte) string {
    cipherBytes, err := hex.DecodeString(cipher)

    if (err != nil) {
        fmt.Println("Invalid cipher provided")
        return ""
    }

    output := make([]byte, len(cipherBytes))

    for index := 0; index < len(output); index++ {
        srcCharCode := cipherBytes[index]
        key := keyBytes[index]

        // P XOR K
        output[index] = srcCharCode ^ key
    }

    return string(output)
}

// DecryptWithRandomBytes method decrypt a cipher using 
// OTP(XOR) with a number as RNG seed to generate pseudo 
// random bytes as key
func DecryptWithRandomBytes(cipher string, seed int64) string {
    rng := lcg.GetRandomNumberGenerator()
    rng.SetSeed(seed)
    keyBytes := rng.RandomBytes(int32(len(cipher) / 2))

    return decryptWithBytes(cipher, keyBytes)
}

// Decrypt method decrypt a cipher using OTP(XOR) 
// with bytes in hex as key
func Decrypt(cipher string, key string) string {
    keyBytes, err := hex.DecodeString(key)

    if (err != nil) {
        fmt.Println("Invalid key provided")
        return ""
    }

    return decryptWithBytes(cipher, keyBytes)
}