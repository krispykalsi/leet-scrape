package usecase

import (
	"fmt"
	"github.com/ISKalsi/leet-scrape/v2/domain/model"
	"github.com/ISKalsi/leet-scrape/v2/domain/repo"
	"github.com/ISKalsi/leet-scrape/v2/internal/errors"
	"log"
	"os"
	"path/filepath"
)

type MakeSolutionFileUseCase struct {
	repo repo.ProblemScrapper
	num  int
	url  string
	name string
}

func NewMakeSolutionFileUseCase(scrapper repo.ProblemScrapper, num int, url string, name string) *MakeSolutionFileUseCase {
	return &MakeSolutionFileUseCase{
		repo: scrapper,
		num:  num,
		url:  url,
		name: name,
	}
}

func (uc MakeSolutionFileUseCase) fetchQuestionData() (*model.Question, error) {
	var err error
	var question *model.Question

	if uc.url != "" {
		question, err = uc.repo.GetByUrl(uc.url)
	} else if uc.num != -1 {
		question, err = uc.repo.GetByNumber(uc.num)
	} else if uc.name != "" {
		question, err = uc.repo.GetByName(uc.name)
	} else {
		return nil, errors.FlagMissing
	}
	return question, err
}

func (uc *MakeSolutionFileUseCase) FromQuestionData(boilerplate string, fPath string) error {
	q, err := uc.fetchQuestionData()
	if err != nil {
		return err
	}

	if len(q.CodeSnippets) == 0 {
		return errors.NoCodeSnippetsFound
	}
	fName := q.TitleSlug + "." + q.CodeSnippets[0].LangSlug
	f, err := os.OpenFile(filepath.Join(fPath, fName), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		err = f.Close()
		if err != nil {
			log.Println(err)
		}
	}(f)

	boilerplate = fixNewline(boilerplate)
	appendStringToFile(f, boilerplate)
	appendStringToFile(f, q.CodeSnippets[0].Code)
	fmt.Println("Created file " + fName + " successfully")
	return nil
}

func appendStringToFile(f *os.File, s string) {
	_, err := f.WriteString(s)
	if err != nil {
		log.Fatal(err)
	}
}
