package api

import "github.com/ISKalsi/leet-scrape/v2/domain/model"

type QuestionListQuery struct {
	QuestionList struct {
		TotalNum int              `json:"totalNum"`
		Data     []model.Question `json:"data"`
	} `json:"questionList"`
}
