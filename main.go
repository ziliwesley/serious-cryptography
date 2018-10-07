package main

import (
    "fmt"
    caeser "github.com/ziliwesley/serious-cryptography/classical/caesar"
    "encoding/hex"
    "math/rand"
    "time"
)

func main() {
    plainText := "Hello from Caesar!"
    cipher := caeser.EncryptClassic(plainText)

    fmt.Printf("Source text: %s\n", plainText)
    fmt.Printf("Cipher: %s\n", cipher)

    decrypted := caeser.DecryptClassic(cipher)

    fmt.Printf("Decrypted: %s\n", decrypted)

    bytes := make([]byte, len(plainText))
    copy(bytes[:], plainText)

    fmt.Printf("%s\n", hex.Dump(bytes))
    // use the "space" flag in the format, putting a space between the % and 
    // the x.
    fmt.Printf("% x\n", plainText)

    // Generating a random, fixed-length byte array in Go
    randBytes := make([]byte, len(plainText))
    // Initialize random seed to make sure the result the result will change 
    // every time `rand.Read()` was called
    rand.Seed(time.Now().UnixNano())
    rand.Read(randBytes)
    fmt.Printf("Random Bytes: %x\n", randBytes)
    encoded := hex.EncodeToString(randBytes)
    fmt.Printf("Encoded: %s\n", encoded)

    // Decode hex string to bytes
    decoded, err := hex.DecodeString("41e8d461c6c11ceee9b40de6d7a9334279ec")
    if (err != nil) {
        fmt.Println("error decoding")
    }
    fmt.Printf("%x\n", decoded)
}
