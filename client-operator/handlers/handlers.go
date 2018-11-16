package handlers

import (
	"../db/models"
	"../ether"
	"../swagger/responses"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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

// @Summary Return Tx by tx Id
// @Description Return Tx by tx Id
// @Accept  text/plain
// @Param   tx     path    string     true        "tx"
// @Produce  application/json
// @Success 200 {array} responses.TxResponse
// @Router /getMyTx [get]
func GetTx(c *gin.Context) {
	db := c.MustGet("test").(*mgo.Database)
	query := bson.M{"data": c.Param("tx")}
	tx := models.Tx{}
	err := db.C(models.CollectionExamples).Find(query).One(&tx)
	if err != nil {
		println("Mistake DB")
	}
	c.JSON(http.StatusOK, gin.H{
		"Tx": tx.Data,
	})
}

// @Summary Put your tx by tx Id
// @Description Put your tx by tx Id
// @Accept  text/plain
// @Param   tx     path    string     true        "tx"
// @Produce  application/json
// @Success 200 {array} responses.SuccessResponse
// @Router /SetMyTx [post]
func SetTx(c *gin.Context) {
	db := c.MustGet("test").(*mgo.Database)
	tx := models.Tx{}
	tx.Data = c.Param("tx")
	err := db.C(models.CollectionExamples).Insert(tx)
	if err != nil {
		println("Mistake DB")
		c.JSON(http.StatusOK, gin.H{
			"Status": "Error",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"Status": "Ok",
	})
}

func GetAllTx(c *gin.Context) {
	db := c.MustGet("test").(*mgo.Database)
	tx := []models.Tx{}
	err := db.C(models.CollectionExamples).Find(nil).All(&tx)
	if err != nil {
		println("Mistake DB")
	}
	c.JSON(http.StatusOK, gin.H{
		"Txs": tx,
	})
}

func GetBalance(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Balance": ether.GetTotalSupply(),
	})
}
