package main

import (
	"github.com/ISKalsi/ltdl/v2/api"
	"github.com/ISKalsi/ltdl/v2/internal/errors"
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
)

const CliName = "leetscrape"

func main() {
	app := &cli.App{
		Name:  "Leetcode Downloader",
		Usage: "Download and create the default empty solution file (with the question statement as docstring at the top) from leetcode.com",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    URL,
				Aliases: []string{"u"},
				Usage:   "Get question by its full `<" + URL + ">`.Eg: " + CliName + " -u https://leetcode.com/problems/two-sum",
			},
			&cli.StringFlag{
				Name:    NAME,
				Aliases: []string{"n"},
				Usage:   "Get question by its `<" + NAME + ">`. The first search result is returned. Eg: " + CliName + " -n \"two sum\"\n\tNote: String should have double quotes if it contains any whitespace",
			},
			&cli.IntFlag{
				Name:    NUMBER,
				Value:   -1,
				Aliases: []string{"N"},
				Usage:   "Get question by its `<" + NUMBER + ">`.Eg: " + CliName + " -N 1",
			},
			&cli.StringFlag{
				Name:    BOILERPLATE,
				Aliases: []string{"b"},
				Usage:   "Add boilerplate code to the Solution file at the top",
			},
			&cli.StringFlag{
				Name:    LOCATION,
				Aliases: []string{"dir"},
				Usage:   "Directory path for the output solution file",
			},
		},
		Action: func(c *cli.Context) error {
			url := c.String(URL)
			num := c.Int(NUMBER)
			name := c.String(NAME)
			boilerplate := c.String(BOILERPLATE)
			path := c.String(LOCATION)

			var err error
			var question *api.Question

			if url != "" {
				question, err = getQuestionByUrl(url)
			} else if num != -1 {

			} else if name != "" {
				question, err = getQuestionByName(name)
			} else {
				handleError(errors.FlagMissing)
			}
			handleError(err)

			err = makeFileFromQuestionData(boilerplate, question, path)
			if err != nil {
				handleError(errors.FileGeneration)
				log.Println(err)
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
