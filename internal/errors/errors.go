package errors

const (
	InvalidURL Error = iota + 100
	LoginRequired
	FlagMissing
	FileGeneration
	QuestionNotFound
	Unexpected
)

type Error int

func (e Error) Error() string {
	return e.GetMessage("")
}

func (e Error) GetCode() int {
	return int(e)
}

func (e Error) GetMessage(cli string) string {
	switch e {
	case InvalidURL:
		return "Please enter valid leetcode url"
	case LoginRequired:
		return "The part of leetcode you're trying to access requires login. Try entering the name of problem instead.  Enter \"" + cli + " help\""
	case FlagMissing:
		return "At least one flag is required. Enter \"" + cli + " help\""
	case FileGeneration:
		return "Something went wrong during the generation of solution template file"
	case QuestionNotFound:
		return "Question not found. Please check the spellings."
	case Unexpected:
		return "This was not supposed to happen. Please raise an issue on the Github repo - https://github.com/ISKalsi/leetscrape"
	default:
		return "I was not prepared for this (>.<) Please raise an issue on the Github repo - https://github.com/ISKalsi/leetscrape"
	}
}
