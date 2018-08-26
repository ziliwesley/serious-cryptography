package caesar

import "testing"

type Fixture struct {
	input, expected string
}

func Test(test *testing.T) {
	fixtures := []Fixture {
		{ "abc", "def" },
		{ "ABC", "DEF" },
		{ "abcdefghijklmnopqrstuvwxyz", "defghijklmnopqrstuvwxyzabc" },
		{ "ABCDEFGHIJKLMNOPQRSTUVWXYZ", "DEFGHIJKLMNOPQRSTUVWXYZABC" },
		{ "ZOO", "CRR" },
		{ "CAESAR", "FDHVDU" },
	}
	for _, c := range fixtures {
		output := Encrypt(c.input)
		if output != c.expected {
			test.Errorf("Encrypt(%q) == (%q), %q expected",
				c.input,
				output,
				c.expected)
		}
	}
}