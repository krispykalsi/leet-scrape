package usecase

import (
	"github.com/ISKalsi/leet-scrape/v2/domain/model"
	"github.com/ISKalsi/leet-scrape/v2/domain/service"
	"github.com/ISKalsi/leet-scrape/v2/internal/util"
	"path/filepath"
	"regexp"
)

type GenerateQuestionFile struct {
	question  *model.Question
	path      string
	writer    service.FileWriter
	imgGetter service.ImageDownloader
}

func NewGenerateQuestionFile(q *model.Question, path string, w service.FileWriter, ig service.ImageDownloader) *GenerateQuestionFile {
	return &GenerateQuestionFile{
		path:      path,
		question:  q,
		writer:    w,
		imgGetter: ig,
	}
}

func (uc *GenerateQuestionFile) Execute() error {
	content := util.FixNewlineAndTabs(uc.question.Content)
	fileName := uc.question.TitleSlug + ".html"
	imgTagRegex := regexp.MustCompile(`<img alt=".*?" src="(.*?)".*?>`)
	imgNameRegex := regexp.MustCompile(`([^/]+)(/$|$)`)

	imgTags := imgTagRegex.FindAllStringSubmatch(content, -1)

	var imgData []struct {
		name string
		url  string
	}

	for _, match := range imgTags {
		url := match[1]
		name := imgNameRegex.FindStringSubmatch(url)[1]
		imgData = append(imgData, struct {
			name string
			url  string
		}{name, url})
	}

	mediaPath := filepath.Join(uc.path, "media")
	for _, data := range imgData {
		imgName := uc.question.TitleSlug + "_" + data.name
		err := uc.imgGetter.DownloadImageFromUrl(imgName, mediaPath, data.url)
		if err != nil {
			return err
		}
	}

	return uc.writer.WriteDataToFile(fileName, uc.path, content)

}
