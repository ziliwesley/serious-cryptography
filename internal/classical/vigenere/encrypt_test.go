package vigenere

import (
	"fmt"
	"serious-cryptography/pkg/common"
	"testing"
)

func ExampleEncrypt() {
	fmt.Println(Encrypt("TheyDrinkTheTea", "DUH"))
	// Output: WblbXylhrWblWyh
}

func TestEncrypt(test *testing.T) {
	fixtures := []common.Fixture{
		{
			Input:    "CRYPTO",
			InputAlt: "DUH",
			Expected: "FLFSNV",
		},
		{
			Input:    "Crypto",
			InputAlt: "duH",
			Expected: "Flfsnv",
		},
	}
	for _, c := range fixtures {
		output := Encrypt(c.Input, c.InputAlt)
		if output != c.Expected {
			test.Errorf("Encrypt(%q, %q) == (%q), %q expected",
				c.Input,
				c.InputAlt,
				output,
				c.Expected)
		}
	}
}
