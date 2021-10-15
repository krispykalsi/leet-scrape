package repo

import "github.com/ISKalsi/leet-scrape/v2/domain/entity"

type Problem interface {
	GetByName(name string) (*entity.Question, error)
	GetByUrl(url string) (*entity.Question, error)
	GetByNumber(num int) (*entity.Question, error)
	GetProblemOfTheDay() (*entity.Question, error)
}
