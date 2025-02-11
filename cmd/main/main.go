package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/trieungochai/go_bookstore-mgmt/pkg/routes"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterBookstoreRoutes(router)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe("3000", router))
}
