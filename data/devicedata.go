package data

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	eventhub "github.com/Azure/azure-event-hubs-go/v3"
)

type DeviceData struct {
	DeviceID     string  `json:"deviceid"`
	TimeStamp    string  `json:"timestamp"`
	AvgFridge    float64 `json:"avgfridgetemp"`
	CountFridge  int     `json:"countfridgetemp"`
	AvgFreezer   float64 `json:"avgfreezertemp"`
	CountFreezer int     `json:"countfreezertemp"`
	
}

type sendDeviceData struct {
	DateTime   string
	AvgFridge  string
	AvgFreezer string
}

var data = []DeviceData{}

func init() {
	log.Println("In DeviceData init function")
	go runEventHubListener()
}

func GetDeviceData(devID string, day int) []sendDeviceData {
	log.Println("Get Device Data for:", devID)
	log.Println("Data for Day:", day)

	sendData := []sendDeviceData{}

	for _, b := range data {
		if b.DeviceID == devID {
			t, _ := time.Parse(time.RFC3339, b.TimeStamp)

			a := sendDeviceData{
				t.Format(time.RFC822),
				fmt.Sprintf("%.2f", b.AvgFridge),
				fmt.Sprintf("%.2f", b.AvgFreezer),
			}
			sendData = append(sendData, a)

		}
	}
	log.Println(sendData)

	//log.Println("Length of Data:", len(data))
	return sendData
}

func runEventHubListener() {
	log.Println("In Event Hub function")
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	str := "Endpoint=sb://goeventhubns.servicebus.windows.net/;SharedAccessKeyName=sendreceive;SharedAccessKey=rP4Lfj2o4hepWWqU9TYkixXines4PSa4emsgkYvfePg=;EntityPath=goeventhub"
	hub, err := eventhub.NewHubFromConnectionString(str)
	if err != nil {
		log.Fatalln("Not able to create event hub from connection string: \n", err)
	}

	h, err := hub.GetRuntimeInformation(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Starting Event Hub Handler")

	handler := func(c context.Context, event *eventhub.Event) error {

		//	fmt.Println((string(event.Data)))

		err := json.Unmarshal([]byte(string(event.Data)), &data)
		if err != nil {
			log.Fatalln("Error json:\n", err)
		} else {
			log.Println("Here:")
		}

		return nil
	}

	log.Println("Listening for Messages: ")
	for _, partitionID := range h.PartitionIDs {  
		listenerHandle, err := hub.Receive(ctx, partitionID, handler, eventhub.ReceiveWithLatestOffset())
		if err != nil {
			log.Fatalln("Error while creating a listener handler")
		}

		defer listenerHandle.Close(ctx)

	}

	// Wait for a signal to quit:
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, os.Kill)
	<-signalChan

	err = hub.Close(context.Background())
	if err != nil {
		fmt.Println("There is error wile closing the hub", err)
	}

}