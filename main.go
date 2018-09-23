package main

import "fmt"
import caeser "github.com/ziliwesley/serious-cryptography/classical/caesar"
import "encoding/hex"

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
    fmt.Printf("% x", plainText)
}
