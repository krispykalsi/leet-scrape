package service

import "github.com/ISKalsi/leet-scrape/v2/internal/errors"

type ImageDownloader struct{}

func NewImageDownloader() *ImageDownloader {
	return &ImageDownloader{}
}

func (d *ImageDownloader) DownloadImageFromUrl(url string) error {
	return errors.NotImplemented
}
