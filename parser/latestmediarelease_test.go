package parser

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestLatestMediaRelese_ParseFrom(t *testing.T) {
	t.Run("should return error if not able to find latest media link", func (t *testing.T) {
	    parser := NewLatestMediaRelease()

		output, err := parser.ParseFrom(strings.NewReader("some random string"))

		assert.Empty(t, output)
		assert.Error(t, err)
	})

	t.Run("should return latest media link if found", func (t *testing.T) {
		input := GetLatestMediaReleaseHtml()
		parser := NewLatestMediaRelease()

		output, err := parser.ParseFrom(strings.NewReader(input))

		assert.NoError(t, err)
		assert.Equal(t, "/coronavirus-update-victoria-14-august-2020", output)
	})
}

