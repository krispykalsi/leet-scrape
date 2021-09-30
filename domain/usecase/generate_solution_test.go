package usecase

import (
	"github.com/ISKalsi/leet-scrape/v2/internal/fixtures"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateSolutionFileUseCase(group *testing.T) {
	testQuestion := &fixtures.TestQuestion
	testPath := "testFiles/cpp"
	testBoilerplate := "// this is a test comment\n\n"
	testLang := "C++"

	group.Run("should return valid constructor", func(t *testing.T) {
		uc := NewGenerateSolutionFile(testQuestion, testPath, testBoilerplate, testLang)
		assert.Equal(t, testQuestion, uc.question)
		assert.Equal(t, testPath, uc.path)
		assert.Equal(t, testBoilerplate, uc.boilerplate)
		assert.Equal(t, testLang, uc.requiredLang)
	})
}
