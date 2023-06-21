package main

import (
	"log"
	"os"

	"graph-gateway/http"
	"graph-gateway/routes"

	"github.com/gin-gonic/gin"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	server := gin.New()
	server.Use(gin.Logger())
	server.Use(gin.Recovery())

	//srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	server.GET("/", http.PlaygroundHandler())
	server.POST("/query", http.GraphQLHandler())

	//comms REST routes
	routes.RegisterUserRoutes(server.Group(""))
	// http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	// http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	server.Run(":" + port)
	// log.Fatal(http.ListenAndServe(":"+port, nil))
}
