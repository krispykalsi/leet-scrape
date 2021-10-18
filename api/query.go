package api

import (
	"log"
	"os"
	"path/filepath"
)

const (
	queryFilesDirPath = "../api"
	GraphqlApiUrl     = "https://leetcode.com/graphql/"
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
	queryBytes, err := os.ReadFile(filepath.Join(queryFilesDirPath, queryFileName))
	if err != nil {
		panic(err)
	}
	return string(queryBytes)
}
