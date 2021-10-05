package service

type ImageDownloader interface {
	DownloadImageFromUrl(url string) error
}
