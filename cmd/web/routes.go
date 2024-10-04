package main

import (
	"net/http"

	"github.com/amledigital/arcxp-circulations/internal/handlers"
	"github.com/gin-gonic/gin"
)

func routes() http.Handler {
	r := gin.Default()

	api := r.Group("/api/v1")

	api.GET("/arc-section/:arcWebsite", handlers.Repo.HandleGetArcSection)

	api.GET("/:documentID", handlers.Repo.HandleGetDocumentCirculation)

	api.GET("/healthcheck", func(c *gin.Context) {
		type healthcheck struct {
			Status  int    `json:"status"`
			Version string `json:"version"`
			Msg     string `json:"msg"`
		}

		var hc healthcheck

		hc.Status = 200
		hc.Version = app.Version
		hc.Msg = "all systems go"

		c.JSON(200, hc)

	})

	return csrfHandler(r)
}
