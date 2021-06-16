package controller

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/Azure/go-autorest/autorest"
	"github.com/ashmintech/azurewithgo-admin/data"
)

var tpl *template.Template
var authorize autorest.Authorizer
var subscriptionID string
var err error
var cook *http.Cookie

func init() {

	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))

	cID := os.Getenv("AZURE_CLIENT_ID")
	cSecret := os.Getenv("AZURE_CLIENT_SECRET")
	tID := os.Getenv("AZURE_TENANT_ID")
	subscriptionID = os.Getenv("AZURE_SUBSCRIPTION_ID")

	// If any of the value is empty
	if cID == "" || cSecret == "" || tID == "" || subscriptionID == "" {
		log.Fatalln("Not able to set environmental variables")
	}

	go data.RunEventHubListener()

}
