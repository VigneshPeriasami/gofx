package cron

import (
	"fmt"
	"log"

	"github.com/vigneshperiasami/analytics/database"
	"github.com/vigneshperiasami/analytics/models"
	"go.uber.org/fx"
)

type CompanyPull struct {
	clientQuery *database.QueryClient
	logger      *log.Logger
}

type CompanyCount struct {
	dbClient *database.DbClient
	logger   *log.Logger
}

func NewCompanyCount(dbClient *database.DbClient, logger *log.Logger) Action {
	return &CompanyCount{dbClient: dbClient, logger: logger}
}

func (c *CompanyCount) Execute() {
	db, _ := c.dbClient.Open()
	rows, _ := db.Query("Select count(*) from Companies")

	var count int
	rows.Next()
	rows.Scan(&count)

	c.logger.Println("Total companies in DB: ", count)
}

func NewCompanyPull(queryClient *database.QueryClient, logger *log.Logger) Action {
	return &CompanyPull{
		clientQuery: queryClient,
		logger:      logger,
	}
}

func (cp *CompanyPull) Execute() {
	queryResult, err := cp.clientQuery.Query("select ID, Ibans, Name, Address from Companies limit 2")

	if err != nil {
		fmt.Println(err)
		return
	}

	companies := database.ReadRows[models.Company](queryResult, func(f func(...any) error) models.Company {
		var company models.Company
		f(&company.Id, &company.Ibans, &company.Name, &company.Address)
		return company
	})

	cp.logger.Printf("Sample Records:\n%v\n", companies)
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
