package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFixNewLine(t *testing.T) {
	testString := `hello\n world\r\n \t this is sparta!\r`
	expected := "hello\n world\r\n \t this is sparta!\r"
	actual := FixNewlineAndTabs(testString)
	assert.Equal(t, expected, actual)
}
