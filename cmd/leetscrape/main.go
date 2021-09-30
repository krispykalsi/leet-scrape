package main

import (
	"github.com/ISKalsi/leet-scrape/v2/api"
	"github.com/ISKalsi/leet-scrape/v2/data/repo"
	"github.com/ISKalsi/leet-scrape/v2/domain/usecase"
	"github.com/ISKalsi/leet-scrape/v2/internal/errors"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

const (
	URL         = "url"
	NAME        = "name"
	NUMBER      = "number"
	BOILERPLATE = "boilerplate"
	LOCATION    = "output-dir"

	QUESTION = "question"
	SOLUTION = "solution"
)

const CliName = "leetscrape"

func main() {
	app := &cli.App{
		Name:    "Leetcode Scrapper",
		Version: "0.1.2",
		Usage:   "Download and create the default empty solution file (with the question statement as docstring at the top) from leetcode.com",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    URL,
				Aliases: []string{"u"},
				Usage:   "Search problem by its `<" + URL + ">`.Eg: " + CliName + " -u https://leetcode.com/problems/two-sum sol",
			},
			&cli.StringFlag{
				Name:    NAME,
				Aliases: []string{"n"},
				Usage:   "Search problem by its `<" + NAME + ">`. Eg: " + CliName + " -n \"two sum\" sol\n\tNote: String should have double quotes if it contains any whitespace",
			},
			&cli.IntFlag{
				Name:    NUMBER,
				Value:   -1,
				Aliases: []string{"N"},
				Usage:   "Search problem by its `<" + NUMBER + ">`.Eg: " + CliName + " -N 1 sol",
			},
			&cli.StringFlag{
				Name:    LOCATION,
				Aliases: []string{"o"},
				Usage:   "Directory `<path>` for the output file",
			},
		},
		Commands: []*cli.Command{
			{
				Name:    QUESTION,
				Aliases: []string{"ques"},
				Usage:   "Download the question statement (including images, if any) as an HTML page",
				Action: func(c *cli.Context) error {
					return errors.NotImplemented
				},
			},
			{
				Name:    SOLUTION,
				Aliases: []string{"sol"},
				Usage:   "Download the starter-solution-snippet file in your desired language format",
				Action: func(c *cli.Context) error {
					args := extractFlagArgs(c)
					s := repo.NewScrapper(api.SolutionPart)
					uc := usecase.NewMakeSolutionFileUseCase(s, args.num, args.url, args.name)
					err := uc.FromQuestionData(args.boilerplate, args.path)
					if err != nil {
						handleError(err)
					}
					return nil
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    BOILERPLATE,
						Aliases: []string{"b"},
						Usage:   "Add boilerplate code to the Solution file at the top. Eg: " + CliName + " -b \"#include<iostream>\n using namespace std\n\"",
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
