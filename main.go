package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"

	cont "github.com/ashmintech/azurewithgo-admin/controller"
)

const hostAddress = ":9000"

func main() {

	r := mux.NewRouter()

	admin := r.Methods(http.MethodGet).Subrouter()
	admin.HandleFunc("/admin", cont.AdminPortal)
	admin.HandleFunc("/admin/devicejobs", cont.DeviceJobs)
	admin.HandleFunc("/admin/devices", cont.Devices)
	admin.HandleFunc("/admin/devices/{id}", cont.DeviceDetails)
	admin.HandleFunc("/admin/customers", cont.Customers)
	admin.HandleFunc("/admin/customers/{id}", cont.CustomerDetails)
	admin.HandleFunc("/admin/files", cont.Files)
	admin.HandleFunc("/admin/devicestatus/{id}", cont.DeviceToggleStatus)
	admin.Use(cont.AdminMiddleware)

	// Static Files
	r.PathPrefix("/admin").Handler(http.FileServer(http.Dir(".")))

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/admin", http.StatusMovedPermanently)
	})

	srv := &http.Server{
		Handler:      r,
		Addr:         hostAddress,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		log.Println(("Starting Admin Portal on Port 9000"))

		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	srv.Shutdown(ctx)
	log.Println("Server shutting down")
	os.Exit(0)
}
