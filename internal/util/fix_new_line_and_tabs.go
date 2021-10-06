package util

import (
	"regexp"
)

func FixNewlineAndTabs(s string) string {
	carriageReturn := regexp.MustCompile(`\\r`)
	newLine := regexp.MustCompile(`\\n`)
	tab := regexp.MustCompile(`\\t`)
	s = carriageReturn.ReplaceAllString(s, "\r")
	s = newLine.ReplaceAllString(s, "\n")
	s = tab.ReplaceAllString(s, "\t")
	return s
}
