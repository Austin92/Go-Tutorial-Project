package main

import (
	"fmt"
	"net/http"
)


func AllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "User endpoint hit")
}

func NewUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "New User endpoint hit")
}

func DeleteUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete User endpoint hit")
}

func UpdateUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "update User endpoint hit")
}
