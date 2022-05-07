package main

import (
	"fmt"
	"log"
	"os"

	"github.com/a6e5h1/ent-sample/pkg/core/registry"

	"github.com/a6e5h1/ent-sample/cmd"
	"github.com/a6e5h1/ent-sample/pkg/api/interfaces"
	"github.com/urfave/cli/v2"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func main() {
	Exec(os.Args)
}

func Exec(args []string) {
	app := cli.NewApp()
	app.Flags = cmd.Flags(
		cmd.RDBFlags,
		cmd.HTTPFlags,
	)
	app.Action = func(ctx *cli.Context) error {
		app := fx.New(
			fx.Provide(
				zap.NewDevelopment,
				cmd.RDBFxConstructor(ctx),
				registry.NewServices,
				interfaces.NewHandler,
			),
			fx.Invoke(
				cmd.HTTPServerFxInvoke(ctx),
			),
		)
		if err := app.Err(); err != nil {
			return err
		}

		app.Run()
		return nil
	}

	if err := app.Run(args); err != nil {
		log.Fatal(fmt.Sprintf("%+v", err))
	}
}
