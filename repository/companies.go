package repository

import (
	"fmt"

	"github.com/vigneshperiasami/analytics/database"
	"github.com/vigneshperiasami/analytics/models"
	"github.com/vigneshperiasami/analytics/sqlhelper"
	"go.uber.org/fx"
)

const TABLE_COMPANY = "companies"

type CompanyClient struct {
	dbClient database.DbClient
}

func NewCompanyClient(dbClient database.DbClient) *CompanyClient {
	return &CompanyClient{
		dbClient: dbClient,
	}
}

func (c *CompanyClient) GetCompanyTotalCount() (int, error) {
	db, err := c.dbClient.Open()
	if err != nil {
		return -1, err
	}
	rows, err := db.Query(fmt.Sprintf("select count(*) from %s", TABLE_COMPANY))
	if err != nil {
		return -1, err
	}
	defer db.Close()
	defer rows.Close()

	rows.Next()
	var count int
	rows.Scan(&count)
	return count, nil
}

// Reads all companies from the database
func (c *CompanyClient) GetAllCompanies() ([]models.Company, error) {
	db, err := c.dbClient.Open()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query(fmt.Sprintf("select id, ibans, name, address from %s", TABLE_COMPANY))
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

func InsertCompanyColumns(company models.Company) []interface{} {
	return []interface{}{company.Id, company.Ibans, company.Name, company.Address}
}

func (c *CompanyClient) InsertCompanies(companies []models.Company) error {
	db, err := c.dbClient.Open()
	if err != nil {
		return err
	}
	defer db.Close()

	queryParams := sqlhelper.InsertQ[models.Company]{
		TableName: TABLE_COMPANY,
		Columns:   []string{"id", "ibans", "name", "address"},
		MapperFn: func(company models.Company) []interface{} {
			return []interface{}{company.Id, company.Ibans, company.Name, company.Address}
		},
	}

	query, args := queryParams.Build(companies)

	_, err = db.Query(query, args...)

	return err
}

func (c *CompanyClient) InsertTransactions(transactions []models.Transaction) error {
	db, err := c.dbClient.Open()
	if err != nil {
		return err
	}
	defer db.Close()

	queryParams := sqlhelper.InsertQ[models.Transaction]{
		TableName: "transactions",
		Columns:   []string{"id", "beneficiary", "sender", "currency", "transactionTime", "amount"},
		MapperFn: func(t models.Transaction) []interface{} {
			return []interface{}{t.Id, t.Beneficiary, t.Sender, t.Currency, t.Timestamp, t.Amount}
		},
	}
	query, args := queryParams.Build(transactions)
	_, err = db.Query(query, args...)
	return err
}

var Module = fx.Options(
	database.Module,
	fx.Provide(
		NewCompanyClient,
	),
)
