package cmd

import (
	"context"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/urfave/cli/v2"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

var HTTPFlags = []cli.Flag{
	&cli.IntFlag{
		Name:    "api-port",
		Aliases: []string{"p"},
		Usage:   "API Server port",
		EnvVars: []string{"API_PORT"},
		Value:   9000,
	},
}

func HTTPServerFxInvoke(ctx *cli.Context) interface{} {
	return func(lc fx.Lifecycle, e *echo.Echo, logger *zap.Logger) {
		listener, err := net.Listen("tcp", ":"+ctx.String("api-port"))
		if err != nil {
			log.Fatalf("net.Listen error: %+v", err)
		}

		server := &http.Server{
			ErrorLog: log.New(os.Stderr, "", log.Ldate|log.Lmicroseconds|log.Llongfile),
			Handler:  e,
		}
		server.RegisterOnShutdown(func() {
			_, _ = fmt.Fprintln(os.Stdout, "Shutdown server")
		})

		lc.Append(fx.Hook{
			OnStart: func(ctx context.Context) error {
				go func() {
					if err := server.Serve(listener); err != nil {
						if !errors.Is(err, http.ErrServerClosed) {
							logger.Fatal("Serve", zap.Error(err))
						}
					}
				}()
				return nil
			},
			OnStop: server.Shutdown,
		})
	}
}
