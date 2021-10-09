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
	group.Run("should return valid Question using it's name", func(tt *testing.T) {
		testQues, _ := testdata.ImportFromFile("two_sum.json")
		testQuery := &model.QuestionQuery{
			Question: testQues,
		}
		data := []string{
			"Two sum",
			"two sum",
			"TWO SUM",
			"tWo SuM",
			"two-sum",
			"TWO-sum",
		}
		for _, testName := range data {
			tt.Run(testName, func(t *testing.T) {
				mockApi := &mock.GraphQLApi{}
				mockApi.On("FetchBySlug", testQues.TitleSlug).Return(testQuery, nil)
				scrapper := NewProblemScrapper(mockApi)

				actualQues, err := scrapper.GetByName(testName)

				assert.Nil(t, err)
				assert.Equal(t, &testQues, actualQues)
			})
		}
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

func TestGetByUrl(group *testing.T) {
	group.Run("should return a valid Question using a valid url", func(tt *testing.T) {
		testQues, _ := testdata.ImportFromFile("two_sum.json")
		testQuery := &model.QuestionQuery{
			Question: testQues,
		}
		data := []string{
			"https://www.leetcode.com/problems/two-sum/",
			"https://www.leetcode.com/problems/two-sum",
			"https://leetcode.com/problems/two-sum/",
			"www.leetcode.com/problems/two-sum",
			"leetcode.com/problems/two-sum",
		}
		for _, testUrl := range data {
			tt.Run(testUrl, func(t *testing.T) {
				mockApi := &mock.GraphQLApi{}
				mockApi.On("FetchBySlug", testQues.TitleSlug).Return(testQuery, nil)
				scrapper := NewProblemScrapper(mockApi)

				actualQues, err := scrapper.GetByUrl(testUrl)

				assert.Nil(t, err)
				assert.Equal(t, &testQues, actualQues)
			})
		}
	})

	group.Run("error verification", func(tt *testing.T) {
		tt.Run("should return InvalidUrl error when domain name is not leetcode.com", func(t *testing.T) {
			testUrl := "https://codeforces.com/problemset/problem/1600/J"
			mockApi := &mock.GraphQLApi{}
			scrapper := NewProblemScrapper(mockApi)

			ques, err := scrapper.GetByUrl(testUrl)

			assert.Nil(t, ques)
			assert.ErrorIs(t, err, internalErr.InvalidURL)
		})

		tt.Run("should return LoginRequired error when problem is from a leetcode curated playlist", func(t *testing.T) {
			testUrl := "https://leetcode.com/explore/interview/card/top-interview-questions-easy/94/trees/631/"
			mockApi := &mock.GraphQLApi{}
			scrapper := NewProblemScrapper(mockApi)

			ques, err := scrapper.GetByUrl(testUrl)

			assert.Nil(t, ques)
			assert.ErrorIs(t, err, internalErr.LoginRequired)
		})

		tt.Run("should return InvalidUrl error when 2nd subdomain of url is not a valid slug string", func(t *testing.T) {
			testUrl := "https://leetcode.com/problems/pow(x,n)/"
			mockApi := &mock.GraphQLApi{}
			scrapper := NewProblemScrapper(mockApi)

			ques, err := scrapper.GetByUrl(testUrl)

			assert.Nil(t, ques)
			assert.ErrorIs(t, err, internalErr.LoginRequired)
		})
	})

	group.Run("should return no Question (nil) in case of an error from api", func(t *testing.T) {
		testUrl := "https://leetcode.com/problems/two-summ/"
		testSlug := "two-summ"
		mockApi := &mock.GraphQLApi{}
		mockApi.On("FetchBySlug", testSlug).Return(nil, errors.New("query does not exist"))
		scrapper := NewProblemScrapper(mockApi)

		ques, err := scrapper.GetByUrl(testUrl)

		assert.Nil(t, ques)
		assert.NotNil(t, err)
	})
}
