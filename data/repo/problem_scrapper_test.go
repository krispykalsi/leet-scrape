package repo

import (
	"errors"
	"github.com/ISKalsi/leet-scrape/v2/data/model"
	internalErr "github.com/ISKalsi/leet-scrape/v2/internal/errors"
	"github.com/ISKalsi/leet-scrape/v2/internal/mock"
	"github.com/ISKalsi/leet-scrape/v2/internal/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetByName(group *testing.T) {
	group.Run("should return valid Question using it's name", func(t *testing.T) {
		testName := "Two sum"
		testQues, _ := testdata.ImportFromFile("two_sum.json")
		testQuery := &model.QuestionQuery{
			Question: testQues,
		}

		mockApi := &mock.GraphQLApi{}
		mockApi.On("FetchBySlug", testQues.TitleSlug).Return(testQuery, nil)
		scrapper := NewProblemScrapper(mockApi)

		actualQues, err := scrapper.GetByName(testName)

		assert.Nil(t, err)
		assert.Equal(t, &testQues, actualQues)
	})

	group.Run("should return no Question (nil) in case of an error from api", func(tt *testing.T) {
		testName := "Twwo sum"
		testSlug := "twwo-sum"

		tt.Run("when no question and random error is returned by the api", func(t *testing.T) {
			mockApi := &mock.GraphQLApi{}
			mockApi.On("FetchBySlug", testSlug).Return(nil, errors.New("random error"))
			scrapper := NewProblemScrapper(mockApi)

			ques, err := scrapper.GetByName(testName)

			assert.NotNil(t, err)
			assert.Nil(t, ques)
		})

		tt.Run("when a dummy question and random error is returned by the api", func(t *testing.T) {
			mockApi := &mock.GraphQLApi{}
			mockApi.On("FetchBySlug", testSlug).Return(&model.QuestionQuery{}, errors.New("random error"))
			scrapper := NewProblemScrapper(mockApi)

			ques, err := scrapper.GetByName(testName)

			assert.NotNil(t, err)
			assert.Nil(t, ques)
		})

		tt.Run("when no question and \"query does not exist\" error is returned by the api", func(t *testing.T) {
			mockApi := &mock.GraphQLApi{}
			mockApi.On("FetchBySlug", testSlug).Return(nil, errors.New("query does not exist"))
			scrapper := NewProblemScrapper(mockApi)

			ques, err := scrapper.GetByName(testName)

			assert.NotNil(t, err)
			assert.Nil(t, ques)
		})
	})

	group.Run("should return QuestionNotFound error when the query does not exist in the api", func(t *testing.T) {
		testName := "Twwo sum"
		testSlug := "twwo-sum"
		mockApi := &mock.GraphQLApi{}
		mockApi.On("FetchBySlug", testSlug).Return(nil, errors.New("query does not exist"))
		scrapper := NewProblemScrapper(mockApi)

		_, err := scrapper.GetByName(testName)

		assert.ErrorIs(t, internalErr.QuestionNotFound, err)
	})
}
