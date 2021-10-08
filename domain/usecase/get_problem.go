package usecase

import (
	"github.com/ISKalsi/leet-scrape/v2/domain/entity"
	"github.com/ISKalsi/leet-scrape/v2/domain/repo"
	"github.com/ISKalsi/leet-scrape/v2/internal/errors"
)

type searchingMethod int

const (
	byName searchingMethod = iota
	byNumber
	byUrl
)

type GetProblem struct {
	repo   repo.ProblemScrapper
	search searchingMethod
	num    int
	url    string
	name   string
}

func NewGetProblemByName(scrapper repo.ProblemScrapper, name string) *GetProblem {
	return &GetProblem{
		repo:   scrapper,
		search: byName,
		name:   name,
	}
}

func NewGetProblemByNumber(scrapper repo.ProblemScrapper, num int) *GetProblem {
	return &GetProblem{
		repo:   scrapper,
		search: byNumber,
		num:    num,
	}
}

func NewGetProblemByUrl(scrapper repo.ProblemScrapper, url string) *GetProblem {
	return &GetProblem{
		repo:   scrapper,
		search: byUrl,
		url:    url,
	}
}

func (uc *GetProblem) Execute() (*entity.Question, error) {
	var question *entity.Question
	var err error
	switch uc.search {
	case byName:
		question, err = uc.repo.GetByName(uc.name)
	case byUrl:
		question, err = uc.repo.GetByUrl(uc.url)
	case byNumber:
		question, err = uc.repo.GetByNumber(uc.num)
	default:
		return nil, errors.InvalidSearchMethod
	}
	return question, err
}
