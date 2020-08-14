package parser

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestCasesByLGA_ParseFrom(t *testing.T) {
	t.Run("should return one output for each input LGA", func (t *testing.T) {
		input := GetCasesByLGAHtml()
	    parser := NewCasesByLGA()
		casesByLGAs, err := parser.ParseFrom(strings.NewReader(input), []string{"Wyndham", "Brimbank"})

		assert.NoError(t, err)
		expectedResponse := []CasesByLGAResponse {
			{
				LGA:            "WYNDHAM",
				ConfirmedCases: "1764",
				ActiveCases:    "917",
			},
			{
				LGA:            "BRIMBANK",
				ConfirmedCases: "1655",
				ActiveCases:    "876",
			},
		}
		assert.Equal(t, expectedResponse, casesByLGAs)
	})
}
