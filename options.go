package randomize

// Letters
// Exported variable for custom letter settings.
var Letters = DefaultLetterSetting()

// ForceRandomOrderOpt
var ForceRandomOrderOpt = false

// DefaultLetterSetting
// is the default setting for random string generation.
func DefaultLetterSetting() []rune {
	return append(UpperLetterSetting(), LowerLetterSetting()...)
}

// UpperLetterSetting
// Returns a slice of runes containing all uppercase letters.
func UpperLetterSetting() []rune {
	return []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
}

// LowerLetterSetting
// Returns a slice of runes containing all lowercase letters.
func LowerLetterSetting() []rune {
	return []rune("abcdefghijklmnopqrstuvwxyz")
}

// NumberLetterSetting
// Returns a slice of runes containing all numbers.
func NumberLetterSetting() []rune {
	return []rune("0123456789")
}

// SpecialLetterSetting
// Returns a slice of runes containing all special characters.
func SpecialLetterSetting() []rune {
	return []rune("!@#$%^&*()_+{}[]:;<>?/.,~`")
}

// SetDefaultLetterSetting
// Sets the Letter setting back to Default.
func SetDefaultLetterSetting() {
	Letters = DefaultLetterSetting()
}
