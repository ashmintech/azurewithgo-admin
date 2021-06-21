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

type DeviceAnomaly struct {
	DeviceID    string  `json:"deviceid"`
	TimeStamp   string  `json:"timestamp"`
	FridgeTemp  float64 `json:"fridgetemp"`
	CountFridge int     `json:"countfridgetemp"`
}

type sendDeviceAnomaly struct {
	DeviceID        string
	DateTime        string
	FridgeTemp      string
	CountFridgeTemp string
}

var anomalyData = []DeviceAnomaly{}

func GetAnomaly(devID string) []sendDeviceAnomaly {

	sendData := []sendDeviceAnomaly{}

	for i := len(anomalyData) - 1; i >= 0; i-- {
		b := anomalyData[i]
		if b.DeviceID == devID {
			t, _ := time.Parse(time.RFC3339, b.TimeStamp)

			a := sendDeviceAnomaly{
				b.DeviceID,
				t.Format(time.RFC822),
				fmt.Sprintf("%.2f", b.FridgeTemp),
				fmt.Sprintf("%d", b.CountFridge),
			}
			sendData = append(sendData, a)
		}
	}
	return sendData
}

func ShowAnomaly() []sendDeviceAnomaly {

	sendData := []sendDeviceAnomaly{}

	for i := len(anomalyData) - 1; i >= 0; i-- {
		b := anomalyData[i]
		t, _ := time.Parse(time.RFC3339, b.TimeStamp)

		a := sendDeviceAnomaly{
			b.DeviceID,
			t.Format(time.RFC822),
			fmt.Sprintf("%.2f", b.FridgeTemp),
			fmt.Sprintf("%d", b.CountFridge),
		}
		sendData = append(sendData, a)

	}
	return sendData
}

const (
	EventHubAnomalyEndPoint = "Endpoint=sb://goeventhubns.servicebus.windows.net/;SharedAccessKeyName=goanomalyjob_eventhuboutput_policy;SharedAccessKey=2ybXYHzX92XqmltxqaSbG0AD67Aa40k7aE5xRg1aMfg=;EntityPath=goanomaly"
)

func RunAnomalyListener() {

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	hub, err := eventhub.NewHubFromConnectionString(EventHubAnomalyEndPoint)
	if err != nil {
		log.Fatalln("Not able to create event hub from connection string: \n", err)
	}

	h, err := hub.GetRuntimeInformation(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	var aData = []DeviceAnomaly{}

	handler := func(c context.Context, event *eventhub.Event) error {
		
		err := json.Unmarshal([]byte(string(event.Data)), &aData)

		for _, a := range aData {
			anomalyData = append(anomalyData, a)
		}

		if err != nil {
			log.Fatalln("Error json:\n", err)
		}
		return nil
	}

	for _, partitionID := range h.PartitionIDs {

		listenerHandle, err := hub.Receive(ctx, partitionID, handler, eventhub.ReceiveFromTimestamp(time.Now().AddDate(0, 0, -7)))
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
		fmt.Println("There is error while closing the hub", err)
	}

}
