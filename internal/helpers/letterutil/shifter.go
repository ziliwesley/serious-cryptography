package letterutil

// LowerCaseA equals to ascii code of letter 'a'
const LowerCaseA = 97

// LowerCaseZ equals to ascii code of letter 'z'
const LowerCaseZ = 122

// UpperCaseA equals to ascii code of letter 'A'
const UpperCaseA = 65

// UpperCaseZ equals to ascii code of letter 'Z'
const UpperCaseZ = 90

// ShiftLetter shifts a rune (letter) by a given
// number of places
func ShiftLetter(charCode rune, shift int32) rune {
	// actual value of shift should not be larger
	// than 26 or smaler than -26
	actualShift := shift % 26

	// Handle lowercase letters (a-z)
	if charCode >= LowerCaseA &&
		charCode <= LowerCaseZ {
		shifted := charCode + actualShift
		if shifted > LowerCaseZ {
			return LowerCaseA + (shifted - LowerCaseZ) - 1
		}
		if shifted < LowerCaseA {
			return LowerCaseZ - (LowerCaseA - shifted) + 1
		}
		return shifted
	}

	// Handle uppercase letters (a-z)
	if charCode >= UpperCaseA &&
		charCode <= UpperCaseZ {
		shifted := charCode + actualShift
		if shifted > UpperCaseZ {
			return UpperCaseA + (shifted - UpperCaseZ) - 1
		}
		if shifted < UpperCaseA {
			return UpperCaseZ - (UpperCaseA - shifted) + 1
		}
		return shifted
	}

	// Keep unshifted
	return charCode
}
