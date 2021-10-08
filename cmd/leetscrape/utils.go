package main

import (
	"github.com/ISKalsi/leet-scrape/v2/api"
	"github.com/ISKalsi/leet-scrape/v2/data/datasource"
	"github.com/ISKalsi/leet-scrape/v2/data/repo"
	"github.com/ISKalsi/leet-scrape/v2/domain/entity"
	"github.com/ISKalsi/leet-scrape/v2/domain/usecase"
	"github.com/ISKalsi/leet-scrape/v2/internal/errors"
	"github.com/machinebox/graphql"
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

func getQuestion(args *flagArgs) (*entity.Question, error) {
	c := graphql.NewClient(api.GraphqlApiUrl)
	a := datasource.NewGraphQLApiImpl(c)
	s := repo.NewProblemScrapper(a)

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
