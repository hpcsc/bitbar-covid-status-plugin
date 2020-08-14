package parser

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestLatestVictoriaNumbers_ParseFrom(t *testing.T) {
	t.Run("should return updated date and current number", func (t *testing.T) {
		input := GetLatestVictoriaNumbersHtml()
	    parser := NewLatestVictoriaNumbers()

		output, err := parser.ParseFrom(strings.NewReader(input))
		
		assert.NoError(t, err)
		expectedOutput := LatestVictoriaNumbersResponse{
			Updated:  "Updated: 14 August 2020   04:00pm",
			NewCases: "372",
		}
		assert.Equal(t, expectedOutput, output)
	})
}

