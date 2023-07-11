package app

import (
	"Prometheus_alerts_wrapper/controllers/alerts_controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func MapUrls() {
	router.POST("/SendAlert", alerts_controller.SendAlert)
	router.GET("/Ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}
