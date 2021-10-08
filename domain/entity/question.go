package entity

type Question struct {
	Id           string        `json:"questionFrontendId"`
	TitleSlug    string        `json:"titleSlug"`
	Content      string        `json:"content"`
	CodeSnippets []CodeSnippet `json:"codeSnippets"`
}
