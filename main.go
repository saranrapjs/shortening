package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/saranrapjs/shortening/pkg/links"
)

func main() {
	r := mux.NewRouter()

	// root handler, for the index
	fs := http.FileServer(http.Dir("dist/"))
	r.PathPrefix("/manage").Handler(http.StripPrefix("/manage", fs))

	// API + 301 redirector
	svc := links.NewLinkService()
	r.HandleFunc("/{slug}", links.BindRoutes(svc))

	// individual edit pages, for links
	r.Handle("/{slug}/edit", DynamicFileServer(fs))

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/manage", http.StatusMovedPermanently)
	})

	http.Handle("/", r)
	srv := &http.Server{
		ReadTimeout: 30 * time.Second,
		Addr:        ":8080",
	}
	log.Fatal(srv.ListenAndServe())
}

func DynamicFileServer(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = ""
		h.ServeHTTP(w, r)
	})
}
