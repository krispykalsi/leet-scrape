package datasource

import (
	"github.com/ISKalsi/leet-scrape/v2/internal/errors"
	"github.com/gocolly/colly/v2"
)

type WebScrapper interface {
	ScrapeNameOfDailyChallenge() (string, error)
}

type WebScrapperImpl struct {
	collector *colly.Collector
}

func NewWebScrapperImpl(c *colly.Collector) *WebScrapperImpl {
	return &WebScrapperImpl{collector: c}
}

func (ws *WebScrapperImpl) ScrapeNameOfDailyChallenge() (string, error) {
	return "", errors.NotImplemented
}
