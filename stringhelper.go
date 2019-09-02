package strings

import (
	"regexp"
	"strings"
)

const (
	camelRegexString = `[ _.-](\w|$)`
	titleRegexString = `[a-z][A-Z]`
)

var (
	camelRegex = regexp.MustCompile(camelRegexString)
	titleRegex = regexp.MustCompile(titleRegexString)
)

// Camelize converts various cases to camelCase
// Example: kev-kev, kev_kev and kev.kev all become kevKev
func Camelize(text string) string {
	newText := camelRegex.ReplaceAllStringFunc(text, func(s string) string {
		return upper(s[1])
	})

	// ensure first letter is lower case
	return strings.ToLower(string(newText[0])) + newText[1:]
}

// Pascalize converts various cases to PascalCase
// Example: kev-kev, kev_kev and kev.kev all become KevKev
func Pascalize(text string) string {
	camelized := Camelize(text)

	// ensure first letter is upper case
	return strings.ToUpper(string(camelized[0])) + camelized[1:]
}

// Titleize converts various cases to human readable titles
// Example: kev-kev, kevKev, kev_kev and kev.kev all become Kev Kev
func Titleize(text string) string {
	pascalized := Pascalize(text)

	return titleRegex.ReplaceAllStringFunc(pascalized, func(s string) string {
		return string(s[0]) + " " + string(s[1])
	})
}

func upper(s uint8) string {
	return strings.ToUpper(string(s))
}
