package repo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertToSlug(tt *testing.T) {
	testStrings := []struct {
		name         string
		normal       string
		expectedSlug string
	}{
		{"capital case", "Two Sum", "two-sum"},
		{"trialing whitespace", "  two sum   ", "two-sum"},
		{"abnormal case", "BinARy SeArCh TrEe", "binary-search-tree"},
		{"abnormal mid word spacing", " Abnormal   spacing", "abnormal-spacing"},
		{"already kebab case", "two-sum", "two-sum"},
	}
	for _, testString := range testStrings {
		tt.Run(testString.name, func(t *testing.T) {
			actualSlug := convertToSlug(testString.normal)
			assert.Equal(t, testString.expectedSlug, actualSlug)
		})
	}
}
