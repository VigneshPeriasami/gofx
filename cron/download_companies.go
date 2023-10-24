package cron

import (
	"log"

	"github.com/vigneshperiasami/analytics/upstream"
)

type CompaniesDownloader struct {
	client upstream.UpstreamClient
	logger *log.Logger
}

func NewCompaniesDownloader(client upstream.UpstreamClient, logger *log.Logger) Action {
	return &CompaniesDownloader{
		client: client,
		logger: logger,
	}
}

func (d *CompaniesDownloader) Execute() {
	companies, err := d.client.GetCompanies()
	if err != nil {
		d.logger.Println("Error downloading companies: ", err)
		return
	}
	d.logger.Printf("Total companies downloaded: %d\n", len(companies))
}
