package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ResponseHistory(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"History": "Ok",
	})
}
