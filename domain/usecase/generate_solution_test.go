package usecase

import (
	"github.com/ISKalsi/leet-scrape/v2/domain/entity"
	"github.com/ISKalsi/leet-scrape/v2/internal/errors"
	"github.com/ISKalsi/leet-scrape/v2/internal/mock"
	"github.com/ISKalsi/leet-scrape/v2/internal/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateSolutionFileUseCase(group *testing.T) {
	testQuestion := &testdata.QuestionWithImg
	testPath := "testFiles/cpp"
	testBoilerplate := "// this is a test comment\n\n"
	testLang := "C++"

	group.Run("should return valid constructor", func(t *testing.T) {
		mockWriter := &mock.FileWriter{}
		uc := NewGenerateSolutionFile(testQuestion, mockWriter, testPath, testBoilerplate, testLang)
		assert.Equal(t, testQuestion, uc.question)
		assert.Equal(t, testPath, uc.path)
		assert.Equal(t, testBoilerplate, uc.boilerplate)
		assert.Equal(t, testLang, uc.language)
	})

	group.Run("should prepend boilerplate code to the given code snippet (with newline chars fixed) before writing it to file", func(t *testing.T) {
		mw := &mock.FileWriter{}
		q := &entity.Question{
			TitleSlug: "two-sum",
			Content:   "<sample content>",
			CodeSnippets: []entity.CodeSnippet{
				{
					"C++",
					"cpp",
					"<sample\\r\\n code\\n>",
				},
			},
		}
		tb := "// sample comment \\n"
		expectedData := "// sample comment \n<sample\r\n code\n>"
		mw.On("WriteDataToFile", q.TitleSlug+".cpp", testPath, expectedData).Return(nil)
		uc := NewGenerateSolutionFile(q, mw, testPath, tb, testLang)
		err := uc.Execute()
		assert.Nil(t, err)
	})

	group.Run("error verification", func(tt *testing.T) {
		tt.Run("should return NoCodeSnippetsFound error when there are no code snippets in the given question", func(t *testing.T) {
			mw := &mock.FileWriter{}
			q := &entity.Question{
				TitleSlug:    "two-sum",
				Content:      "<question description>",
				CodeSnippets: []entity.CodeSnippet{},
			}
			uc := NewGenerateSolutionFile(q, mw, testPath, testBoilerplate, testLang)
			err := uc.Execute()
			assert.ErrorIs(t, err, errors.NoCodeSnippetsFound)
		})

		tt.Run("should return ExtensionNotFound error when there is no file extension corresponding to the given language", func(t *testing.T) {
			mw := &mock.FileWriter{}
			q := &entity.Question{
				TitleSlug: "two-sum",
				Content:   "<question description>",
				CodeSnippets: []entity.CodeSnippet{
					{
						"D",
						"d",
						"<sample code>",
					},
				},
			}
			uc := NewGenerateSolutionFile(q, mw, testPath, testBoilerplate, q.CodeSnippets[0].Lang)
			err := uc.Execute()
			assert.ErrorIs(t, err, errors.ExtensionNotFound)
		})

		tt.Run("should return SnippetNotFoundInGivenLang error when there is no code snippets corresponding to the given language", func(t *testing.T) {
			mw := &mock.FileWriter{}
			q := &entity.Question{
				TitleSlug: "two-sum",
				Content:   "<question description>",
				CodeSnippets: []entity.CodeSnippet{
					{
						"D",
						"d",
						"<sample code>",
					},
				},
			}
			uc := NewGenerateSolutionFile(q, mw, testPath, testBoilerplate, testLang)
			err := uc.Execute()
			assert.ErrorIs(t, err, errors.SnippetNotFoundInGivenLang)
		})
	})
}
