package main

import (
	"github.com/ISKalsi/leet-scrape/v2/internal/errors"
	"github.com/urfave/cli/v2"
)

func handleError(err errors.Error) cli.ExitCoder {
	return cli.Exit(err.GetMessage(CliName), err.GetCode())
}

func exitCli(err error) cli.ExitCoder {
	switch e := err.(type) {
	case errors.Error:
		return handleError(e)
	case nil:
		return nil
	default:
		return handleError(errors.Unexpected)
	}
}
