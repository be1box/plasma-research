package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"../swagger/responses"
)

type Responses struct {
	HistoryResponse responses.HistoryResponse
}
// @Summary Return History of blocks
// @Description Return History of blocks
// @Produce  application/json
// @Success 200 {array} responses.HistoryResponse
// @Router /getHistoryPart [get]
func ResponseHistory(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"History": "Array",
	})
}

func SignTx(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{
		"Signed": "true",
	})
}