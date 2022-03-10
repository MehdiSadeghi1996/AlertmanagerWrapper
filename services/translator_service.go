package services

import (
	"Prometheus_alerts_wrapper/configurations"
	"Prometheus_alerts_wrapper/domain/alerts"
	"Prometheus_alerts_wrapper/domain/messages"
	"log"
)

const (
	alertname        = "نام هشدار"
	service          = "نام سرویس"
	status           = "وضعیت"
	problemStartedAt = "زمان شروع مشکل"
	problemEndAt     = "زمان پایان مشکل"
	Description      = "توضیحات"
	colon            = " : "
	newline          = "\n"
)

func TranslateAlertToPersian(alert *alerts.Alert) *string {

	var messagetemp = &messages.MessageTemplate{
		AlertName:        alert.GroupLabels["value"],
		Service:          alert.Receiver,
		Status:           alert.Alerts[0].Status,
		ProblemStartedAt: alert.Alerts[0].StartsAt,
		ProblemEndAt:     alert.Alerts[0].EndsAt,
		Description:      alert.Alerts[0].Annotations["description"],
	}

	var d configurations.Dictionary
	conf := d.GetConf()

	log.Println(conf.Maps["mehdi"]) // test

	var body = alertname + colon + conf.Maps[messagetemp.AlertName] + newline +
		service + colon + conf.Maps[messagetemp.Service] + newline +
		status + colon + conf.Maps[messagetemp.Status] + newline +
		problemStartedAt + colon + messagetemp.ProblemStartedAt.Format("2006-01-02 15:04") +
		problemEndAt + colon + messagetemp.ProblemEndAt.Format("2006-01-02 15:04") +
		Description + colon + conf.Maps[messagetemp.Description]

	return &body

}
