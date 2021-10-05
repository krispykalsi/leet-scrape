package usecase

import (
	"github.com/ISKalsi/leet-scrape/v2/domain/model"
	"github.com/ISKalsi/leet-scrape/v2/domain/service"
	"github.com/ISKalsi/leet-scrape/v2/internal/errors"
)

type GenerateQuestionFile struct {
	question  *model.Question
	path      string
	writer    service.FileWriter
	imgGetter service.ImageDownloader
}

func NewGenerateQuestionFile(q *model.Question, path string, w service.FileWriter, ig service.ImageDownloader) *GenerateQuestionFile {
	return &GenerateQuestionFile{
		path:      path,
		question:  q,
		writer:    w,
		imgGetter: ig,
	}
}

func (uc *GenerateQuestionFile) Execute() error {
	return errors.NotImplemented
}
