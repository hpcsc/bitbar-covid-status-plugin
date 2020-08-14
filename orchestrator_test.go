package main

import (
	"bitbar-covid-status-plugin/config"
	"bitbar-covid-status-plugin/parser"
	"errors"
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"strings"
	"testing"
)

type stubConfigReader struct {
	lga []string
	returnsWithError error
}

func (s stubConfigReader) Read() (config.Config, error) {
	return config.Config{LGA: s.lga}, s.returnsWithError
}

func newStubConfigReader(err error, lga ...string) stubConfigReader {
	return stubConfigReader{lga: lga, returnsWithError: err}
}

type stubDownloader struct {
	casesByLGAResponse            string
	casesByLGAError               error
	latestMediaReleaseResponse    string
	latestMediaReleaseError       error
	latestVictoriaNumbersResponse string
	latestVictoriaNumbersError    error
}

func newStubDownloader() stubDownloader {
	return stubDownloader{
		casesByLGAResponse: parser.GetCasesByLGAHtml(),
		latestVictoriaNumbersResponse: parser.GetLatestVictoriaNumbersHtml(),
		latestMediaReleaseResponse: parser.GetLatestMediaReleaseHtml(),
	}
}

func (s stubDownloader) DownloadFrom(path string) (io.ReadCloser, error) {
	if path == "/media-hub-coronavirus-disease-covid-19" {
		return newReadCloser(s.latestMediaReleaseResponse), s.latestMediaReleaseError
	}

	if path == "/coronavirus" {
		return newReadCloser(s.latestVictoriaNumbersResponse), s.latestVictoriaNumbersError
	}

	return newReadCloser(s.casesByLGAResponse), s.casesByLGAError
}

func newReadCloser(input string) io.ReadCloser {
	return ioutil.NopCloser(strings.NewReader(input))
}

func TestOrchestrator_Process(t *testing.T) {
	t.Run("return error when fail to download latest victoria number", func(t *testing.T) {
		d := newStubDownloader()
		d.latestVictoriaNumbersError = errors.New("some error")
		c := newStubConfigReader(nil)
		orchestrator := NewOrchestrator(d, c)

		output := orchestrator.Process()

		assert.Contains(t, output, "Failed to get latest Victoria numbers")
	})

	t.Run("return latest Victoria numbers when successful", func(t *testing.T) {
		d := newStubDownloader()
		c := newStubConfigReader(nil)
		orchestrator := NewOrchestrator(d, c)

		output := orchestrator.Process()

		assert.Contains(t, output, "VIC: 372 cases")
	})

	t.Run("return error when not able to read config file", func(t *testing.T) {
		d := newStubDownloader()
		c := newStubConfigReader(errors.New("some config reading error"))
		orchestrator := NewOrchestrator(d, c)

		output := orchestrator.Process()

		assert.Contains(t, output, "Failed to get config")
	})

	t.Run("return one row for each LGA when successful", func(t *testing.T) {
		d := newStubDownloader()
		c := newStubConfigReader(nil, "WYNDHAM", "BRIMBANK")
		orchestrator := NewOrchestrator(d, c)

		output := orchestrator.Process()

		assert.Contains(t, output, "WYNDHAM: 1764 confirmed, 917 active")
		assert.Contains(t, output, "BRIMBANK: 1655 confirmed, 876 active")
	})

	t.Run("return link to latest media release", func(t *testing.T) {
		d := newStubDownloader()
		c := newStubConfigReader(nil, "WYNDHAM")
		orchestrator := NewOrchestrator(d, c)

		output := orchestrator.Process()

		assert.Contains(t, output, "/coronavirus-update-victoria-14-august-2020")
	})
}
