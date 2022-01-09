package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


// declaring the class variables
type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

// creating the array and doing the print out to the server.
type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Test Title", Desc: "Test Desc", Content: "Welcome to Adriot Go!!"},
	}

	fmt.Println("Endpoint Hit: ALl Articles Endpoint")
	json.NewEncoder(w).Encode(articles)
}

// homepage endpoint
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "homepage Endpoint hit")
}

func testPostArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "testPostArticle Endpoint hit")
}

func handlerRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", allArticles).Methods("GET")
	myRouter.HandleFunc("/articles", testPostArticles).Methods("Post")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	handlerRequests()
}
