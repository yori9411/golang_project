package main

import (
	"golangAPI/database"
	"golangAPI/src"

	_ "golangAPI/docs"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/gin-gonic/gin"
)

// @title           Swagger API
// @version         1.0
// @description     This is sample of practicing API to get data in MySQL
// @termsOfService  http://swagger.io/terms/

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      127.0.0.1:8000
// @BasePath  /v1

func main() {
	router := gin.Default()
	v1 := router.Group("/v1")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	src.AddUserRouter(v1)
	go func() {
		database.SQLDB()
		database.REDISDB()
	}()
	router.Run(":8000")
}
