package caesar

import "testing"
import "github.com/ziliwesley/serious-cryptography/common"

func TestDecrypt(test *testing.T)  {
	fixtures := []common.Fixture {
		{ Input: "bcd", Expected: "abc" },
		{ Input: "BCD", Expected: "ABC" },
		{ Input: "bcdefghijklmnopqrstuvwxyza", Expected: "abcdefghijklmnopqrstuvwxyz" },
		{ Input: "BCDEFGHIJKLMNOPQRSTUVWXYZA", Expected: "ABCDEFGHIJKLMNOPQRSTUVWXYZ" },
		{ Input: "APP", Expected: "ZOO" },
		{ Input: "DBFTBS", Expected: "CAESAR" },
	}

	for _, c := range fixtures {
		output := Decrypt(c.Input, 1)
		if output != c.Expected {
			test.Errorf("Decrypt(%q, 1) == (%q), %q expected",
				c.Input,
				output,
				c.Expected)
		}
	}
}

func TestDecryptClassic(test *testing.T)  {
	fixtures := []common.Fixture {
		{ Input: "def", Expected: "abc" },
		{ Input: "DEF", Expected: "ABC" },
		{ Input: "defghijklmnopqrstuvwxyzabc", Expected: "abcdefghijklmnopqrstuvwxyz" },
		{ Input: "DEFGHIJKLMNOPQRSTUVWXYZABC", Expected: "ABCDEFGHIJKLMNOPQRSTUVWXYZ" },
		{ Input: "CRR", Expected: "ZOO" },
		{ Input: "FDHVDU", Expected: "CAESAR" },
	}

	for _, c := range fixtures {
		output := DecryptClassic(c.Input)
		if output != c.Expected {
			test.Errorf("DecryptClassic(%q) == (%q), %q expected",
				c.Input,
				output,
				c.Expected)
		}
	}
}