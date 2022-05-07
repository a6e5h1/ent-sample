package interfaces

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/a6e5h1/ent-sample/pkg/api/interfaces/graphql/generated"
	"github.com/a6e5h1/ent-sample/pkg/api/interfaces/graphql/resolver"
	"github.com/a6e5h1/ent-sample/pkg/core/registry"
	"github.com/labstack/echo/v4"
)

func NewHandler(servs *registry.Services) *echo.Echo {
	e := echo.New()

	server := handler.New(generated.NewExecutableSchema(generated.Config{
		Resolvers: &resolver.Resolver{},
	}))
	server.AddTransport(transport.POST{})
	server.Use(extension.Introspection{})

	playgroundHandler := playground.Handler("GraphQL", "/graphql")

	e.POST("/graphql", func(c echo.Context) error {
		server.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	e.GET("/playground", func(c echo.Context) error {
		playgroundHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	return e
}
