package main

import (
	"github.com/ISKalsi/leet-scrape/v2/api"
	"github.com/ISKalsi/leet-scrape/v2/data/repo"
	"github.com/ISKalsi/leet-scrape/v2/domain/model"
	"github.com/ISKalsi/leet-scrape/v2/domain/usecase"
	"github.com/ISKalsi/leet-scrape/v2/internal/errors"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

const (
	URL      = "url"
	NAME     = "name"
	NUMBER   = "number"
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
		Version:   "0.1.2",
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
					args := extractFlagArgs(c)
					ques, err := getQuestion(args)
					if err != nil {
						return exitCli(err)
					}
					generateFile := usecase.NewGenerateQuestionFile(ques, args.path)
					err = generateFile.Execute()
					if err != nil {
						return exitCli(err)
					}
					return nil
				},
			},
			{
				Name:    SOLUTION,
				Aliases: []string{"sol"},
				Usage:   "Download the starter-solution-snippet file in your desired language format",
				Action: func(c *cli.Context) error {
					args := extractFlagArgs(c)
					ques, err := getQuestion(args)
					if err != nil {
						return exitCli(err)
					}
					generateFile := usecase.NewGenerateSolutionFile(ques, args.path, args.boilerplate, args.lang)
					err = generateFile.Execute()
					if err != nil {
						return exitCli(err)
					}
					return nil
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    BOILERPLATE,
						Aliases: []string{"b"},
						Usage:   "Add boilerplate code to the Solution file at the top. Eg: " + CliName + " -n \"Two Sum\" sol -b \"#include<iostream>\\n using namespace std\\n\"",
					},
					&cli.StringFlag{
						Name:     LANGUAGE,
						Aliases:  []string{"l"},
						Usage:    "Add boilerplate code to the Solution file at the top. Eg: " + CliName + " -n \"Two Sum\" sol -l C++\n\t\tGenerally available options: C++, C, C#, Kotlin, Java, Python, Python3, Swift, \n\t\tGo, PHP, Racket, Rust, Ruby, JavaScript, TypeScript, Scala, ErLang, Elixir",
						Required: true,
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

func getQuestion(args *flagArgs) (*model.Question, error) {
	s := repo.NewScrapper(api.SolutionPart)

	var getProblem *usecase.GetProblem
	if args.url != "" {
		getProblem = usecase.NewGetProblemByUrl(s, args.url)
	} else if args.name != "" {
		getProblem = usecase.NewGetProblemByName(s, args.name)
	} else if args.num != -1 {
		getProblem = usecase.NewGetProblemByNumber(s, args.num)
	} else {
		return nil, errors.FlagMissing
	}

	ques, err := getProblem.Execute()
	if err != nil {
		return nil, err
	}
	return ques, nil
}
