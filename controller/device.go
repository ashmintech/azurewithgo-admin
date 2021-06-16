package controller

import (
	"log"
	"net/http"
	"path"

	"github.com/ashmintech/azurewithgo-admin/data"
)

func Devices(w http.ResponseWriter, r *http.Request) {
	if err := tpl.ExecuteTemplate(w, "devices.gohtml", data.GetAllDevices()); err != nil {
		log.Fatalln("Not able to call the template", err)
	}
}

func existDevice(d string) (*data.Device, bool) {
	return data.GetDevice(d)
}

func findCustomer4Device(d string) (*data.Customer, bool) {
	return data.GetCustomer4Device(d)
}

func DeviceToggleStatus(w http.ResponseWriter, r *http.Request) {

	devID := path.Base(r.URL.Path)

	d, ok := data.ToggleDeviceStatus(devID)

	if !ok {
		http.Redirect(w, r, "/admin/devices", http.StatusSeeOther)
		return
	} else {
		http.Redirect(w, r, "/admin/devices/"+d.DeviceID, http.StatusSeeOther)
		return

	}

}

func DeviceDetails(w http.ResponseWriter, r *http.Request) {

	devID := path.Base(r.URL.Path)
	var d *data.Device
	var c *data.Customer

	d, found := existDevice(devID)

	if !found {
		http.Redirect(w, r, "/admin/devices", http.StatusSeeOther)
		return
	}

	c, _ = findCustomer4Device(d.DeviceID)

	type sendData struct {
		Dev  *data.Device
		Cust *data.Customer
	}

	if err := tpl.ExecuteTemplate(w, "devicedetails.gohtml", sendData{d, c}); err != nil {
		log.Fatalln("Not able to call the template", err)
	}
}


