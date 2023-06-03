package regex

import "regexp"

// Grep is the function to be like the grep GNU tool.
func Grep(pattern, str string) bool {
	return regexp.MustCompile(pattern).MatchString(str)
}
