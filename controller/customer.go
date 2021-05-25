package controller

import (
	"log"
	"net/http"
	"path"

	"github.com/ashmintech/azurewithgo-admin/module4/data"
)

func Customers(w http.ResponseWriter, r *http.Request) {

	if err := tpl.ExecuteTemplate(w, "customers.gohtml", data.GetCustomers()); err != nil {
		log.Fatalln("Not able to call the template", err)
	}
}

func existCustomer(c string) (*data.Customer, bool) {
	return data.GetCustomer(c)
}

func findDevices4Customer(c string) (data.Devices) {
	return data.GetDevices4Customer(c)
}

func CustomerDetails(w http.ResponseWriter, r *http.Request) {

	custID := path.Base(r.URL.Path)
	var d data.Devices
	var c *data.Customer

	c, found := existCustomer(custID)

	if !found {
		http.Redirect(w, r, "/admin/customers", http.StatusSeeOther)
		return
	}

	d = findDevices4Customer(c.CustID)

	type sendData struct {
		Dev  data.Devices
		Cust *data.Customer
	}

	if err := tpl.ExecuteTemplate(w, "customerdetails.gohtml", sendData{d, c}); err != nil {
		log.Fatalln("Not able to call the template", err)
	}
}
