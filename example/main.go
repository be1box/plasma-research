package main

import (
	"./db"
	"./db/middlewares"
	"./handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	db.Connect()

	r := gin.Default()
	r.Use(middlewares.Connect)
	r.Use(middlewares.ErrorHandler)
	r.GET("/getHistoryPart", handlers.ResponseHistory)
	r.PUT("/putMyTx", handlers.ResponseHistory)
	r.Run() // listen and serve on 0.0.0.0:8080
}
