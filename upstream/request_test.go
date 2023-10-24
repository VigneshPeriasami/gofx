package upstream

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vigneshperiasami/analytics/environment"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
)

func TestUpstreamFx(t *testing.T) {
	var client UpstreamClient
	app := fxtest.New(t, environment.Module, fx.Provide(NewUpstreamClient),
		fx.Invoke(func(u UpstreamClient) {
			client = u
		}),
	)
	defer app.RequireStart().RequireStop()
	assert.NotEqual(t, client, nil)
}
