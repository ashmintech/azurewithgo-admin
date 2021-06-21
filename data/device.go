package data

import (
	"log"

	"github.com/globalsign/mgo/bson"
)

const (
	DeviceCollName = "devices"
)

// Devices is a collection of device
type Devices []*Device

type Device struct {
	DeviceID           string `json:"deviceid"`
	DeviceName         string `json:"devicename"`
	DeviceModel        string `json:"devicemodel"`
	DeviceType         string `json:"devicetype"`
	DeviceStatus       string `json:"devicestatus"`
	DeviceCreationDate string `json:"devicecreationdate"`
}

var DeviceList = []*Device{}

/*
var DeviceList = []*Device{
	{
		DeviceID:           "smartdevice1",
		DeviceName:         "Device Name 1",
		DeviceModel:        "Globomantics Fridge",
		DeviceType:         "Fridge",
		DeviceStatus:       "Active",
		DeviceCreationDate: "May 4, 2021",
	},

	{
		DeviceID:           "smartdevice2",
		DeviceName:         "Device Name 2",
		DeviceModel:        "Globomantics Fridge",
		DeviceType:         "Fridge",
		DeviceStatus:       "Inactive",
		DeviceCreationDate: "May 14, 2021",
	},
}
*/

func GetAllDevices() Devices {

	mcoll := GetCollection(DeviceCollName)

	err = mcoll.Find(nil).Iter().All(&DeviceList)
	if err != nil {
		log.Println("Error while querying the Collection:\n", err)
		return nil
	}

	return DeviceList
}

func GetDeviceCount() int {

	mcoll := GetCollection(DeviceCollName)

	n, err := mcoll.Find(nil).Count()
	if err != nil {
		log.Println("Error while querying the Collection:\n", err)
		return 0
	}

	return n
}

func GetDevice(d string) (*Device, bool) {

	var dev Device

	mcoll := GetCollection(DeviceCollName)

	err := mcoll.Find(bson.M{"deviceid": d}).One(&dev)
	if err != nil {
		log.Println("Error while querying the Collection:\n", err)
		return nil, false
	}
	return &dev, true
}

func GetDevices(deviceList []string) Devices {

	var dList Devices
	for _, devID := range deviceList {
		if d, found := GetDevice(devID); found {
			dList = append(dList, d)
		}
	}
	return dList
}

func ToggleDeviceStatus(d string) (*Device, bool) {

	mcoll := GetCollection(DeviceCollName)

	if dev, found := GetDevice(d); found {

		if dev.DeviceStatus == "Active" {
			dev.DeviceStatus = "Inactive"
		} else {
			dev.DeviceStatus = "Active"
		}
		err := mcoll.Update(bson.M{"deviceid": d}, dev)
		if err != nil {
			log.Println("Error while updating the Document:\n", err)
			return nil, false
		}
		return dev, true
	} else {
		return nil, false
	}

}
