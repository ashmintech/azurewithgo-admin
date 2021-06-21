package controller

import (
	"log"
	"net/http"
	"path"

	"github.com/ashmintech/azurewithgo-admin/data"
)

func DeviceAnomaly(w http.ResponseWriter, r *http.Request) {

	devID := path.Base(r.URL.Path)

	if _, found := existDevice(devID); !found {
		http.Redirect(w, r, "/admin/devices", http.StatusSeeOther)
		return
	}

	if err := tpl.ExecuteTemplate(w, "deviceanomaly.gohtml", data.GetAnomaly(devID)); err != nil {
		log.Fatalln("Not able to call the template", err)
	}
}

func ShowAnomaly(w http.ResponseWriter, r *http.Request) {

	if err := tpl.ExecuteTemplate(w, "showanomaly.gohtml", data.ShowAnomaly()); err != nil {
		log.Fatalln("Not able to call the template", err)
	}
}
