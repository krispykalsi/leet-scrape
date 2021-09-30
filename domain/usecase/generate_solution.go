package usecase

import (
	"github.com/ISKalsi/leet-scrape/v2/domain/model"
	"github.com/ISKalsi/leet-scrape/v2/internal/errors"
	"strings"
)

type GenerateSolutionFile struct {
	generateFile
	boilerplate  string
	requiredLang string
}

func NewGenerateSolutionFile(q *model.Question, path string, boilerplate string, lang string) *GenerateSolutionFile {
	return &GenerateSolutionFile{
		generateFile: generateFile{
			path:     path,
			question: q,
		},
		boilerplate:  boilerplate,
		requiredLang: lang,
	}
}

func (uc *GenerateSolutionFile) Execute() error {
	if len(uc.question.CodeSnippets) == 0 {
		return errors.NoCodeSnippetsFound
	}

	for _, snippet := range uc.question.CodeSnippets {
		if strings.ToLower(snippet.Lang) == strings.ToLower(uc.requiredLang) {
			data := uc.boilerplate + snippet.Code
			data = fixNewline(data)

			fileName := uc.question.TitleSlug + "." + fileExtensions[snippet.Lang]
			return uc.writeDataToFile(fileName, data)
		}
	}
	return errors.LanguageNotFound
}
