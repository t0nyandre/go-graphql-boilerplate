package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/joho/godotenv"
	gql "github.com/mattdamon108/gqlmerge/lib"
	"github.com/t0nyandre/go-graphql-boilerplate/api/handler"
	"github.com/t0nyandre/go-graphql-boilerplate/api/mutation"
	"github.com/t0nyandre/go-graphql-boilerplate/api/query"
	"github.com/t0nyandre/go-graphql-boilerplate/internal/logger"
	"github.com/t0nyandre/go-graphql-boilerplate/internal/storage/postgres"
	"github.com/t0nyandre/go-graphql-boilerplate/internal/user"
)

// Resolver consists of all our queries and mutations
type Resolver struct {
	*query.Query
	*mutation.Mutation
}

func main() {
	logger := logger.NewLogger()
	defer logger.Sync()

	if os.Getenv("APP_ENV") != "production" {
		if err := godotenv.Load("./config/env/.env"); err != nil {
			logger.Fatalw("Error loading .env file",
				"path", "./config/env/.env",
				"error", err.Error(),
			)
		}
	}

	ctx := context.Background()

	db := postgres.NewPostgresConnect(logger)

	opts := []graphql.SchemaOpt{graphql.UseFieldResolvers()}
	schema := graphql.MustParseSchema(*gql.Merge(" ", "api/graphql"), &Resolver{}, opts...)

	userRepo := user.NewRepository(db)
	userService := user.NewService(userRepo, logger)
	ctx = context.WithValue(ctx, "userService", userService)

	r := chi.NewRouter()
	r.Handle("/query", handler.AddContext(ctx, &relay.Handler{Schema: schema}))

	if err := http.ListenAndServe(fmt.Sprintf("%s:%s", os.Getenv("APP_HOST"), os.Getenv("APP_PORT")), r); err != nil {
		logger.Fatalw("Could not start GraphQL server",
			"host", os.Getenv("APP_HOST"),
			"port", os.Getenv("APP_PORT"),
			"error", err.Error(),
		)
	}

}
