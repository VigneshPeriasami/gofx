package cron

import (
	"log"

	"github.com/vigneshperiasami/analytics/repository"
	"github.com/vigneshperiasami/analytics/upstream"
)

type CompaniesDownloader struct {
	client upstream.UpstreamClient
	repo   *repository.CompanyClient
	logger *log.Logger
}

func NewCompaniesDownloader(client upstream.UpstreamClient, repo *repository.CompanyClient, logger *log.Logger) Action {
	return &CompaniesDownloader{
		client: client,
		logger: logger,
		repo:   repo,
	}
}

func (d *CompaniesDownloader) Execute() {
	companies, err := d.client.GetCompanies()
	if err != nil {
		d.logger.Println("Error downloading companies: ", err)
		return
	}
	d.logger.Printf("Total companies downloaded: %d\n", len(companies))
	err = d.repo.InsertCompanies(companies)
	if err != nil {
		d.logger.Fatalf("Error inserting companies: %s", err)
	}
}
