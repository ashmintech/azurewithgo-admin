package controller

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/services/preview/streamanalytics/mgmt/2020-03-01-preview/streamanalytics"
	"github.com/Azure/go-autorest/autorest/to"
	//	"github.com/Azure/azure-sdk-for-go/services/streamanalytics/mgmt/2016-03-01/streamanalytics"
)

func Files(w http.ResponseWriter, r *http.Request) {

	// This is a test bed for stream analytics.... check if they have the functionallity

	// 1. See if you can start the job..

	//	if err := tpl.ExecuteTemplate(w, "devices.gohtml", data.GetAllDevices()); err != nil {
	//		log.Fatalln("Not able to call the template", err)
	//	}

	//sjob := streamanalytics.newstreamingjobsclient(subscriptionID)

	sjobclient := streamanalytics.NewStreamingJobsClient(subscriptionID)

	sjobclient.Authorizer = authorize

	//	sjobclient.Start()

	//	sjob := &streamanalytics.stre

	//	fmt.Println(sjoblist)
	//	sjobclient.Get()
	for sjoblist, err := sjobclient.List(context.Background(), ""); sjoblist.NotDone(); err = sjoblist.Next() {
		if err != nil {
			log.Fatal(err)

		}

		fmt.Println("In here")
		for _, job := range sjoblist.Values() {
			state := to.String(job.StreamingJobProperties.JobState)
			fmt.Println("StreamJob Name: ", to.String(job.Name))
			fmt.Println("StreamJob Name: ", state)

			if state == "Stoppessd" {
				//start the job
				param := &streamanalytics.StartStreamingJobParameters{
					OutputStartMode: streamanalytics.LastOutputEventTime,
				}
				s, err := sjobclient.Start(context.Background(), "rg-goweb", "gostream", param)
				if err != nil {
					log.Println("Not able to start the stream job")
				}

				log.Println(s.Status())
				//s.Result()

			}
			//	resList, _ := resClient.ListByResourceGroup(context.Background(), to.String(resGrp.Name), "", "", nil)
			//	for _, res := range resList.Values() {
			//		fmt.Println("\t- Resource Name: ", to.String(res.Name), " | Resource Type: ", to.String(res.Type))
			//	}

		}

	}

	//	sjoblist.Next()

}
