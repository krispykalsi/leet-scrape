package main

import (
	"github.com/ISKalsi/leet-scrape/v2/internal/errors"
	"github.com/urfave/cli/v2"
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
