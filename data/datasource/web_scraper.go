package datasource

import (
	"github.com/ISKalsi/leet-scrape/v2/internal/errors"
	"github.com/gocolly/colly/v2"
)

type WebScrapper interface {
	ScrapeNameOfProblemOfTheDay() (string, error)
}

type WebScrapperImpl struct {
	collector *colly.Collector
}

func NewWebScrapperImpl(c *colly.Collector) *WebScrapperImpl {
	return &WebScrapperImpl{collector: c}
}

func (ws *WebScrapperImpl) ScrapeNameOfProblemOfTheDay() (string, error) {
	return "", errors.NotImplemented
}
