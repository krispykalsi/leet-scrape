package api

import "os"

const (
	QueryFilePath = "api/query.gql"
)

func GetQuestionDataQuery() string {
	queryBytes, err := os.ReadFile(QueryFilePath)
	if err != nil {
		panic(err)
	}
	return string(queryBytes)
}
