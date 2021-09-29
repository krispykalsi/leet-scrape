package api

import (
	"log"
	"os"
	"path/filepath"
)

const (
	queryFilesDirPath = "api"
)

type PartOfProblem int

const (
	QuestionPart PartOfProblem = iota
	SolutionPart
	BothParts
)

func GetQuery(part PartOfProblem) string {
	var queryFileName string
	switch part {
	case QuestionPart:
		queryFileName = "question.gql"
	case SolutionPart:
		queryFileName = "solution.gql"
	case BothParts:
		queryFileName = "both.gql"
	default:
		log.Fatal("Invalid part of problem")
	}
	queryBytes, err := os.ReadFile(filepath.Join(queryFilesDirPath, queryFileName))
	if err != nil {
		panic(err)
	}
	return string(queryBytes)
}
