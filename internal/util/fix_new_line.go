package util

import (
	"regexp"
)

func FixNewline(s string) string {
	carriageReturn := regexp.MustCompile(`\\r`)
	newLine := regexp.MustCompile(`\\n`)
	s = carriageReturn.ReplaceAllString(s, "\r")
	s = newLine.ReplaceAllString(s, "\n")
	return s
}
