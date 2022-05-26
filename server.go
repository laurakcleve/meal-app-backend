package main

import (
	"laurakcleve/meal/db"
	"laurakcleve/meal/graph"
	"laurakcleve/meal/graph/generated"

	"context"
	"log"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

const defaultPort = "8080"

func graphqlHandler() gin.HandlerFunc {
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}	
	
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}
	
	db.InitDB()
	defer db.Conn.Close(context.Background())

	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "origin")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, PUT")
	})
	
	r.OPTIONS("/", func(c *gin.Context) {
		c.JSON(200, nil)
	})
	r.POST("/query", graphqlHandler())
	r.GET("/", playgroundHandler())
	r.Run(":" + port)
}
