package repository

import (
	"github.com/vigneshperiasami/analytics/database"
	"github.com/vigneshperiasami/analytics/models"
	"go.uber.org/fx"
)

type CompanyClient struct {
	dbClient *database.DbClient
}

func NewCompanyClient(dbClient *database.DbClient) *CompanyClient {
	return &CompanyClient{
		dbClient: dbClient,
	}
}

func (c *CompanyClient) GetCompanyTotalCount() int {
	db, _ := c.dbClient.Open()
	rows, _ := db.Query("select count(*) from Companies")

	rows.Next()
	var count int
	rows.Scan(&count)
	return count
}

// Reads all companies from the database
func (c *CompanyClient) GetAllCompanies() ([]models.Company, error) {
	db, err := c.dbClient.Open()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("select ID, Ibans, Name, Address from Companies limit 2")
	if err != nil {
		return nil, err
	}

	defer db.Close()
	defer rows.Close()

	companies := []models.Company{}
	for rows.Next() {
		var company models.Company
		rows.Scan(&company.Id, &company.Ibans, &company.Name, &company.Address)
		companies = append(companies, company)
	}
	return companies, nil
}

var Module = fx.Options(
	database.Module,
	fx.Provide(
		NewCompanyClient,
	),
)
