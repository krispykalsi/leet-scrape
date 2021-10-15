package usecase

import (
	"github.com/ISKalsi/leet-scrape/v2/domain/entity"
	"github.com/ISKalsi/leet-scrape/v2/domain/service"
	"github.com/ISKalsi/leet-scrape/v2/internal/errors"
	"github.com/ISKalsi/leet-scrape/v2/internal/util"
	"strings"
)

type GenerateSolutionFile struct {
	writer      service.FileWriter
	question    *entity.Question
	fileName    string
	path        string
	boilerplate string
	language    string
}

func NewGenerateSolutionFile(q *entity.Question, w service.FileWriter, fileName string, path string, boilerplate string, lang string) *GenerateSolutionFile {
	return &GenerateSolutionFile{
		writer:      w,
		question:    q,
		fileName:    fileName,
		path:        path,
		boilerplate: boilerplate,
		language:    lang,
	}
}

func (uc *GenerateSolutionFile) Execute() error {
	if len(uc.question.CodeSnippets) == 0 {
		return errors.NoCodeSnippetsFound
	}

	for _, snippet := range uc.question.CodeSnippets {
		if strings.ToLower(snippet.Lang) == strings.ToLower(uc.language) {
			data := uc.boilerplate + snippet.Code
			data = util.FixNewlineAndTabs(data)

			ext := fileExtensions[snippet.Lang]
			if ext == "" {
				return errors.ExtensionNotFound
			}
			fileNameWithExtension := uc.fileName + "." + ext
			return uc.writer.WriteDataToFile(fileNameWithExtension, uc.path, data)
		}
	}
	return errors.SnippetNotFoundInGivenLang
}

var fileExtensions = map[string]string{
	"C++":        "cpp",
	"Java":       "java",
	"Python":     "py",
	"Python3":    "py",
	"C":          "c",
	"C#":         "cs",
	"JavaScript": "js",
	"Ruby":       "rb",
	"Swift":      "swift",
	"Go":         "golang",
	"Scala":      "scala",
	"Kotlin":     "kt",
	"Rust":       "rs",
	"PHP":        "php",
	"TypeScript": "ts",
	"Racket":     "rkt",
	"ErLang":     "erl",
	"Elixir":     "ex",
}
