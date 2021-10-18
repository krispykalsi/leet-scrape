package datasource

import "github.com/stretchr/testify/mock"

type WebScrapper struct {
	mock.Mock
}

func (w *WebScrapper) ScrapeNameOfDailyChallenge() (string, error) {
	args := w.Called()
	r0 := args.String(0)
	r1 := args.Error(1)
	return r0, r1
}
