package fixtures

import (
	"encoding/json"
	"github.com/ISKalsi/leet-scrape/v2/data/model"
	"github.com/ISKalsi/leet-scrape/v2/domain/entity"
	"os"
	"path/filepath"
	"runtime"
)

func ImportFromFile(fixtureName string) (entity.Question, error) {
	_, pathToFixtures, _, _ := runtime.Caller(0)
	_ = os.Chdir(filepath.Dir(pathToFixtures))
	file, err := os.ReadFile(fixtureName)
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
