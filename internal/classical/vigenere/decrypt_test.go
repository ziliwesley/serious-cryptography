package vigenere

import (
	"fmt"
	"serious-cryptography/pkg/common"
	"testing"
)

func ExampleDecrypt() {
	fmt.Println(Decrypt("WblbXylhrWblWyh", "DUH"))
	// Output: TheyDrinkTheTea
}

func TestDecrypt(test *testing.T) {
	fixtures := []common.Fixture{
		{
			Input:    "FLFSNV",
			InputAlt: "DUH",
			Expected: "CRYPTO",
		},
		{
			Input:    "Flfsnv",
			InputAlt: "duH",
			Expected: "Crypto",
		},
	}
	for _, c := range fixtures {
		output := Decrypt(c.Input, c.InputAlt)
		if output != c.Expected {
			test.Errorf("Decrypt(%q, %q) == (%q), %q expected",
				c.Input,
				c.InputAlt,
				output,
				c.Expected)
		}
	}
}
