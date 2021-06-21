package data

import (
	"log"

	"github.com/globalsign/mgo/bson"
)

const (
	CustomerCollName = "customers"
)

type Customer struct {
	CustomerID   string `json:"customerid"`
	FName        string `json:"fname"`
	LName        string `json:"lname"`
	Address      string `json:"address"`
	Phone        string `json:"phone"`
	Email        string `json:"email"`
	SubType      string `json:"subtype"`
	Active       bool   `json:"active"`
	CreationDate string `json:"creation"`
}

var customerList = []*Customer{}

// Customers is a collection of customer
type Customers []*Customer

func GetCustomers() Customers {

	mcoll := GetCollection(CustomerCollName)

	err = mcoll.Find(nil).Iter().All(&customerList)
	if err != nil {
		log.Println("Customer: Error while querying the Collection:\n", err)
		return nil
	}

	return customerList
}

func GetCustomerCount() int {
	mcoll := GetCollection(CustomerCollName)

	n, err := mcoll.Find(nil).Count()
	if err != nil {
		log.Println("Customer: Error while querying the Collection:\n", err)
		return 0
	}

	return n
}

func GetCustomer(custID string) (*Customer, bool) {

	var cust Customer

	mcoll := GetCollection(CustomerCollName)

	err := mcoll.Find(bson.M{"customerid": custID}).One(&cust)
	if err != nil {
		log.Println("Customer: Error while querying the Collection:\n", err)
		return nil, false
	}
	return &cust, true
}
