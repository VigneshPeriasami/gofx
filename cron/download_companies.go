package cron

import (
	"fmt"
	"log"

	"github.com/vigneshperiasami/analytics/models"
	"github.com/vigneshperiasami/analytics/upstream"
)

type CompaniesDownloader struct {
	client upstream.UpstreamClient
	logger *log.Logger
}

func NewCompaniesDownloader(client *upstream.UpstreamClient, logger *log.Logger) Action {
	return &CompaniesDownloader{
		client: *client,
		logger: logger,
	}
}

func (d *CompaniesDownloader) Execute() {
	resp, err := d.client.Get(upstream.COMPANIES_PATH)
	if err != nil {
		fmt.Println(err)
	}
	var companies []models.Company
	upstream.ReadJsonResponse[[]models.Company](resp, &companies)
	d.logger.Printf("Total companies downloaded: %d\n", len(companies))
}
