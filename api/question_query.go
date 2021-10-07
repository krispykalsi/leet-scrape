package api

import "github.com/ISKalsi/leet-scrape/v2/domain/model"

type QuestionQuery struct {
	model.Question `json:"question"`
}
