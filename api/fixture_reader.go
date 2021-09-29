package api

import (
	"encoding/json"
	"os"
)

func GetQuestionFromFixture(filepath string) (Question, error) {
	file, err := os.ReadFile(filepath)
	if err != nil {
		return Question{}, err
	}
	var q map[string]QuestionDataQuery
	err = json.Unmarshal(file, &q)
	if err != nil {
		return Question{}, err
	}
	return q["data"].Question, nil
}
