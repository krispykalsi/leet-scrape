package usecase

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFixNewLine(t *testing.T) {
	testString := `hello\n world\r\n this is sparta!\r`
	expected := "hello\n world\r\n this is sparta!\r"
	actual := fixNewline(testString)
	assert.Equal(t, expected, actual)
}
