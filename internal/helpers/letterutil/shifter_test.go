package letterutil

import (
	"fmt"
	"serious-cryptography/pkg/common"
	"testing"
)

func ExampleShiftLetter() {
	fmt.Printf("%c\n", ShiftLetter('n', 1))
	// Output: o
}

func TestShiftLetterBy25(test *testing.T) {
	fixtures := []common.Fixture{
		{InputRune: 'a', ExpectedRune: 'z'},
		{InputRune: 'b', ExpectedRune: 'a'},
		{InputRune: 'c', ExpectedRune: 'b'},
		{InputRune: 'A', ExpectedRune: 'Z'},
		{InputRune: 'B', ExpectedRune: 'A'},
		{InputRune: 'C', ExpectedRune: 'B'},
		{InputRune: ',', ExpectedRune: ','},
		{InputRune: '中', ExpectedRune: '中'},
	}

	for _, c := range fixtures {
		output := ShiftLetter(c.InputRune, 25)
		if output != c.ExpectedRune {
			test.Errorf("ShiftLetter(%q, 25) == (%q), %q expected",
				c.InputRune,
				output,
				c.ExpectedRune)
		}
	}
}

func TestShiftLetterByNeg25(test *testing.T) {
	fixtures := []common.Fixture{
		{InputRune: 'z', ExpectedRune: 'a'},
		{InputRune: 'a', ExpectedRune: 'b'},
		{InputRune: 'b', ExpectedRune: 'c'},
		{InputRune: 'Z', ExpectedRune: 'A'},
		{InputRune: 'A', ExpectedRune: 'B'},
		{InputRune: 'B', ExpectedRune: 'C'},
	}

	for _, c := range fixtures {
		output := ShiftLetter(c.InputRune, -25)
		if output != c.ExpectedRune {
			test.Errorf("ShiftLetter(%q, -25) == (%q), %q expected",
				c.InputRune,
				output,
				c.ExpectedRune)
		}
	}
}
