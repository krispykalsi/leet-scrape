package main

import (
	"github.com/ISKalsi/ltdl/v2/internal/errors"
	"github.com/urfave/cli/v2"
	"regexp"
	"strings"
)

func exitCli(err errors.Error) {
	cli.HandleExitCoder(cli.Exit(err.GetMessage(CliName), err.GetCode()))
}

func handleError(err error) {
	if err != nil {
		switch e := err.(type) {
		case errors.Error:
			exitCli(e)
		default:
			exitCli(errors.Unexpected)
		}
	}
}

func fixNewline(s string) string {
	carriageReturn := regexp.MustCompile(`\\r`)
	newLine := regexp.MustCompile(`\\n`)
	s = carriageReturn.ReplaceAllString(s, "\r")
	s = newLine.ReplaceAllString(s, "\n")
	return s
}

func convertToSlug(s string) string {
	trailingWhitespaceRemoved := strings.Trim(s, " ")
	allLowerCase := strings.ToLower(trailingWhitespaceRemoved)
	midWordWhitespace := regexp.MustCompile(`\s+`)
	dashes := midWordWhitespace.ReplaceAllString(allLowerCase, "-")
	return dashes
}
