package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

const (
	URL    = "url"
	NAME   = "name"
	NUMBER = "number"
	TODAY  = "daily-challenge"

	LOCATION = "output-dir"

	QUESTION = "question"
	SOLUTION = "solution"

	BOILERPLATE = "boilerplate"
	LANGUAGE    = "lang"
)

const CliName = "leetscrape"

func main() {
	app := &cli.App{
		Name:      "Leetcode Scrapper",
		Version:   "0.2.3",
		Usage:     "Download and create the default empty solution file (with the question statement as docstring at the top) from leetcode.com",
		UsageText: "leetscrape [global options] command [command options]\n    Examples -\n\t1. " + CliName + " --name \"Two Sum\" solution --lang C++\n\t2. " + CliName + " -N 455 question",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    URL,
				Aliases: []string{"u"},
				Usage:   "Search problem by its `<" + URL + ">`.Eg: " + CliName + " -u https://leetcode.com/problems/two-sum sol",
			},
			&cli.StringFlag{
				Name:    NAME,
				Aliases: []string{"n"},
				Usage:   "Search problem by its `<" + NAME + ">`. Eg: " + CliName + " -n \"two sum\" sol\n\tNote: `<" + NAME + ">` should have double quotes if it contains any whitespace",
			},
			&cli.IntFlag{
				Name:    NUMBER,
				Value:   -1,
				Aliases: []string{"N"},
				Usage:   "Search problem by its `<" + NUMBER + ">`.Eg: " + CliName + " -N 1 sol",
			},
			&cli.BoolFlag{
				Name:    TODAY,
				Aliases: []string{"d"},
				Usage:   "Get problem of the day. Eg: " + CliName + " -d sol -l C++",
			},
			&cli.StringFlag{
				Name:    LOCATION,
				Aliases: []string{"o"},
				Usage:   "Directory `<path>` for the output file",
			},
		},
		Commands: []*cli.Command{question, solution},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
