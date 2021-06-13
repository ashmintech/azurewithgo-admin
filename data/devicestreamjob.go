package data

import "log"

type jobtype string

const (
	Monitor  jobtype = "gomonitor"
	Anomaly  jobtype = "goanolmaly"
	RealTime jobtype = "gostream"
)

type devicejob struct {
	DeviceID       string  `json:"deviceid"`
	DeviceJobID    string  `json:"devicejobid"`
	DeviceJobName  string  `json:"devicejobname"`
	DeviceJobType  jobtype `json:"devicejobtype"`
	DeviceJobState string  `json:"devicejobstate"`
}

// DevicesJob is a collection of devicejob
type DeviceJobs []*devicejob

var DeviceJobList = []*devicejob{
	{
		DeviceID:       "testiotdevice1",
		DeviceJobID:    "devicejobid",
		DeviceJobName:  "",
		DeviceJobType:  RealTime,
		DeviceJobState: "Stopped",
	},
	{
		DeviceID:       "testiotdevice1",
		DeviceJobID:    "devicejobid",
		DeviceJobName:  "",
		DeviceJobType:  Monitor,
		DeviceJobState: "Stopped",
	},
	{
		DeviceID:       "testiotdevice1",
		DeviceJobID:    "devicejobid",
		DeviceJobName:  "",
		DeviceJobType:  Anomaly,
		DeviceJobState: "Stopped",
	},
}

func GetDeviceJobs(d string) *DeviceJobs {

	var dList DeviceJobs

	for _, devID := range DeviceJobList {
		if devID.DeviceID == d {
			dList = append(dList, devID)
			log.Println("Found Device Job:", devID.DeviceJobType)
		}
	}

	return &dList

}

/*

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

*/
