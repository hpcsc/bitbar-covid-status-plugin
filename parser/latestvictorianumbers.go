package parser

import (
	"github.com/PuerkitoBio/goquery"
	"io"
	"strings"
)

type LatestVictoriaNumbersResponse struct {
	Updated  string
	NewCases string
}

type LatestVictoriaNumbers struct {
}

func NewLatestVictoriaNumbers() LatestVictoriaNumbers {
	return LatestVictoriaNumbers{}
}

func (c LatestVictoriaNumbers) ParseFrom(input io.Reader) (LatestVictoriaNumbersResponse, error) {
	doc, err := goquery.NewDocumentFromReader(input)
	if err != nil {
		return LatestVictoriaNumbersResponse{}, err
	}

	latestNumbersDiv := doc.Find(".covid-latest-numbers").First()
	updated := strings.TrimSpace(latestNumbersDiv.Find(".updated").Text())

	covidNumbersDiv := latestNumbersDiv.Find(".covid-number-box")
	newCases := strings.TrimSpace(covidNumbersDiv.First().Find(".numbers").Text())

	return LatestVictoriaNumbersResponse{
		Updated:  updated,
		NewCases: newCases,
	}, nil
}
