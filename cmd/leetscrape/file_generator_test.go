package main

import (
	"github.com/ISKalsi/ltdl/v2/api"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

func TestMakeFileFromQuestionData(t *testing.T) {
	boilerplate := "// boilerplate comment"
	q, err := api.GetQuestionFromFixture("../../api/convert_sorted_array_to_binary_search_tree.json")
	assert.Nil(t, err)
	err = makeFileFromQuestionData(boilerplate, &q, "")
	assert.Nil(t, err)

	cppFileName := "convert-sorted-array-to-binary-search-tree.cpp"
	file, err := os.ReadFile(cppFileName)
	assert.Nil(t, err)
	defer func() {
		err = os.Remove(cppFileName)
		if err != nil {
			log.Println(err)
		}
	}()

	expected := boilerplate + q.CodeSnippets[0].Code
	actual := string(file)
	assert.Equal(t, expected, actual)
}
