package cron

import (
	"log"

	"github.com/vigneshperiasami/analytics/repository"
	"github.com/vigneshperiasami/analytics/upstream"
	"go.uber.org/fx"
)

type Downloader struct {
	client upstream.UpstreamClient
	repo   *repository.CompanyClient
	logger *log.Logger
}

func NewDownloader(client upstream.UpstreamClient, repo *repository.CompanyClient, logger *log.Logger) Downloader {
	return Downloader{
		client: client,
		repo:   repo,
		logger: logger,
	}
}

var Module = fx.Options(
	fx.Provide(
		NewDownloader,
		fx.Annotate(
			NewCompaniesDownloader,
			fx.ResultTags(ACTION_TAG),
		),
		fx.Annotate(
			NewTransactionsDownloader,
			fx.ResultTags(ACTION_TAG),
		),
	),
)
