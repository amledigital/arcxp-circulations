package handlers

import (
	"fmt"

	"github.com/amledigital/arcxp-circulations/utils/httpclient"
	"github.com/gin-gonic/gin"
)

func HandleGetDocumentCirculation(c *gin.Context) {
	client := httpclient.NewHttpClient("GET", fmt.Sprintf("%s", Repo.App.ArcContentBase), nil)
}
