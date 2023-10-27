package upstream

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vigneshperiasami/analytics/models"
	"go.uber.org/fx"
)

type UpstreamClient interface {
	GetCompanies() ([]models.Company, error)
	GetTransactions() ([]models.Transaction, error)
}

type UpstreamClientResult struct {
	upstreamUrl string
}

type UpstreamParams struct {
	fx.In
	UpstreamUrl string `name:"upstreamUrl"`
}

func NewUpstreamClient(config UpstreamParams) UpstreamClient {
	return &UpstreamClientResult{
		upstreamUrl: config.UpstreamUrl,
	}
}

func (up *UpstreamClientResult) Get(path string) (*http.Response, error) {
	return http.Get(fmt.Sprintf("%s/%s", up.upstreamUrl, path))
}

func ReadJsonResponse[T any](resp *http.Response, data *T) error {
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(data); err != nil {
		return err
	}
	return nil
}

var Module = fx.Options(
	fx.Provide(NewUpstreamClient),
)
