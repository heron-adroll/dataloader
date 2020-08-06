package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/heron.rossi/dataloader/dataloaders"
	"github.com/heron.rossi/dataloader/graph"
	"github.com/heron.rossi/dataloader/graph/generated"
)

const defaultPort = "8080"

func GinContextToContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), "GinContextKey", c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	r := gin.New()
	r.Use(GinContextToContextMiddleware())
	r.Use(dataloaders.DataLoaderMiddleware())

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	h := playground.Handler("GraphQL Playground", "/query")

	r.POST("/query", func(c *gin.Context) { srv.ServeHTTP(c.Writer, c.Request) })
	r.GET("/", func(c *gin.Context) { h.ServeHTTP(c.Writer, c.Request) })

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
	gin.Logger()
}
