package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vigneshperiasami/analytics/cron"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
)

func TestAppFx(t *testing.T) {
	t.Run("Build dependencies for App", func(t *testing.T) {
		var crons []cron.Action
		type fxParams struct {
			fx.In
			Actions []cron.Action `group:"executors"`
		}
		app := fxtest.New(t, Options, fx.Invoke(func(params fxParams) {
			crons = append(crons, params.Actions...)
		}))
		defer app.RequireStart().RequireStop()
		assert.Equal(t, 2, len(crons))
	})
}
