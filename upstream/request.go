package upstream

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vigneshperiasami/analytics/environment"
	"go.uber.org/fx"
)

type UpstreamClient struct {
	upstreamUrl string
}

func NewUpstreamClient(config *environment.ConfigResult) *UpstreamClient {
	return &UpstreamClient{
		upstreamUrl: config.UpstreamUrl,
	}
}

func (up *UpstreamClient) Get(path string) (*http.Response, error) {
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
