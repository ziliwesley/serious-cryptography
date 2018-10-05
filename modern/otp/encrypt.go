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

func encryptWithBytes(src string, keyBytes []byte) string {
    output := make([]byte, len(src))

    for index := 0; index < len(output); index++ {
        srcCharCode := src[index]
        key := keyBytes[index]

        // P XOR K
        output[index] = srcCharCode ^ key
    }

    return hex.EncodeToString(output)
}

// EncryptWithRandomBytes method encrypt a source of  
// type using OTP(XOR) with a number as RNG seed to
// generate pseudo random bytes as key
func EncryptWithRandomBytes(src string, seed int64) string {
    rng := lcg.GetRandomNumberGenerator()
    rng.SetSeed(seed)
    keyBytes := rng.RandomBytes(int32(len(src)))

    return encryptWithBytes(src, keyBytes)
}

// Encrypt method encrypt a source of text using 
// OTP(XOR) with bytes in hex as key
func Encrypt(src string, key string) string {
    keyBytes, err := hex.DecodeString(key)

    if (err != nil) {
        fmt.Println("Invalid key provided")
        keyBytes = make([]byte, len(src))
    }

    return encryptWithBytes(src, keyBytes)
}