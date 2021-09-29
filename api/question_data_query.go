package api

type CodeSnippet struct {
	Lang     string `json:"lang"`
	LangSlug string `json:"langSlug"`
	Code     string `json:"code"`
}

type Question struct {
	TitleSlug    string        `json:"titleSlug"`
	Content      string        `json:"content"`
	CodeSnippets []CodeSnippet `json:"codeSnippets"`
}

type QuestionDataQuery struct {
	Question `json:"question"`
}
