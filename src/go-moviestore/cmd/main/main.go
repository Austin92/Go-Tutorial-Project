package main

import (
	"log"
	"net/http"

	"github.com/Austin92/go-moviestore/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

/* creating the router */
func main() {
	r := mux.NewRouter() //router//
	routes.RegisterMovieStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
