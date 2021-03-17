package otp

import (
	"fmt"
	"serious-cryptography/pkg/common"
	"testing"
)

func ExampleEncrypt() {
	fmt.Println(Encrypt("hello", "545e028120"))
	// Output: 3c3b6eed4f
}

func ExampleEncryptWithRandomBytes() {
	fmt.Println(EncryptWithRandomBytes("destiny", 1538754611))
	// Output: 924a7326b408ea
}

func TestEncrypt(test *testing.T) {
	fixtures := []common.Fixture{
		{Input: "hello", Expected: "3c3b6eed4f"},
		{Input: "world", Expected: "233170ed44"},
		{Input: "times", Expected: "20376fe453"},
	}
	for _, c := range fixtures {
		output := Encrypt(c.Input, "545e028120")
		if output != c.Expected {
			test.Errorf("Encrypt(%q, 2) == (%q), %q expected",
				c.Input,
				output,
				c.Expected)
		}
	}
}
