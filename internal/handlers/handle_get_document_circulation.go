package handlers

import (
	"fmt"

	"github.com/amledigital/arcxp-circulations/internal/models"
	"github.com/amledigital/arcxp-circulations/utils/httpclient"
	"github.com/gin-gonic/gin"
)

func (hr *HandlerRepo) HandleGetDocumentCirculation(c *gin.Context) {

	documentID := c.Param("documentID")

	circChan := make(chan *models.Circulation)
	errorChan := make(chan error)
	doneChan := make(chan bool)

	hr.App.WG.Add(1)

	go func() {
		defer hr.App.WG.Done()

		client := httpclient.NewHttpClient("GET", fmt.Sprintf("%s/draft/v1/story/%s/circulation", hr.App.ArcContentBase, documentID), hr.App.ArcAccessToken, nil)

		circulations, err := client.FetchCirculationsByID(documentID, hr.App.ArcWebsite)

		if err != nil {
			errorChan <- err
			fmt.Println(err)
			doneChan <- true
		}

		if len(circulations.Circulations) > 0 {
			for _, v := range circulations.Circulations {
				circChan <- &v
			}
			doneChan <- true
		} else {
			doneChan <- true
		}

	}()

	var circulations []*models.Circulation

	for {
		select {
		case err := <-errorChan:
			c.AbortWithStatusJSON(400, err.Error())
			return
		case circulation := <-circChan:
			fmt.Println("getting circulation")
			circulations = append(circulations, circulation)
		case <-doneChan:
			close(circChan)
			close(errorChan)
			close(doneChan)
			data := map[string]interface{}{}
			data["circulations"] = circulations
			c.JSON(200, data)
		}
	}

}
