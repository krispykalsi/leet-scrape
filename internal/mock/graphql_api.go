package mock

import (
	"github.com/ISKalsi/leet-scrape/v2/data/model"
	"github.com/stretchr/testify/mock"
)

type GraphQLApi struct {
	mock.Mock
}

func (g *GraphQLApi) FetchBySlug(titleSlug string) (*model.QuestionQuery, error) {
	args := g.Called(titleSlug)
	r0, _ := args.Get(0).(*model.QuestionQuery)
	r1 := args.Error(1)
	return r0, r1
}

func (g *GraphQLApi) FetchByNumber(id string) (*model.QuestionListQuery, error) {
	args := g.Called(id)
	r0, _ := args.Get(0).(*model.QuestionListQuery)
	r1 := args.Error(1)
	return r0, r1
}
