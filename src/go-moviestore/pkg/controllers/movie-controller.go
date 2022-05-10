package controllers

import (
	"GO-MOVIESTORE/pkg/models"
	"GO-MOVIESTORE/pkg/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Austin92/go-moviestore/pkg/models"
	"github.com/Austin92/go-moviestore/pkg/utils"
	"github.com/gorilla/mux"
)

var NewMovies models.Movie

func GetMovie(w http.ResponseWriter, r *http.Request){
	NewMovies := models.GetAllMovies()
	res, _ := json.Marshal(NewMovies)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetMovieById(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	movieId := vars["movieId"]
	ID, err := strconv.ParseInt(movieId,0,0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	movieDetails, _:= models.GetMovieById(ID)
	res, _ := json.Marshal(movieDetails)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateMovie(w http.ResponseWriter, r *http.Request){
	CreateMovie := &models.Movie{}
	utils.ParseBody(r, CreateMovie)
	b := CreateMovie.CreateMovie()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteMovie(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	movieId := vars["movieId"]
	ID, err := strconv.ParseInt(movieId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	movie := models.DeleteMovie(ID)
	res, _ := json.Marshal(movie)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func UpdateMovie(w http.ResponseWriter, r *http.Request){
	var updateMovie := &models.Movie{}
	utils.Parse(r, updateMovie)
	vars := mux.Vars(r)
	movieId : vars["movieId"]
	ID, err := strconv.ParseInt(movieId, 0, 0)
	if err != nil  {
		fmt.Println("error while parsing")
	}
	movieDetails, db:=models.GetMovieById(ID)
	if updateMovie.Id != " "{
		movieDetails.Id = updateMovie.Id
	}
	if updateMovie.Isbn != " "{
		movieDetails.Isbn = updateMovie.Isbn
	}
	if updateMovie.Title != " "{
		movieDetails.Title = updateMovie.Title
	}
	db.Save(&movieDetails)
	res, _ := json.Marshal(movieDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}