package main

import (
	"fmt"
	"github.com/ISKalsi/ltdl/v2/api"
	"github.com/ISKalsi/ltdl/v2/internal/errors"
	"log"
	"os"
	"path/filepath"
)

func makeFileFromQuestionData(boilerplate string, q *api.Question, fPath string) error {
	var fName string
	if q != nil {
		fName = q.TitleSlug + "." + q.CodeSnippets[0].LangSlug
	} else {
		return errors.Unexpected
	}
	f, err := os.OpenFile(filepath.Join(fPath, fName), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer func(f *os.File) {
		err = f.Close()
		if err != nil {
			log.Fatal(err)
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
