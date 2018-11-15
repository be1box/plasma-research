package handlers

import (
	"../db/models"
	"../swagger/responses"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"net/http"
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

func SignTx(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Signed": "true",
	})
}

func PutTx(c *gin.Context) {
	db := c.MustGet("test").(*mgo.Database)
	tx := models.Tx{}
	tx.Data = c.Param("tx")
	err := db.C(models.CollectionExamples).Insert(tx)
	if err != nil {
		println("Mistake DB")
	}

}
