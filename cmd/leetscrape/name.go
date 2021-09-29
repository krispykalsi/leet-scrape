package main

import (
	"context"
	"github.com/ISKalsi/ltdl/v2/api"
	"github.com/ISKalsi/ltdl/v2/internal/errors"
	"github.com/machinebox/graphql"
	"strings"
)

const (
	ApiUrl = "https://leetcode.com/graphql"
)

func getQuestionByName(name string) (*api.Question, error) {
	nameSlug := convertToSlug(name)
	client := graphql.NewClient(ApiUrl)
	query := api.GetQuestionDataQuery()

	req := graphql.NewRequest(query)
	req.Var("titleSlug", nameSlug)
	req.Header.Set("Content-Type", "application/json")

	var q api.QuestionDataQuery
	err := client.Run(context.Background(), req, &q)
	if err != nil {
		if strings.Contains(err.Error(), "query does not exist") {
			return nil, errors.QuestionNotFound
		} else {
			return nil, err
		}
	}

	return &q.Question, nil
}
