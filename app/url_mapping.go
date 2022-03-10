package app

import "Prometheus_alerts_wrapper/controllers/alerts_controller"

func MapUrls() {
	router.POST("/SendAlert", alerts_controller.SendAlert)
}
