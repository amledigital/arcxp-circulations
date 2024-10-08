package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/justinas/nosurf"
)

func (hr *HandlerRepo) HandlePutCirculateDocument(c *gin.Context) {

	t := c.GetHeader("X-CSRF-TOKEN")

	if !nosurf.VerifyToken(nosurf.Token(c.Request), t) {
		c.AbortWithStatusJSON(400, "not authorized")
		return
	}

	fmt.Println("hit")

	var body interface{}

	err := json.NewDecoder(c.Request.Body).Decode(&body)

	if err != nil {
		if errors.Is(err, io.EOF) {
			c.AbortWithStatusJSON(400, "no body")
			return
		}
	}
	fmt.Printf("%+v", body)
	c.JSON(200, body)

}
