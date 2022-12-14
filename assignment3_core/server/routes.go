package server

import (
	"assignment3-core/controller"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	app := gin.Default()

	app.GET("/weather", controller.GetWeather)
	app.POST("/weather", controller.SetWeather)

	app.Run(":8081")
}
