package main

import (
	"context"
	"fmt"
	"log"

	_ "github.com/joho/godotenv/autoload"
	"github.com/vigneshperiasami/analytics/cron"
	"go.uber.org/fx"
)

type MainProgramResult string

type RunProgParams struct {
	fx.In
	Actions []cron.Action `group:"executors"`
}

func RunProgram(
	params RunProgParams,
	logger *log.Logger,
	lc fx.Lifecycle,
	sd fx.Shutdowner) MainProgramResult {

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			for _, v := range params.Actions {
				v.Execute()
			}
			// Nothing to do after this, hence shutdown
			sd.Shutdown()
			return nil
		},

		OnStop: func(ctx context.Context) error {
			logger.Println("Done my job shutting down")
			return nil
		},
	})
	return MainProgramResult(fmt.Sprintf("Executing total tasks: %v", len(params.Actions)))
}

var MainProgram = fx.Options(
	fx.Provide(
		RunProgram,
	),
)

func main() {
	fx.New(
		Options,
		MainProgram,
		fx.NopLogger,
		fx.Invoke(
			func(c MainProgramResult, logger *log.Logger) {
				logger.Println(c)
			},
		),
	).Run()
}
