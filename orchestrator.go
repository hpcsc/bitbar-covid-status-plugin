package main

import (
	"bitbar-covid-status-plugin/config"
	"bitbar-covid-status-plugin/download"
	"bitbar-covid-status-plugin/parser"
	"fmt"
	"strings"
)

type Orchestrator struct {
	downloader                  download.Downloader
	configReader                config.ConfigReader
	latestMediaReleaseParser    parser.LatestMediaRelese
	latestVictoriaNumbersParser parser.LatestVictoriaNumbers
	casesByLGAParser            parser.CasesByLGA
}

func NewOrchestrator(downloader download.Downloader, configReader config.ConfigReader) Orchestrator {
	return Orchestrator{
		downloader:                  downloader,
		configReader:                configReader,
		latestMediaReleaseParser:    parser.NewLatestMediaRelease(),
		latestVictoriaNumbersParser: parser.NewLatestVictoriaNumbers(),
		casesByLGAParser:            parser.NewCasesByLGA(),
	}
}

func (o Orchestrator) Process() string {
	latestVictoriaNumbersReader, err := o.downloader.DownloadFrom("/coronavirus")
	defer latestVictoriaNumbersReader.Close()
	if err != nil {
		return FormatTopLevelError(fmt.Sprintf("Failed to get latest Victoria numbers: %v | color=red", err))
	}

	latestVictoriaNumbers, err := o.latestVictoriaNumbersParser.ParseFrom(latestVictoriaNumbersReader)
	if err != nil {
		return FormatTopLevelError(fmt.Sprintf("Failed to parse latest Victoria numbers: %v | color=red", err))
	}

	output := []string{
		fmt.Sprintf("VIC: %s cases | color=green", latestVictoriaNumbers.NewCases),
		"---",
	}

	numbersByLGA := o.processNumbersByLGA()
	if numbersByLGA != "" {
		output = append(output, numbersByLGA)
	}

	output = append(output, fmt.Sprintf(`---
%s | color=white
---`, latestVictoriaNumbers.Updated))

	return strings.Join(output, "\n")
}

func (o Orchestrator) processNumbersByLGA() string {
	c, err := o.configReader.Read()
	if err != nil {
		return fmt.Sprintf("Failed to get config: %v | color=red", err)
	}

	if len(c.LGA) == 0 {
		return ""
	}

	latestMediaReleaseReader, err := o.downloader.DownloadFrom("/media-hub-coronavirus-disease-covid-19")
	defer latestMediaReleaseReader.Close()
	if err != nil {
		return fmt.Sprintf("Failed to get latest media link: %v | color=red", err)
	}

	latestMediaLink, err := o.latestMediaReleaseParser.ParseFrom(latestMediaReleaseReader)
	if err != nil {
		return fmt.Sprintf("Failed to parse latest media link: %v | color=red", err)
	}

	casesByLGAReader, err := o.downloader.DownloadFrom(latestMediaLink)
	defer casesByLGAReader.Close()
	if err != nil {
		return fmt.Sprintf("Failed to get cases by LGA: %v | color=red", err)
	}

	casesByLGAs, err := o.casesByLGAParser.ParseFrom(casesByLGAReader, c.LGA)
	if err != nil {
		return fmt.Sprintf("Failed to parse latest media link: %v | color=red", err)
	}

	var casesByLGAOutput []string
	for _, casesByLGA := range casesByLGAs {
		casesByLGAOutput = append(casesByLGAOutput, fmt.Sprintf("%s: %s confirmed, %s active | color=white",
			casesByLGA.LGA,
			casesByLGA.ConfirmedCases,
			casesByLGA.ActiveCases))
	}

	casesByLGAOutput = append(casesByLGAOutput, fmt.Sprintf("---\nOpen media release | href=%s", o.downloader.GetUrl(latestMediaLink)))

	return strings.Join(casesByLGAOutput, "\n")
}

func FormatTopLevelError(error string) string {
	return fmt.Sprintf(`ERROR | color=red
---
%s |color=red`, error)
}
