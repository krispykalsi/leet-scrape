package fixtures

import (
	"encoding/json"
	"github.com/ISKalsi/leet-scrape/v2/api"
	"github.com/ISKalsi/leet-scrape/v2/domain/model"
	"os"
)

func ImportFromFile(filepath string) (model.Question, error) {
	file, err := os.ReadFile(filepath)
	if err != nil {
		return model.Question{}, err
	}
	var q map[string]api.QuestionQuery
	err = json.Unmarshal(file, &q)
	if err != nil {
		return model.Question{}, err
	}
	return q["data"].Question, nil
}
