package controller

import (
	"log"
	"net/http"
	"path"

	"github.com/ashmintech/azurewithgo-admin/data"
)

func Customers(w http.ResponseWriter, r *http.Request) {

	if err := tpl.ExecuteTemplate(w, "customers.gohtml", data.GetCustomers()); err != nil {
		log.Fatalln("Not able to call the template", err)
	}
}

func existCustomer(c string) (*data.Customer, bool) {
	return data.GetCustomer(c)
}

func findDevices4Customer(c string) data.Devices {
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

	d = findDevices4Customer(c.CustomerID)

	type sendData struct {
		Dev  data.Devices
		Cust *data.Customer
	}

	if err := tpl.ExecuteTemplate(w, "customerdetails.gohtml", sendData{d, c}); err != nil {
		log.Fatalln("Not able to call the template", err)
	}
}

func EditCustomer(w http.ResponseWriter, r *http.Request) {

	custID := path.Base(r.URL.Path)

	var c *data.Customer

	c, found := existCustomer(custID)

	if !found {
		http.Redirect(w, r, "/admin/customers", http.StatusSeeOther)
		return
	}

	if err := tpl.ExecuteTemplate(w, "customerprofile.gohtml", c); err != nil {
		log.Fatalln("Not able to call the template", err)
	}

}

func CustomerProfile(w http.ResponseWriter, r *http.Request) {

	custID := r.FormValue("custid")

	c, found := existCustomer(custID)

	if !found {
		http.Redirect(w, r, "/admin/customers", http.StatusSeeOther)
		return
	}

	done := data.PutCustomer(c, r.FormValue("phone"), r.FormValue("subtype"))

	if !done {
		http.Redirect(w, r, "/admin/customers", http.StatusSeeOther)
		return
	} else {
		http.Redirect(w, r, "/admin/customers/"+custID, http.StatusSeeOther)
		return
	}

}
