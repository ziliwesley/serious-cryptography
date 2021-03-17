package otp

import (
	"fmt"
	"serious-cryptography/pkg/common"
	"testing"
)

func ExampleDecrypt() {
	fmt.Println(Decrypt("3c3b6eed4f", "545e028120"))
	// Output: hello
}

func ExampleDecryptWithRandomBytes() {
	fmt.Println(DecryptWithRandomBytes("924a7326b408ea", 1538754611))
	// Output: destiny
}

func TestDecrypt(test *testing.T) {
	fixtures := []common.Fixture{
		{Input: "07ad19a9d6", Expected: "hello"},
		{Input: "18a707a9dd", Expected: "world"},
		{Input: "1ba118a0ca", Expected: "times"},
	}
	for _, c := range fixtures {
		output := Decrypt(c.Input, "6fc875c5b9")
		if output != c.Expected {
			test.Errorf("Decrypt(%q, 2) == (%q), %q expected",
				c.Input,
				output,
				c.Expected)
		}
	}
}
