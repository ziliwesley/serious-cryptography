package caesar

import "testing"
import "github.com/ziliwesley/serious-cryptography/common"

func TestEncrypt(test *testing.T) {
	fixtures := []common.Fixture {
		{ Input: "abc", Expected: "def" },
		{ Input: "ABC", Expected: "DEF" },
		{ Input: "abcdefghijklmnopqrstuvwxyz", Expected: "defghijklmnopqrstuvwxyzabc" },
		{ Input: "ABCDEFGHIJKLMNOPQRSTUVWXYZ", Expected: "DEFGHIJKLMNOPQRSTUVWXYZABC" },
		{ Input: "ZOO", Expected: "CRR" },
		{ Input: "CAESAR", Expected: "FDHVDU" },
	}
	for _, c := range fixtures {
		output := Encrypt(c.Input)
		if output != c.Expected {
			test.Errorf("Encrypt(%q) == (%q), %q expected",
				c.Input,
				output,
				c.Expected)
		}
	}
}