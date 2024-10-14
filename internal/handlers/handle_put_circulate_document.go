package handlers

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/justinas/nosurf"
)

type DocumentCirculationParams struct {
	DocumentID              string
	MigrateToList           []string
	WebsiteID               string
	WebsitePrimarySectionID string
	WebsiteURL              string
}

func (hr *HandlerRepo) HandlePutCirculateDocument(c *gin.Context) {

	t := c.GetHeader("X-CSRF-TOKEN")

	if !nosurf.VerifyToken(nosurf.Token(c.Request), t) {
		c.AbortWithStatusJSON(400, "not authorized")
		return
	}

	var params DocumentCirculationParams

	params.DocumentID = c.Query("documentID")
	params.MigrateToList = strings.Split(c.Param("sections"), ",")
	params.WebsiteID = c.Query("website_id")
	params.WebsitePrimarySectionID = c.Query("primary_section")
	params.WebsiteURL = c.Query("website_url")

	// do some quick validation

	if params.WebsiteID == "" {
		params.WebsiteID = hr.App.ArcWebsite
	}

	if params.WebsitePrimarySectionID == "" {
		params.WebsitePrimarySectionID = params.MigrateToList[0]
	}

	if params.DocumentID == "" || len(params.MigrateToList) == 0 || params.WebsiteID == "" || params.WebsitePrimarySectionID == "" || params.WebsiteURL == "" {
		c.AbortWithStatusJSON(400, map[string]any{"error": "missing a document circulation parameter"})
		return
	}

	// build circulation HandlePutCirculateDocument

	primarySection := c.Query("primary_section")
	if len(primarySection) == 0 || primarySection == "" {
		c.AbortWithStatusJSON(400, map[string]any{"error": "no primary section provided"})
		return
	}

	websiteID := c.Query("website_id")

	if len(websiteID) == 0 || websiteID == "" {
		websiteID = hr.App.ArcWebsite
	}

	websiteURL := c.Query("website_url")

	c.JSON(200, documentID)

}
