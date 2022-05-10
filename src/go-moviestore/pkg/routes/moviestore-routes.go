package routes

import (
	"github.com/Austin92/go-moviestore/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterMovieStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/movies", controllers.getMovies).Methods("GET")
	router.HandleFunc("/movies/{id}", controllers.getMovies).Methods("GET")
	router.HandleFunc("/movies/", controllers.createMovie).Methods("POST")
	router.HandleFunc("/movies/{id}", controllers.updateMovie).Methods("PUT")
	router.HandleFunc("/movies/{id}", controllers.deleteMovies).Methods("DELETE")
}
