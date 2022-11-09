package review

import (
	"regexp"
)

var emailPattern = regexp.MustCompile(`.+@.+\..+`)

// Checks if the given string looks like a generic email, like x@x.x
func IsEmail(reviewer string) bool {
	return emailPattern.MatchString(reviewer)
}
