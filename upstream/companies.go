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
