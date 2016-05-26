package links

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type LinkService interface {
	Get(string) (*Link, error)
	Update(slug string, url string) (*Link, error)
	Delete(string) error
	List() ([]*Link, error)
}

type errResponse struct {
	Message string `json:"message"`
}

func BindList(svc LinkService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var response interface{}
		response, err := svc.List()
		fmt.Println(response, err)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response = errResponse{err.Error()}
		}
		w.Header().Set("Content-Type", "application/javascript")
		json.NewEncoder(w).Encode(&response)
	}
}

func BindRoutes(svc LinkService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		slug := vars["slug"]
		r.ParseForm()
		url := r.PostFormValue("url")
		var response interface{}
		var err error
		switch r.Method {
		case "GET":
			link, err := svc.Get(slug)
			if err == nil {
				http.Redirect(w, r, link.ToRedirect(), http.StatusSeeOther)
				return
			}
		case "PUT", "POST":
			response, err = svc.Update(slug, url)
		case "DELETE":
			err = svc.Delete(slug)
		}
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response = errResponse{err.Error()}
		}
		w.Header().Set("Content-Type", "application/javascript")
		json.NewEncoder(w).Encode(&response)
	}
}
