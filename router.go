package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func ConfigureRouter() {
	if GetConfig().ENV != "develop" {
		gin.SetMode(gin.ReleaseMode)
	}
}

func CreateRouter() {
	router = gin.New()

	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowCredentials: true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Authentication", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Authentication", "Authorization", "Content-Type"},
	}))

	public := router.Group("/")
	{
		public.GET("/ping", TestPingPong)
	}

}

func RunRouter() {
	herokuPort, _ := os.LookupEnv("PORT")

	if herokuPort == "" {
		herokuPort = "80"
	}

	router.Run(":" + herokuPort)
}

func TestPingPong(c *gin.Context) {
	c.JSON(200, "pong")
}
