package main

import (
	"log"
	"net/http"

	"os"

	"github.com/gin-gonic/gin"
)

type Repo struct {
	User string
	Name string
	Path string
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")

	router.GET("/:repoUser/:repoName", func(c *gin.Context) {
		repo := &Repo{
			User: c.Params.ByName("repoUser"),
			Name: c.Params.ByName("repoName"),
			Path: "/" + c.Params.ByName("repoUser") + "/" + c.Params.ByName("repoName"),
		}

		c.HTML(http.StatusOK, "go-get.tmpl.html", repo)
	})

	router.Run(":" + port)
}
