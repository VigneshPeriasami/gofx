package cron

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
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
