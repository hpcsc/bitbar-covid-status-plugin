package parser

import (
	"github.com/PuerkitoBio/goquery"
	"io"
	"strings"
)

type CasesByLGAResponse struct {
	LGA            string
	ConfirmedCases string
	ActiveCases    string
}

type CasesByLGA struct {
}

func NewCasesByLGA() CasesByLGA {
	return CasesByLGA{}
}

func (c CasesByLGA) ParseFrom(input io.Reader, lgas []string) ([]CasesByLGAResponse, error) {
	doc, err := goquery.NewDocumentFromReader(input)
	if err != nil {
		return nil, err
	}

	var result []CasesByLGAResponse
	doc.Find("table tbody tr").Each(func(i int, s *goquery.Selection) {
		tds := s.Find("td")
		lga := strings.TrimSpace(tds.Eq(0).Text())

		if contains(lgas, lga) {
			confirmedCases := strings.TrimSpace(tds.Eq(1).Text())
			activeCases := strings.TrimSpace(tds.Eq(2).Text())

			result = append(result, CasesByLGAResponse{
				LGA:            lga,
				ConfirmedCases: confirmedCases,
				ActiveCases:    activeCases,
			})
		}
	})

	return result, nil
}

func contains(lgas []string, current string) bool {
	for _, lga := range lgas {
		if strings.EqualFold(lga, current) {
			return true
		}
	}

	return false
}
