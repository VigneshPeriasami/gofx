package environment

import (
	"fmt"
	"os"

	"go.uber.org/fx"
)

type ConfigResult struct {
	fx.Out
	DbConn      string `name:"dbconn"`
	UpstreamUrl string `name:"upstreamUrl"`
}

func NewConfig() ConfigResult {
	dbconn := fmt.Sprintf(
		"%s:%s@(%s:3306)/%s",
		os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_DATABASE"),
	)
	return ConfigResult{
		DbConn:      dbconn,
		UpstreamUrl: os.Getenv("ANALYTICS_UPSTREAM"),
	}
}

var Module = fx.Options(
	fx.Provide(NewConfig),
)
