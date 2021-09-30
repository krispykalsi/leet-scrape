package usecase

import (
	"github.com/ISKalsi/leet-scrape/v2/domain/model"
	"github.com/ISKalsi/leet-scrape/v2/internal/errors"
)

type GenerateQuestionFile struct {
	generateFile
}

func NewGenerateQuestionFile(q *model.Question, path string) *GenerateQuestionFile {
	return &GenerateQuestionFile{
		generateFile: generateFile{
			path:     path,
			question: q,
		},
	}
}

func (uc *GenerateQuestionFile) Execute() error {
	return errors.NotImplemented
}
