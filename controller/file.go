package controller

import (
	"log"
	"net/http"

	"github.com/ashmintech/azurewithgo-admin/data"
)

func Files(w http.ResponseWriter, r *http.Request) {

	if err := tpl.ExecuteTemplate(w, "files.gohtml", data.GetFiles()); err != nil {
		log.Fatalln("Not able to call the template", err)
	}

}
func AddFile(w http.ResponseWriter, r *http.Request) {

	if err := tpl.ExecuteTemplate(w, "addfile.gohtml", nil); err != nil {
		log.Fatalln("Not able to call the template", err)
	}

}

func UploadFile(w http.ResponseWriter, r *http.Request) {

	resp, ok := data.PutFiles(r)

	if ok {
		log.Println("Operation Completed:", resp)
	} else {
		log.Println("Issue with operation:")
	}

	http.Redirect(w, r, "/admin/files", http.StatusSeeOther)

}
