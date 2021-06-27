package controller

import (
	"net/http"
)

func Files(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
	
}
