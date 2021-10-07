package repo

import (
	"github.com/ISKalsi/leet-scrape/v2/data/datasource"
	"github.com/ISKalsi/leet-scrape/v2/domain/model"
	"github.com/ISKalsi/leet-scrape/v2/internal/errors"
	"github.com/ISKalsi/leet-scrape/v2/internal/util"
	"regexp"
	"strconv"
	"strings"
)

type ProblemScrapper struct {
	api datasource.GraphQLApi
}

func NewProblemScrapper(c datasource.GraphQLApi) *ProblemScrapper {
	return &ProblemScrapper{api: c}
}

func (s *ProblemScrapper) GetByName(name string) (*model.Question, error) {
	nameSlug := util.ConvertToSlug(name)
	response, err := s.api.FetchBySlug(nameSlug)
	if err != nil {
		if strings.Contains(err.Error(), "query does not exist") {
			return nil, errors.QuestionNotFound
		} else {
			return nil, err
		}
	}
	return &response.Question, nil
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
	numString := strconv.Itoa(num)
	response, err := s.api.FetchByNumber(numString)
	if err != nil {
		return nil, err
	}

	if response.QuestionList.TotalNum == 0 {
		return nil, errors.QuestionIdOutOfRange
	} else {
		ques := response.QuestionList.Data[0]
		if ques.Id != numString {
			return nil, errors.QuestionIdOutOfRange
		} else {
			return &ques, nil
		}
	}
}
