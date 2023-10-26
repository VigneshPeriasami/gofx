package upstream

import "github.com/vigneshperiasami/analytics/models"

func (u *UpstreamClientResult) GetCompanies() ([]models.Company, error) {
	resp, err := u.Get(COMPANIES_PATH)
	if err != nil {
		return nil, err
	}
	var companies []models.Company
	ReadJsonResponse[[]models.Company](resp, &companies)
	return companies, nil
}

func (u *UpstreamClientResult) GetTransactions() ([]models.Transaction, error) {
	resp, err := u.Get(TRANSACTIONS_PATH)

	if err != nil {
		return nil, err
	}
	var transactions []models.Transaction
	ReadJsonResponse[[]models.Transaction](resp, &transactions)
	return transactions, nil
}
