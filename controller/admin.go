package controller

import (
	"log"
	"net/http"

	"github.com/Azure/go-autorest/autorest/azure/auth"

	//	data "github.com/ashmintech/azurewithgo-admin/moule3/data"

	"github.com/pborman/uuid"
)

func AdminPortal(w http.ResponseWriter, r *http.Request) {

	if err := tpl.ExecuteTemplate(w, "admin.gohtml", nil); err != nil {
		log.Fatalln("Not able to call the template", err)
	}
}

func AdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Set cookie and session
		cook, err = r.Cookie("adminsession")
		if err == http.ErrNoCookie {
			if authorize, err = auth.NewAuthorizerFromEnvironment(); err != nil {
				log.Fatal("Error while Authorization using env variables\n", err)
			}

			if subscriptionID == "" {
				http.Error(w, "Azure Subscription details does not exist", http.StatusNotFound)
				return
			}

			cook = &http.Cookie{
				Name:   "adminsession",
				Value:  uuid.New(),
				MaxAge: 600,
			}
			http.SetCookie(w, cook)
		}

		next.ServeHTTP(w, r)
	})
}
