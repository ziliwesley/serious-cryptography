package common

// Fixture defines basic structure of inputs and output
// of a test
type Fixture struct {
	Input,
	InputAlt,
	Expected string
	InputRune    rune
	ExpectedRune rune
}
