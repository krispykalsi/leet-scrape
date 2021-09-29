package repo

import "github.com/ISKalsi/leet-scrape/v2/domain/model"

type ProblemScrapper interface {
	GetByName(name string) (*model.Question, error)
	GetByUrl(url string) (*model.Question, error)
	GetByNumber(num int) (*model.Question, error)
}
