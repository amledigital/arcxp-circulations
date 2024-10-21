package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/amledigital/arcxp-circulations/internal/models"
	"github.com/gin-gonic/gin"
)

func (hr *HandlerRepo) HandleGetAllSections(c *gin.Context) {

	_, cancel := context.WithTimeout(hr.App.CTX, time.Second*15)

	defer cancel()

	sectionChan := make(chan models.QResult)
	errorChan := make(chan error)

	hr.App.WG.Add(1)

	go func() {
		defer close(sectionChan)
		defer close(errorChan)
		defer hr.App.WG.Done()
		var client = &http.Client{}

		req, err := http.NewRequest("GET", fmt.Sprintf("%s/site/v3/website/%s/section", hr.App.ArcContentBase, hr.App.ArcWebsite), nil)

		if err != nil {
			c.AbortWithStatusJSON(400, hr.ConstructEnvelope("error", err.Error()))
			return
		}

		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", hr.App.ArcAccessToken))

		resp, err := client.Do(req)

		if resp.Body != nil {
			defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)

			if err != nil {
				errorChan <- err
				return
			}

			var QResult models.QResult
			err = json.Unmarshal(body, &QResult)

			if err != nil {
				errorChan <- err
				return
			}

			sectionChan <- QResult

		}
	}()

	select {
	case err := <-errorChan:
		c.AbortWithStatusJSON(400, hr.ConstructEnvelope("error", err.Error()))
		return
	case body := <-sectionChan:
		c.JSON(200, body)
	}

}
