package controller

import (
	"context"
	"log"
	"net/http"
	"path"

	"github.com/Azure/azure-sdk-for-go/services/preview/streamanalytics/mgmt/2020-03-01-preview/streamanalytics"
	"github.com/Azure/go-autorest/autorest/to"
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

func findDeviceJobs(devID string) *data.DeviceJobs {
	// This function will return the state of 3 jobs...for that device.
	return data.GetDeviceJobs(devID)

}

func DeviceDetails(w http.ResponseWriter, r *http.Request) {

	devID := path.Base(r.URL.Path)
	var d *data.Device
	var c *data.Customer
	var j *data.DeviceJobs

	d, found := existDevice(devID)

	if !found {
		http.Redirect(w, r, "/admin/devices", http.StatusSeeOther)
		return
	}

	c, _ = findCustomer4Device(d.DeviceID)
	j = findDeviceJobs(d.DeviceID)

	type sendData struct {
		Dev    *data.Device
		Cust   *data.Customer
		DevJob *data.DeviceJobs
	}

	if err := tpl.ExecuteTemplate(w, "devicedetails.gohtml", sendData{d, c, j}); err != nil {
		log.Fatalln("Not able to call the template", err)
	}
}

const (
	RegGrp       = "rg-goweb"
	GoMonitorJob = "gostream"
)

func DeviceJobs(w http.ResponseWriter, r *http.Request) {
	//fmt.Sprintln(w, "In Device Job")

	// Find the current status of 3 streaming jobs and show the status here

	type jobData struct {

		JobName   string
		JobDisplayName string
		JobStatus string
	}

	var jobs []jobData

	// Right now, its just 1 job
	// resource Group: rg-goweb
	// jobname: gostream
	// Monitor Job

	sjobclient := streamanalytics.NewStreamingJobsClient(subscriptionID)
	sjobclient.Authorizer = authorize

	job, err := sjobclient.Get(context.Background(), RegGrp, GoMonitorJob, "")
	if err != nil {
		log.Fatal("Error:", err)
	}

	defer job.Body.Close()

	state := to.String(job.StreamingJobProperties.JobState)
	log.Println("StreamJob Name: ", to.String(job.Name))
	log.Println("StreamJob Name: ", state)

	monitorjob := jobData{
		JobName:   to.String(job.Name),
		JobDisplayName: "Monitoring Iot Data",
		JobStatus: state,
	}

	jobs = append(jobs, monitorjob)

	if err := tpl.ExecuteTemplate(w, "devicejobs.gohtml", jobs); err != nil {
		log.Fatalln("Not able to call the template", err)
	}

}
