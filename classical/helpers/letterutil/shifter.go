package letterutil

const lowerCaseA = 97
const lowerCaseZ = 122

const upperCaseA = 65
const upperCaseZ = 90

// ShiftLetter shifts a rune (letter) by a given
// number of places
func ShiftLetter(charCode rune, shift int32) rune {
	// actual value of shift should not be larger
	// than 26 or smaler than -26
	actualShift := shift % 26

	// Handle lowercase letters (a-z)
	if charCode >= lowerCaseA &&
		charCode <= lowerCaseZ {
		shifted := charCode + actualShift
		if shifted > lowerCaseZ {
			return lowerCaseA + (shifted - lowerCaseZ) - 1
		}
		if shifted < lowerCaseA {
			return lowerCaseZ - (lowerCaseA - shifted) + 1
		}
		return shifted
	}

	// Handle uppercase letters (a-z)
	if charCode >= upperCaseA &&
		charCode <= upperCaseZ {
		shifted := charCode + actualShift
		if shifted > upperCaseZ {
			return upperCaseA + (shifted - upperCaseZ) - 1
		}
		if shifted < upperCaseA {
			return upperCaseZ - (upperCaseA - shifted) + 1
		}
		return shifted
	}

	// Keep unshifted
	return charCode
}
