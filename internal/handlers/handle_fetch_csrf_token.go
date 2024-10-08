package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/justinas/nosurf"
)

func (hr *HandlerRepo) HandleFetchCSRFToken(c *gin.Context) {
	token := nosurf.Token(c.Request)

	c.JSON(200, map[string]string{"token": token})
}
