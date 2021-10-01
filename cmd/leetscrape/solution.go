package main

import (
	"github.com/ISKalsi/leet-scrape/v2/domain/usecase"
	"github.com/urfave/cli/v2"
)

var solution = &cli.Command{
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
}
