package repo

import (
	"context"
	"github.com/ISKalsi/leet-scrape/v2/api"
	"github.com/ISKalsi/leet-scrape/v2/domain/model"
	"github.com/ISKalsi/leet-scrape/v2/internal/errors"
	"github.com/machinebox/graphql"
	"regexp"
	"strings"
)

const (
	ApiUrl = "https://leetcode.com/graphql"
)

type ProblemScrapper struct {
	problemPart api.PartOfProblem
}

func NewScrapper(problemPart api.PartOfProblem) *ProblemScrapper {
	return &ProblemScrapper{problemPart}
}

func (s *ProblemScrapper) GetByName(name string) (*model.Question, error) {
	nameSlug := convertToSlug(name)
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

func (s *ProblemScrapper) GetByNumber(_ int) (*model.Question, error) {
	return nil, errors.NotImplemented
}
