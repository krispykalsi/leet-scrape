package model

import "github.com/ISKalsi/leet-scrape/v2/domain/entity"

type QuestionQuery struct {
	entity.Question `json:"question"`
}
