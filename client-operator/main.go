package main

import (
	"./db"
	"./db/middlewares"
	"./handlers"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"os"

	_ "./docs"
)

var (
	TYPE = os.Getenv("TYPE")
)

func main() {

	// @title Swagger History API
	// @version 0.0.1
	// @description This is Plasma-Client implementation

	// @contact.name API Support
	// @contact.email nk@bankexfoundation.org, kk@bankexfoundation.org, ig@bankexfoundation.org, av@bankexfoundation.org

	// @license.name MIT
	// @license.url https://opensource.org/licenses/MIT

	if TYPE == "operator" {
		db.Connect()
		r := gin.Default()
		r.Use(middlewares.Connect)
		r.Use(middlewares.ErrorHandler)
		r.GET("/getHistoryPart", handlers.ResponseHistory)
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		r.POST("/putMyTx/:tx", handlers.SetTx)
		r.GET("/getMyTx/:tx", handlers.GetTx)
		r.GET("/getAllTx", handlers.GetAllTx)
		r.Run(":8080")
	} else if TYPE == "client" {
		db.Connect()
		r := gin.Default()
		r.Use(middlewares.Connect)
		r.Use(middlewares.ErrorHandler)
		r.GET("/getBalance", handlers.ResponseHistory)
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		r.POST("/withdraw", handlers.ResponseHistory)
		r.POST("/Deposit", handlers.ResponseHistory)
		r.Run(":3000")
	}
}