package links

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type LinkService interface {
	Get(string) (*Link, error)
	Update(slug string, link string) (*Link, error)
	Remove(string) error
}

type linkService struct {
}

func NewLinkService() *linkService {
	return &linkService{}
}

func (l linkService) Get(slug string) (*Link, error) {
	return &Link{"http://apple.com", slug}, nil
}

func (l linkService) Update(slug string, link string) (*Link, error) {
	return &Link{"http://apple.com", slug}, nil
}

func (l linkService) Remove(slug string) error {
	return nil
}

func BindRoutes(svc LinkService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		slug := vars["slug"]
		link, _ := vars["link"]
		var response interface{}
		var err error
		switch r.Method {
		case "GET":
			response, err = svc.Get(slug)
		case "PUT", "POST":
			response, err = svc.Update(slug, link)
		case "DELETE":
			err = svc.Remove(slug)
		}
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response = err
		}
		w.Header().Set("Content-Type", "application/javascript")
		json.NewEncoder(w).Encode(&response)
	}
}
