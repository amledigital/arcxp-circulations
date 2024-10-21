package handlers

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/amledigital/arcxp-circulations/internal/models"
	"github.com/amledigital/arcxp-circulations/utils/httpclient"
	"github.com/gin-gonic/gin"
	"github.com/justinas/nosurf"
)

func (hr *HandlerRepo) HandlePutCirculateDocument(c *gin.Context) {

	t := c.GetHeader("X-CSRF-TOKEN")

	if !nosurf.VerifyToken(nosurf.Token(c.Request), t) {
		c.AbortWithStatusJSON(400, "not authorized")
		return
	}

	var cd *models.Circulation

	body, err := io.ReadAll(c.Request.Body)

	if err != nil {
		c.AbortWithStatusJSON(400, hr.ConstructEnvelope("error", err.Error()))
		return
	}

	if err = json.Unmarshal(body, &cd); err != nil {
		c.AbortWithStatusJSON(400, hr.ConstructEnvelope("error", err.Error()))
		return
	}

	client := httpclient.NewHttpClient("PUT",
		fmt.Sprintf("%s/draft/v1/%s/%s/circulation/%s",
			hr.App.ArcContentBase,
			"story",
			cd.DocumentID,
			hr.App.ArcWebsite),
		hr.App.ArcAccessToken,
		body)

	opResponse, err := client.CirculateADocument()

	if err != nil {
		c.AbortWithStatusJSON(400, hr.ConstructEnvelope("error", err.Error()))
		return
	}

	c.JSON(200, opResponse)

}
