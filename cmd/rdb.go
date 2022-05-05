package cmd

import (
	"fmt"
	"github.com/a6e5h1/ent-sample/ent"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	"honnef.co/go/tools/config"
)

var RDBFlags = []cli.Flag{
	&cli.StringFlag{
		Name:    "db-host",
		Usage:   "db host",
		EnvVars: []string{"DB_HOST"},
		Value:   "http://127.0.0.1",
	},
	&cli.IntFlag{
		Name:    "db-port",
		Usage:   "db port",
		EnvVars: []string{"DB_PORT"},
		Value:   3306,
	},
	&cli.StringFlag{
		Name:     "db-name",
		Usage:    "db name",
		EnvVars:  []string{"DB_NAME"},
		Required: true,
	},
	&cli.StringFlag{
		Name:     "db-user",
		Usage:    "db user",
		EnvVars:  []string{"DB_USER"},
		Required: true,
	},
	&cli.StringFlag{
		Name:     "db-password",
		Usage:    "db password",
		EnvVars:  []string{"DB_PASSWORD"},
		Required: true,
	},
}

func RDBFxConstructor(ctx *cli.Context) func(*config.Config, *zap.Logger) *ent.Client {
	return func(cfg *config.Config, logger *zap.Logger) *ent.Client {
		dsn := (&mysql.Config{
			User:      ctx.String("db-user"),
			Passwd:    ctx.String("db-password"),
			Net:       "tcp",
			Addr:      fmt.Sprintf("%s:%d", ctx.String("db-host"), ctx.Int("db-port")),
			DBName:    ctx.String("db-name"),
			Collation: "utf8mb4_bin",
		}).FormatDSN()

		logger.Debug("rdb", zap.String("dsn", dsn))

		client, err := ent.Open("mysql", dsn)
		if err != nil {
			log.Fatalf("db open error: %+v", err)
		}

		return client
	}
}
