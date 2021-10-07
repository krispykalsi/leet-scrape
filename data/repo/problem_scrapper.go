package repo

import (
	"context"
	"github.com/ISKalsi/leet-scrape/v2/api"
	"github.com/ISKalsi/leet-scrape/v2/domain/model"
	"github.com/ISKalsi/leet-scrape/v2/internal/errors"
	"github.com/ISKalsi/leet-scrape/v2/internal/util"
	"github.com/machinebox/graphql"
	"regexp"
	"strconv"
	"strings"
)

type ProblemScrapper struct{}

func NewProblemScrapper() *ProblemScrapper {
	return &ProblemScrapper{}
}

func (s *ProblemScrapper) GetByName(name string) (*model.Question, error) {
	nameSlug := util.ConvertToSlug(name)
	client := graphql.NewClient(api.GraphqlApiUrl)
	query := api.GetQuery(api.Question)

	req := graphql.NewRequest(query)
	req.Var("titleSlug", nameSlug)
	req.Header.Set("Content-Type", "application/json")

	var q api.QuestionQuery
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
	client := graphql.NewClient(api.GraphqlApiUrl)
	query := api.GetQuery(api.QuestionList)

	numString := strconv.Itoa(num)

	req := graphql.NewRequest(query)
	req.Var("categorySlug", "")
	req.Var("limit", 1)
	req.Var("skip", 0)
	req.Var("filters", map[string]string{
		"searchKeywords": numString,
	})
	req.Header.Set("Content-Type", "application/json")

	var q api.QuestionListQuery
	err := client.Run(context.Background(), req, &q)
	if err != nil {
		return nil, err
	}

	if q.QuestionList.TotalNum == 0 {
		return nil, errors.QuestionIdOutOfRange
	} else {
		ques := q.QuestionList.Data[0]
		if ques.Id != numString {
			return nil, errors.QuestionIdOutOfRange
		} else {
			return &ques, nil
		}
	}
}
