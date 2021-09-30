package usecase

import (
	"github.com/ISKalsi/leet-scrape/v2/internal/errors"
	"github.com/ISKalsi/leet-scrape/v2/internal/fixtures"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetProblemUseCase(group *testing.T) {
	testName := "Two Sum"
	testUrl := "https://leetcode.com/problems/two-sum/"
	testNum := 45

	testQuestion := &fixtures.TestQuestion

	group.Run("should return valid GetProblem use case", func(tt *testing.T) {
		tt.Run("by name", func(t *testing.T) {
			mockScrapper := mockProblemScrapper{}
			getProblem := NewGetProblemByName(&mockScrapper, testName)
			assert.Equal(t, "", getProblem.url)
			assert.Equal(t, 0, getProblem.num)
			assert.Equal(t, testName, getProblem.name)
			assert.Equal(t, byName, getProblem.search)
			mockScrapper.AssertNotCalled(t, "Execute")
		})

		tt.Run("by num", func(t *testing.T) {
			mockScrapper := mockProblemScrapper{}
			getProblem := NewGetProblemByUrl(&mockScrapper, testUrl)
			assert.Equal(t, testUrl, getProblem.url)
			assert.Equal(t, 0, getProblem.num)
			assert.Equal(t, "", getProblem.name)
			assert.Equal(t, byUrl, getProblem.search)
			mockScrapper.AssertNotCalled(t, "Execute")
		})

		tt.Run("by url", func(t *testing.T) {
			mockScrapper := mockProblemScrapper{}
			getProblem := NewGetProblemByNumber(&mockScrapper, testNum)
			assert.Equal(t, "", getProblem.url)
			assert.Equal(t, testNum, getProblem.num)
			assert.Equal(t, "", getProblem.name)
			assert.Equal(t, byNumber, getProblem.search)
			mockScrapper.AssertNotCalled(t, "Execute")
		})
	})

	group.Run("should call the correct repository method to fetch problem data", func(tt *testing.T) {
		tt.Run("by name", func(t *testing.T) {
			mockScrapper := mockProblemScrapper{}
			mockScrapper.On("GetByName", testName).Return(testQuestion, nil)
			getProblem := NewGetProblemByName(&mockScrapper, testName)
			actualQuestion, err := getProblem.Execute()
			assert.Nil(t, err)
			assert.Equal(t, testQuestion, actualQuestion)
			mockScrapper.AssertCalled(t, "GetByName", testName)
			mockScrapper.AssertExpectations(t)
		})

		tt.Run("by url", func(t *testing.T) {
			mockScrapper := mockProblemScrapper{}
			mockScrapper.On("GetByUrl", testUrl).Return(&fixtures.TestQuestion, nil)
			getProblem := NewGetProblemByUrl(&mockScrapper, testUrl)
			actualQuestion, err := getProblem.Execute()
			assert.Nil(t, err)
			assert.Equal(t, testQuestion, actualQuestion)
			mockScrapper.AssertCalled(t, "GetByUrl", testUrl)
			mockScrapper.AssertExpectations(t)
		})

		tt.Run("by num", func(t *testing.T) {
			mockScrapper := mockProblemScrapper{}
			mockScrapper.On("GetByNumber", testNum).Return(&fixtures.TestQuestion, nil)
			getProblem := NewGetProblemByNumber(&mockScrapper, testNum)
			actualQuestion, err := getProblem.Execute()
			assert.Nil(t, err)
			assert.Equal(t, testQuestion, actualQuestion)
			mockScrapper.AssertCalled(t, "GetByNumber", testNum)
			mockScrapper.AssertExpectations(t)
		})
	})

	group.Run("should return InvalidSearchMethod error when incorrect method is provided", func(t *testing.T) {
		mockScrapper := mockProblemScrapper{}
		getProblem := NewGetProblemByNumber(&mockScrapper, testNum)
		getProblem.search = -1
		actualQuestion, err := getProblem.Execute()
		assert.Nil(t, actualQuestion)
		assert.Error(t, err)
		assert.Equal(t, errors.InvalidSearchMethod, err)
		mockScrapper.AssertNotCalled(t, "GetByNumber", testNum)
	})
}
