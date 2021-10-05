package repo

import (
	"context"
	"github.com/ISKalsi/leet-scrape/v2/api"
	"github.com/ISKalsi/leet-scrape/v2/domain/model"
	"github.com/ISKalsi/leet-scrape/v2/internal/errors"
	"github.com/ISKalsi/leet-scrape/v2/internal/util"
	"github.com/gocolly/colly/v2"
	"github.com/machinebox/graphql"
	"regexp"
	"strconv"
	"strings"
)

const (
	ApiUrl = "https://leetcode.com/graphql"
)

type ProblemScrapper struct {
	problemPart api.PartOfProblem
}

func NewProblemScrapper(problemPart api.PartOfProblem) *ProblemScrapper {
	return &ProblemScrapper{problemPart}
}

func (s *ProblemScrapper) GetByName(name string) (*model.Question, error) {
	nameSlug := util.ConvertToSlug(name)
	client := graphql.NewClient(ApiUrl)
	query := api.GetQuery(s.problemPart)

	req := graphql.NewRequest(query)
	req.Var("titleSlug", nameSlug)
	req.Header.Set("Content-Type", "application/json")

	var q model.QuestionDataQuery
	err := client.Run(context.Background(), req, &q)
	if err != nil {
		if strings.Contains(err.Error(), "query does not exist") {
			return nil, errors.QuestionNotFound
		} else {
			return nil, err
		}
	}

	return &q.Question, nil
}

func (s *ProblemScrapper) GetByUrl(url string) (*model.Question, error) {
	isLeetcodeUrl, _ := regexp.MatchString(`leetcode\.com`, url)
	if isLeetcodeUrl {
		problemSetRegex := regexp.MustCompile(`.*problems/([\w-]*).*`)
		isFromProblemSet := problemSetRegex.MatchString(url)
		if isFromProblemSet {
			captureGroups := problemSetRegex.SubexpNames()
			if len(captureGroups) == 2 {
				slug := problemSetRegex.ReplaceAllString(url, "$1")
				q, err := s.GetByName(slug)
				if err != nil {
					return nil, err
				} else {
					return q, nil
				}
			} else {
				return nil, errors.InvalidURL
			}
		} else {
			return nil, errors.LoginRequired
		}
	} else {
		return nil, errors.InvalidURL
	}
}

func (s *ProblemScrapper) GetByNumber(num int) (*model.Question, error) {
	c := colly.NewCollector(
		colly.AllowedDomains("leetcode.com"),
	)

	var problemSubdirectories []string
	elementSelector := "a.h-5.truncate.hover\\:text-primary-s.dark\\:hover\\:text-dark-primary-s"
	c.OnHTML(elementSelector, func(e *colly.HTMLElement) {
		subdirectory := e.Attr("href")
		problemSubdirectories = append(problemSubdirectories, subdirectory)
	})
	err := c.Visit("https://leetcode.com/problemset/all/?search=" + strconv.Itoa(num) + "&page=1")
	if err != nil {
		return nil, err
	}
	c.Wait()

	if len(problemSubdirectories) == 0 {
		return nil, errors.QuestionNotFound
	}
	problemUrl := "leetcode.com" + problemSubdirectories[0]
	q, err := s.GetByUrl(problemUrl)
	if err != nil {
		return nil, err
	} else {
		return q, nil
	}
}
