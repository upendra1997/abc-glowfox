package util

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func HandleError(c *gin.Context, msg string, err error) {
	slog.Error(msg, "err", err)
	c.IndentedJSON(http.StatusBadRequest, map[string]string{
		"error": msg,
	})
}
