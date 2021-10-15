package repo

import (
	"github.com/ISKalsi/leet-scrape/v2/domain/entity"
	"github.com/stretchr/testify/mock"
)

type Problem struct {
	mock.Mock
}

func (s *Problem) GetProblemOfTheDay() (*entity.Question, error) {
	args := s.Called()
	ques, _ := args.Get(0).(*entity.Question)
	err := args.Error(1)
	return ques, err
}

func (s *Problem) GetByName(name string) (*entity.Question, error) {
	args := s.Called(name)
	ques, _ := args.Get(0).(*entity.Question)
	err := args.Error(1)
	return ques, err
}

func (s *Problem) GetByUrl(url string) (*entity.Question, error) {
	args := s.Called(url)
	ques, _ := args.Get(0).(*entity.Question)
	err := args.Error(1)
	return ques, err
}

func (s *Problem) GetByNumber(num int) (*entity.Question, error) {
	args := s.Called(num)
	ques, _ := args.Get(0).(*entity.Question)
	err := args.Error(1)
	return ques, err
}
