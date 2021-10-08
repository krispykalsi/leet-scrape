package fixtures

import (
	"encoding/json"
	"github.com/ISKalsi/leet-scrape/v2/data/model"
	"github.com/ISKalsi/leet-scrape/v2/domain/entity"
	"os"
)

func ImportFromFile(filepath string) (entity.Question, error) {
	file, err := os.ReadFile(filepath)
	if err != nil {
		return entity.Question{}, err
	}
	var q map[string]model.QuestionQuery
	err = json.Unmarshal(file, &q)
	if err != nil {
		return entity.Question{}, err
	}
	return q["data"].Question, nil
}
