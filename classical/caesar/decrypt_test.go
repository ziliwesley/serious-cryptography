package caesar

import "testing"
import "github.com/ziliwesley/serious-cryptography/common"

func TestDecrypt(test *testing.T)  {
	fixtures := []common.Fixture {
		{ Input: "def", Expected: "abc" },
		{ Input: "DEF", Expected: "ABC" },
		{ Input: "defghijklmnopqrstuvwxyzabc", Expected: "abcdefghijklmnopqrstuvwxyz" },
		{ Input: "DEFGHIJKLMNOPQRSTUVWXYZABC", Expected: "ABCDEFGHIJKLMNOPQRSTUVWXYZ" },
		{ Input: "CRR", Expected: "ZOO" },
		{ Input: "FDHVDU", Expected: "CAESAR" },
	}

	for _, c := range fixtures {
		output := Decrypt(c.Input)
		if output != c.Expected {
			test.Errorf("Decrypt(%q) == (%q), %q expected",
				c.Input,
				output,
				c.Expected)
		}
	}
}