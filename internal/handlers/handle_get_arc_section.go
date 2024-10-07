package handlers

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/amledigital/arcxp-circulations/utils/httpclient"
	"github.com/gin-gonic/gin"
)

type Nested struct {
	Path  string `json:"path"`
	Query Query  `json:"query"`
}

type Must struct {
	Term  map[string]any `json:"term,omitempty"`
	Terms map[string]any `json:"terms,omitempty"`
}

type MultiMust struct {
	Terms map[string]any `json:"terms,omitempty"`
}

type Bool struct {
	Must []map[string]any `json:"must,omitempty"`
}

type Query struct {
	Bool Bool `json:"bool"`
}

type ElasticQuery struct {
	Query `json:"query"`
}

func (hr *HandlerRepo) HandleGetArcSection(c *gin.Context) {

	sectionID := c.Query("sectionID")

	website := c.Param("website")

	if website == "" {
		website = hr.App.ArcWebsite
	}

	var q ElasticQuery

	revisionTerm := make(map[string]any)

	revisionTerm["revision.published"] = true

	q.Query.Bool.Must = append(q.Query.Bool.Must, map[string]interface{}{"term": revisionTerm})

	typeTerm := make(map[string]any)

	typeTerm["type"] = "story"

	q.Query.Bool.Must = append(q.Query.Bool.Must, map[string]interface{}{"term": typeTerm})

	var ns Nested

	ns.Path = "taxonomy.sections"

	ns.Query.Bool.Must = []map[string]any{}

	sectionTerm := make(map[string]any)

	sectionTerm["terms"] = map[string]any{
		"taxonomy.sections._id": []string{sectionID},
	}

	ns.Query.Bool.Must = append(ns.Query.Bool.Must, map[string]interface{}{"terms": sectionTerm})

	nsSectionWebsite := make(map[string]any)

	nsSectionWebsite["taxonomy.sections._website"] = website

	ns.Query.Bool.Must = append(ns.Query.Bool.Must, map[string]interface{}{"term": nsSectionWebsite})

	q.Query.Bool.Must = append(q.Query.Bool.Must, map[string]interface{}{"nested": ns})

	out, err := json.Marshal(q)

	if err != nil {
		log.Fatalln(err)
	}

	client := httpclient.NewHttpClient("GET", fmt.Sprintf("%s/content/v4/search?website=%s&body=%s",
		hr.App.ArcContentBase,
		hr.App.ArcWebsite, out), hr.App.ArcAccessToken, nil)

	articles, next, err := client.FetchArticlesBySectionID(sectionID)

	c.JSON(200, map[string]interface{}{
		"articles": articles,
		"next":     next,
		"error":    err,
	})

}
