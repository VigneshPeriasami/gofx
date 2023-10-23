package environment

import (
	"os"

	"go.uber.org/fx"
)

type ConfigResult struct {
	fx.Out
	DbConn      string `name:"dbconn"`
	UpstreamUrl string `name:"upstreamUrl"`
}

func NewConfig() ConfigResult {
	return ConfigResult{
		DbConn:      os.Getenv("DATABASE_ANALYTICS"),
		UpstreamUrl: os.Getenv("ANALYTICS_UPSTREAM"),
	}
}

var Module = fx.Options(
	fx.Provide(NewConfig),
)
