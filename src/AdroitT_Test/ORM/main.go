package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type UserData struct {
	allusers    string
	newusers    string
	deleteusers string
	updateusers string
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", greet)
	r.HandleFunc("/allusers", AllUsers)
	r.HandleFunc("/newusers", NewUsers)
	r.HandleFunc("/deleteusers", DeleteUsers)
	r.HandleFunc("/updateusers", UpdateUsers)
	http.Handle("/", r)
}
