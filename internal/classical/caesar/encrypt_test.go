package caesar

import (
	"fmt"
	"serious-cryptography/pkg/common"
	"testing"
)

func ExampleEncrypt() {
	fmt.Println(Encrypt("hello", 1))
	// Output: ifmmp
}

func ExampleEncryptClassic() {
	fmt.Println(EncryptClassic("CAESAR"))
	// Output: FDHVDU
}

func TestEncrypt(test *testing.T) {
	fixtures := []common.Fixture{
		{Input: "abc", Expected: "cde"},
		{Input: "ABC", Expected: "CDE"},
		{Input: "abcdefghijklmnopqrstuvwxyz", Expected: "cdefghijklmnopqrstuvwxyzab"},
		{Input: "ABCDEFGHIJKLMNOPQRSTUVWXYZ", Expected: "CDEFGHIJKLMNOPQRSTUVWXYZAB"},
		{Input: "ZOO", Expected: "BQQ"},
		{Input: "CAESAR", Expected: "ECGUCT"},
	}
	for _, c := range fixtures {
		output := Encrypt(c.Input, 2)
		if output != c.Expected {
			test.Errorf("Encrypt(%q, 2) == (%q), %q expected",
				c.Input,
				output,
				c.Expected)
		}
	}
}

func TestEncryptClassic(test *testing.T) {
	fixtures := []common.Fixture{
		{Input: "abc", Expected: "def"},
		{Input: "ABC", Expected: "DEF"},
		{Input: "abcdefghijklmnopqrstuvwxyz", Expected: "defghijklmnopqrstuvwxyzabc"},
		{Input: "ABCDEFGHIJKLMNOPQRSTUVWXYZ", Expected: "DEFGHIJKLMNOPQRSTUVWXYZABC"},
		{Input: "ZOO", Expected: "CRR"},
		{Input: "CAESAR", Expected: "FDHVDU"},
	}
	for _, c := range fixtures {
		output := EncryptClassic(c.Input)
		if output != c.Expected {
			test.Errorf("EncryptClassic(%q) == (%q), %q expected",
				c.Input,
				output,
				c.Expected)
		}
	}
}
