package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/namsral/flag"
	"github.com/saranrapjs/shortening/pkg/links"
)

func main() {
	r := mux.NewRouter()

	// root handler, for the index
	fs := http.FileServer(http.Dir("dist/"))
	r.PathPrefix("/manage").Handler(http.StripPrefix("/manage", fs))

	var dbHost string
	flag.StringVar(&dbHost, "dbhost", "", "db hostname")

	flag.Parse()

	dsn := fmt.Sprintf("%s/links", dbHost)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err, dsn)
	}

	// API + 301 redirector
	svc := links.NewLinkService(db)
	r.HandleFunc("/{slug}", links.BindRoutes(svc))

	// individual edit pages, for links
	r.Handle("/{slug}/edit", DynamicFileServer(fs))

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/manage", http.StatusMovedPermanently)
	})

	http.Handle("/", handlers.LoggingHandler(os.Stdout, r))
	srv := &http.Server{
		ReadTimeout: 30 * time.Second,
		Addr:        ":8080",
	}
	log.Print("Listening at 8080...")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func DynamicFileServer(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = ""
		h.ServeHTTP(w, r)
	})
}
