package main

import (
	"log"

	"doki.life/initialize"
	"doki.life/middles"
	"doki.life/routers"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

// var ctx context.Context

// @title API Server
// @version 1.0.0
// @description api service
// @host 127.0.0.1:8080
// @BasePath /api

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

func main() {
	// ctx = context.Background()

	// config load
	initialize.InitConfig(".")

	// gin
	server := gin.Default()
	server.Use(requestid.New())
	server.Use(middles.AddCors())

	api := server.Group("/api")
	routers.InitCommonRouter(api)
	routers.InitSwaggerRouter(server)

	//swagger url
	log.Println("swagger url:" + initialize.GetConfig().BaseUrl + "/swagger/index.html")
	log.Fatal(server.Run(":" + initialize.GetConfig().Port))
}
