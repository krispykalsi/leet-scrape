package fixtures

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestImportFromFile(named *testing.T) {
	named.Run("should return valid Question from json fixture", func(t *testing.T) {
		expected := TestQuestion
		actual, err := ImportFromFile("convert_sorted_array_to_binary_search_tree.json")
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})

	named.Run("should change working directory to fixtures before importing file", func(t *testing.T) {
		err := os.Chdir("/")
		assert.Nil(t, err)

		expected := TestQuestion
		actual, err := ImportFromFile("convert_sorted_array_to_binary_search_tree.json")
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})
}
