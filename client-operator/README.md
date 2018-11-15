# Hello
There are some frameworks here:

go-gin: https://github.com/gin-gonic/gin

go-gin-swagger: https://github.com/swaggo/gin-swagger/

# DB
Right now we are using: MongoDb as temporary example.

FoundationDb for production later ;)

# Swagger


Here is examples how to comment all: https://swaggo.github.io/swaggo.io/declarative_comments_format/

How to add functions for Swagger:

``` js
// @Description Return History of blocks
// @Produce  application/json
// @Success 200 {array} responses.HistoryResponse
// @Router /getHistoryPart [get]

func ResponseHistory(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H {
      "History": "Array",
  })
}

```

1.  go get -u github.com/swaggo/swag/cmd/swag
2.  swag init
3.  Ready

