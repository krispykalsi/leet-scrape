package repo

import "github.com/ISKalsi/leet-scrape/v2/domain/entity"

type ProblemScrapper interface {
	GetByName(name string) (*entity.Question, error)
	GetByUrl(url string) (*entity.Question, error)
	GetByNumber(num int) (*entity.Question, error)
}
