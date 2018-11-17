package handlers

import (
	"../db"
	list "../ether/listener"
	"../ether"
	"../swagger/responses"
	"fmt"
	"github.com/gin-gonic/gin"
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
	data, err := db.Tx("database").Get([]byte(c.Param("tx")))
	if err != nil {
		println("Mistake DB")
	}
	c.JSON(http.StatusOK, gin.H{
		"Tx": string(data),
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
	rawTransaction := []byte(c.Param("tx"))
	txHash := ether.GetTxHash(rawTransaction)
	err := db.Tx("database").Put(txHash, rawTransaction)
	fmt.Print(err)
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
	tx, err := db.Tx("database").GetAll()
	if err != nil {
		println("Mistake DB")
	}
	c.JSON(http.StatusOK, gin.H{
		"Txs": tx,
	})
}

func GetBalance(c *gin.Context) {
	fmt.Println(1)
	c.JSON(http.StatusOK, gin.H{
		"Balance": list.Balance,
	})
}
