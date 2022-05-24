package utility

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckLength(s string, max int) bool {
	if len(s) < max {
		return true
	}
	return false
}

func RedirectError(c *gin.Context, err interface{}) {
	c.HTML(http.StatusOK, "error.html", gin.H{
		"error": err,
		"title": "Page not found",
	})
}
