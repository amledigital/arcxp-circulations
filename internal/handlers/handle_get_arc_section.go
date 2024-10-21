package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strconv"
	"strings"

	"github.com/amledigital/arcxp-circulations/utils/httpclient"
	"github.com/gin-gonic/gin"
)

type Nested struct {
	Path  string `json:"path,omitempty"`
	Query Query  `json:"query,omitempty"`
}

type NestedParent struct {
	Nested `json:"nested,omitempty"`
}

type Term map[string]string

type Terms map[string]string

type Must interface {
	Term | Terms | Nested
}

type MustSlice []any

type Bool struct {
	Must MustSlice `json:"must,omitempty"`
}

type Query struct {
	Bool Bool `json:"bool,omitempty"`
}

func newElasticQuery() *ElasticQuery {
	es := &ElasticQuery{}

	es.Query = Query{}

	es.Query.Bool = Bool{}

	es.Query.Bool.Must = MustSlice{}

	return es
}

type ElasticQuery struct {
	Query `json:"query,omitempty"`
}

func (eq *ElasticQuery) appendToMust(v any) {
	eq.Query.Bool.Must = append(eq.Bool.Must, v)
}

func (hr *HandlerRepo) HandleGetArcSection(c *gin.Context) {

	sectionID := c.Query("sectionID")

	parsedSectionID, err := url.ParseQuery(c.Request.URL.Path)

	if err != nil {
		c.AbortWithStatusJSON(400, hr.ConstructEnvelope("error", "there was an error parsing the query"))
		return
	}

	fmt.Println(parsedSectionID)

	filter := c.Query("filter")

	var filterParams []string

	if len(filter) > 0 {
		filterParams = strings.Split(filter, ",")
	}

	website := c.Param("website")

	from, err := strconv.Atoi(c.Query("from"))

	if err != nil {
		from = 0
	}

	if website == "" {
		website = hr.App.ArcWebsite
	}

	var q = newElasticQuery()

	var revisionTerm = make(map[string]any)

	revisionTerm["revision.published"] = true

	q.appendToMust(map[string]interface{}{"term": revisionTerm})

	var typeTerm = make(map[string]any)

	typeTerm["type"] = "story"

	q.appendToMust(map[string]interface{}{"term": typeTerm})

	var n Nested

	n.Path = "taxonomy.sections"

	n.Query.Bool.Must = []any{}

	n.Query.Bool.Must = append(n.Query.Bool.Must, map[string]any{
		"terms": map[string]any{"taxonomy.sections._id": []string{sectionID}},
	})

	n.Query.Bool.Must = append(n.Query.Bool.Must, map[string]any{"term": map[string]any{"taxonomy.sections._website": website}})

	q.appendToMust(NestedParent{n})
	out, err := json.Marshal(q)

	if err != nil {
		log.Fatalln(err)
	}

	var filterString = ""

	if len(filterParams) > 0 {
		filterString = fmt.Sprintf("&_sourceInclude=%s", url.QueryEscape(strings.Join(filterParams, ",")))
	}

	client := httpclient.NewHttpClient("GET",
		fmt.Sprintf("%s/content/v4/search?website=%s&body=%s&sort=%s"+filterString,
			hr.App.ArcContentBase,
			hr.App.ArcWebsite,
			url.QueryEscape((string(out))),
			url.QueryEscape("publish_date:desc")),
		hr.App.ArcAccessToken,
		nil)

	articles, next, err := client.FetchArticlesBySectionID(from)

	c.JSON(200, map[string]interface{}{
		"articles": articles,
		"next":     next,
		"error":    err,
	})
}
