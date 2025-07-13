package main

import (
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/karchx/goQL/graph"
	"github.com/vektah/gqlparser/v2/ast"

	log "github.com/gothew/l-og"
	"github.com/karchx/goQL/config"
	"github.com/karchx/goQL/models"
	"github.com/karchx/goQL/utils"
)

const defaultPort = "8080"

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Errorf("Error config: %s", err.Error())
		return
	}

	err = utils.ConnectDB(config)
	if err != nil {
		log.Errorf("Error config db: %s", err.Error())
		return
	}

	err = models.Migrate(utils.DB)
	if err != nil {
		log.Errorf("Error migrate models: %s", err.Error())
		return
	}

	// GraphQL Server
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Infof("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
