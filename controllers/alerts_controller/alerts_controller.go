package alerts_controller

import (
	"Prometheus_alerts_wrapper/configurations"
	"Prometheus_alerts_wrapper/domain/alerts"
	"Prometheus_alerts_wrapper/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func SendAlert(c *gin.Context) {
	var alrt alerts.Alert
	if err := c.ShouldBindJSON(&alrt); err != nil {
		c.JSON(http.StatusBadRequest, "Bad Request")
		return
	}

	translated := services.TranslateAlertToPersian(&alrt)

	var d configurations.Dictionary

	conf := d.GetConf()

	var phoneNumbers = conf.PhoneNumbers
	for _, element := range phoneNumbers {
		req, err := http.NewRequest("GET", "http://url/singlesend", nil)
		if err != nil {
			panic(err)
		}

		q := req.URL.Query()
		q.Add("key", "key")
		q.Add("mobileNumber", element)
		q.Add("content", translated)
		req.URL.RawQuery = q.Encode()
		fmt.Println(req.URL.String())
		client := &http.Client{}
		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(body))

	}

	c.JSON(http.StatusAccepted, "alert sent to receivers!")
}
