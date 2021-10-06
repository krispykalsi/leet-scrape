package service

type ImageDownloader interface {
	DownloadImageFromUrl(fName string, path string, url string) error
}
