package main

import (
	"github.com/ISKalsi/leet-scrape/v2/domain/usecase"
	"github.com/urfave/cli/v2"
)

var question = &cli.Command{
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
}
