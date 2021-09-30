package usecase

import (
	"github.com/ISKalsi/leet-scrape/v2/domain/model"
	"github.com/stretchr/testify/mock"
)

type mockProblemScrapper struct {
	mock.Mock
}

func (s *mockProblemScrapper) GetByName(name string) (*model.Question, error) {
	args := s.Called(name)
	ques, _ := args.Get(0).(*model.Question)
	err := args.Error(1)
	return ques, err
}

func (s *mockProblemScrapper) GetByUrl(url string) (*model.Question, error) {
	args := s.Called(url)
	ques, _ := args.Get(0).(*model.Question)
	err := args.Error(1)
	return ques, err
}

func (s *mockProblemScrapper) GetByNumber(num int) (*model.Question, error) {
	args := s.Called(num)
	ques, _ := args.Get(0).(*model.Question)
	err := args.Error(1)
	return ques, err
}
