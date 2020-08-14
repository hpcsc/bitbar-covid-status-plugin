package main

import (
	"bitbar-covid-status-plugin/config"
	"bitbar-covid-status-plugin/download"
	"fmt"
)

func main() {
	downloader := download.NewDHHSDownloader()
	configReader := config.NewJsonConfigReader("./.bitbar-covid-status-plugin.json")
	orchestrator := NewOrchestrator(downloader, configReader)

	output := orchestrator.Process()

	fmt.Println(output)
}
