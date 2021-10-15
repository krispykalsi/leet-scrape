package mock

import "github.com/stretchr/testify/mock"

type ImageDownloader struct {
	mock.Mock
}

func (i *ImageDownloader) DownloadImageFromUrl(fName string, path string, url string) error {
	args := i.Called(fName, path, url)
	r0 := args.Error(0)
	return r0
}
