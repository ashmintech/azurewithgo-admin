package data

import (
	"errors"
)

type Customer struct {
	CustID       string `json:"customerid"`
	FName        string `json:"fname"`
	LName        string `json:"lname"`
	Address      string `json:"address"`
	Phone        string `json:"phone"`
	Email        string `json:"email"`
	SubType      string `json:"subtype"`
	Active       bool   `json:"active"`
	CreationDate string `json:"creation"`
}

var customerList = []*Customer{
	{
		CustID:       "custid1",
		FName:        "Ashish",
		LName:        "Minocha",
		Address:      "Canada",
		Phone:        "(123) 456-7890",
		Email:        "ashmintech@outlook.com",
		SubType:      "Premium",
		Active:       true,
		CreationDate: "Apr 10, 2021",
	},
	{
		CustID:       "custid2",
		FName:        "Ashish",
		LName:        "Minocha",
		Address:      "USA",
		Phone:        "(987) 654-3210",
		Email:        "ashmintech@doesnotexist.com",
		SubType:      "Standard",
		Active:       false,
		CreationDate: "Mar 2 2021",
	},
}

// Customers is a collection of customer
type Customers []*Customer

func GetCustomers() Customers {
	return customerList
}

func GetCustomerCount() int {
	return len(customerList)
}

func GetCustomer(custID string) (*Customer, bool) {

	for _, b := range customerList {
		if b.CustID == custID {
			return b, true
		}
	}
	return nil, false
}

func AddCustomer(p *Customer) (bool, error) {

	for _, b := range customerList {
		if b.CustID == p.CustID {
			return false, errors.New("cust id already exists")
		}
	}
	customerList = append(customerList, p)
	return true, nil
}