package repo

import (
	"regexp"
	"strings"
)

func convertToSlug(s string) string {
	trailingWhitespaceRemoved := strings.Trim(s, " ")
	allLowerCase := strings.ToLower(trailingWhitespaceRemoved)
	midWordWhitespace := regexp.MustCompile(`\s+`)
	dashes := midWordWhitespace.ReplaceAllString(allLowerCase, "-")
	return dashes
}
