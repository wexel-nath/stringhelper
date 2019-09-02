package strings

import (
	"regexp"
	"strings"
)

const (
	pascalRegexString = `[ _.-](\w|$)`
	titleRegexString = `[a-z][A-Z]`
)

var (
	pascalRegex = regexp.MustCompile(pascalRegexString)
	titleRegex = regexp.MustCompile(titleRegexString)
)

// Titleize converts various cases to human readable titles.
// Example: kev-kev, kevKev, kev_kev and kev.kev all become Kev Kev
func Titleize(text string) string {
	pascalized := pascalRegex.ReplaceAllStringFunc(text, func(s string) string {
		return upper(s[1])
	})

	titled := titleRegex.ReplaceAllStringFunc(pascalized, func(s string) string {
		return string(s[0]) + " " + string(s[1])
	})

	return upper(titled[0]) + titled[1:]
}

func upper(s uint8) string {
	return strings.ToUpper(string(s))
}
