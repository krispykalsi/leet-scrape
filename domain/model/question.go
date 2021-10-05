package model

type Question struct {
	TitleSlug    string        `json:"titleSlug"`
	Content      string        `json:"content"`
	CodeSnippets []CodeSnippet `json:"codeSnippets"`
}
