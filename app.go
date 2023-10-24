package main

import (
	"log"
	"os"

	"github.com/vigneshperiasami/analytics/cron"
	"github.com/vigneshperiasami/analytics/environment"
	"github.com/vigneshperiasami/analytics/repository"
	"github.com/vigneshperiasami/analytics/upstream"
	"go.uber.org/fx"
)

func NewLogger() *log.Logger {
	return log.New(
		os.Stdout,
		"[Analytics] ",
		log.Ldate|log.Ltime|log.Lmicroseconds|log.LUTC|log.Lshortfile,
	)
}

var Options = fx.Options(
	environment.Module,
	upstream.Module,
	repository.Module,
	cron.Module,
	fx.Provide(NewLogger),
)
