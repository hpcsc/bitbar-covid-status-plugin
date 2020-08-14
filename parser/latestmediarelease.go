package parser

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"io"
)

type LatestMediaRelese struct {
}

func NewLatestMediaRelease() LatestMediaRelese {
	return LatestMediaRelese{}
}

func (l LatestMediaRelese) ParseFrom(input io.Reader) (string, error) {
	doc, err := goquery.NewDocumentFromReader(input)
	if err != nil {
		return "", err
	}

	relativeLink, exists := doc.Find(".page-content .field--name-field-dhhs-rich-text-text ul").
		First().
		Find("a").
		First().
		Attr("href")
	if !exists {
		return "", errors.New("cannot find latest media release link")
	}

	return relativeLink, nil
}
