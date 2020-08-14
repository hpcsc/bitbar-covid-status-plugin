package download

import (
	"fmt"
	"io"
	"net/http"
)

type Downloader interface {
	DownloadFrom(path string) (io.ReadCloser, error)
	GetUrl(path string) string
}

type DHHSDownloader struct {
	baseUrl string
}

func NewDHHSDownloader() DHHSDownloader {
	return DHHSDownloader{
		baseUrl: "https://www.dhhs.vic.gov.au",
	}
}

func (d DHHSDownloader) DownloadFrom(path string) (io.ReadCloser, error) {
	res, err := http.Get(fmt.Sprintf("%s%s", d.baseUrl, path))
	if err != nil || res.StatusCode != 200 {
		return nil, err
	}

	return res.Body, nil
}

func (d DHHSDownloader) GetUrl(path string) string {
	return fmt.Sprintf("%s%s", d.baseUrl, path)
}