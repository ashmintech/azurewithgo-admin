package controller

import (
	"log"
	"net/http"

	"github.com/ashmintech/azurewithgo-admin/data"
)

func DeviceDataHistorical(w http.ResponseWriter, r *http.Request) {

	if err := tpl.ExecuteTemplate(w, "showhistoricaldata.gohtml", data.ShowHistoricalData()); err != nil {
		log.Fatalln("Not able to call the template", err)
	}
}
