package handlers

import "github.com/gin-gonic/gin"

func (hr *HandlerRepo) HandleGetArcSection(c *gin.Context) {

	site := c.Params.ByName("arcWebsite")

	sectionID := c.Query("sectionID")

	type payload struct {
		Website string `json:"website"`
		Section string `json:"section_id"`
	}

	c.JSON(200, payload{Website: site, Section: sectionID})

}
