package datasource

import (
	"context"
	"github.com/ISKalsi/leet-scrape/v2/api"
	"github.com/machinebox/graphql"
)

type GraphQLApi interface {
	FetchBySlug(titleSlug string) (*api.QuestionQuery, error)
	FetchByNumber(id string) (*api.QuestionListQuery, error)
}

type GraphQLApiImpl struct {
	client *graphql.Client
}

func NewGraphQLApiImpl(c *graphql.Client) *GraphQLApiImpl {
	return &GraphQLApiImpl{client: c}
}

func (g *GraphQLApiImpl) FetchBySlug(titleSlug string) (*api.QuestionQuery, error) {
	query := api.GetQuery(api.Question)
	req := graphql.NewRequest(query)
	req.Var("titleSlug", titleSlug)
	req.Header.Set("Content-Type", "application/json")

	var response api.QuestionQuery
	err := g.client.Run(context.Background(), req, &response)
	if err != nil {
		return nil, err
	} else {
		return &response, nil
	}
}

func (g *GraphQLApiImpl) FetchByNumber(id string) (*api.QuestionListQuery, error) {
	query := api.GetQuery(api.QuestionList)
	req := graphql.NewRequest(query)
	req.Var("categorySlug", "")
	req.Var("limit", 1)
	req.Var("skip", 0)
	req.Var("filters", map[string]string{
		"searchKeywords": id,
	})
	req.Header.Set("Content-Type", "application/json")

	var response api.QuestionListQuery
	err := g.client.Run(context.Background(), req, &response)
	if err != nil {
		return nil, err
	} else {
		return &response, nil
	}
}
