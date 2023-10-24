package main

import (
	"context"
	"fmt"
	"log"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/vigneshperiasami/analytics/cron"
	"github.com/vigneshperiasami/analytics/database"
	"go.uber.org/fx"
)

type MainProgramResult string

type RunProgParams struct {
	fx.In
	Actions  []cron.Action `group:"executors"`
	DbClient database.DbClient
	Sd       fx.Shutdowner
}

func WaitUntilDbIsLive(params RunProgParams, logger *log.Logger) {
	logger.Println("Waiting for database instance")
	for {
		db, err := params.DbClient.Open()
		if err != nil {
			logger.Fatalf("Error connecting db: %s", err)
			return
		}
		err = db.Ping()
		if err == nil {
			break
		}
		time.Sleep(1 * time.Second)

	}
	logger.Println("Database Instance Found")

	logger.Println("Executing crons")
	for _, v := range params.Actions {
		v.Execute()
	}

	// Shutdown after executing all actions
	params.Sd.Shutdown()
}

func RunProgram(
	params RunProgParams,
	logger *log.Logger,
	lc fx.Lifecycle) MainProgramResult {

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go WaitUntilDbIsLive(params, logger)
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
