package main

import (
	"context"
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/vigneshperiasami/analytics/cron"
	"github.com/vigneshperiasami/analytics/database"
	"github.com/vigneshperiasami/analytics/environment"
	"github.com/vigneshperiasami/analytics/upstream"
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

func NewLogger() *log.Logger {
	return log.New(
		os.Stderr,
		"[Analytics] ",
		log.Ldate|log.Ltime|log.Lmicroseconds|log.LUTC|log.Lshortfile,
	)
}

var MainProgram = fx.Options(
	fx.Provide(
		NewLogger,
		RunProgram,
	),
)

func main() {
	fx.New(
		environment.Module,
		upstream.Module,
		database.Module,
		cron.Module,
		MainProgram,
		fx.NopLogger,
		fx.Invoke(
			func(c MainProgramResult, logger *log.Logger) {
				logger.Println(c)
			},
		),
	).Run()
}
