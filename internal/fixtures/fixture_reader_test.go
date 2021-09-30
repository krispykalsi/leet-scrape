package fixtures

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetQuestionFromFixture(t *testing.T) {
	expected := TestQuestion
	actual, err := ImportFromFile("convert_sorted_array_to_binary_search_tree.json")
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}
