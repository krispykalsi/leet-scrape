package api

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
)

const (
	GraphqlApiUrl = "https://leetcode.com/graphql/"
)

type Query int

const (
	Question Query = iota
	QuestionList
	DailyChallenges
)

func GetQuery(part Query) string {
	var queryFileName string
	switch part {
	case Question:
		queryFileName = "question.gql"
	case QuestionList:
		queryFileName = "question_list.gql"
	case DailyChallenges:
		queryFileName = "daily_challenges.gql"
	default:
		log.Fatal("Invalid part of problem")
	}
	_, pathToFixtures, _, _ := runtime.Caller(0)
	queryBytes, err := os.ReadFile(filepath.Join(filepath.Dir(pathToFixtures), queryFileName))
	if err != nil {
		panic(err)
	}
	return string(queryBytes)
}
