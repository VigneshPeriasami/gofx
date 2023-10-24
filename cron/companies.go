package cron

import (
	"log"

	"github.com/vigneshperiasami/analytics/repository"
	"go.uber.org/fx"
)

type CompanyPull struct {
	companyClient *repository.CompanyClient
	logger        *log.Logger
}

type CompanyCount struct {
	companyClient *repository.CompanyClient
	logger        *log.Logger
}

func NewCompanyCount(client *repository.CompanyClient, logger *log.Logger) Action {
	return &CompanyCount{companyClient: client, logger: logger}
}

func (c *CompanyCount) Execute() {
	count, _ := c.companyClient.GetCompanyTotalCount()
	c.logger.Println("Total companies in DB: ", count)
}

func NewCompanyPull(companyClient *repository.CompanyClient, logger *log.Logger) Action {
	return &CompanyPull{
		companyClient: companyClient,
		logger:        logger,
	}
}

func (cp *CompanyPull) Execute() {
	companies, err := cp.companyClient.GetAllCompanies()
	if err != nil {
		cp.logger.Println(err)
		return
	}

	if len(companies) > 2 {
		cp.logger.Printf("Sample Records:\n%v\n", companies[:2])
	}
}

var Module = fx.Options(
	fx.Provide(
		fx.Annotate(
			NewCompanyPull,
			fx.ResultTags(ACTION_TAG),
		),
		fx.Annotate(
			NewCompanyCount,
			fx.ResultTags(ACTION_TAG),
		),
		fx.Annotate(
			NewCompaniesDownloader,
			fx.ResultTags(ACTION_TAG),
		),
	),
)
