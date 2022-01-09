package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Model for course - file

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json: "price"`
	Author      *Author `json: "author`
}

type Author struct {
	Fullname string `json:"Fullname"`
	Website  string `json:"website"`
}

//fake DB
var course []Course

//middleware, helper - file
func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

func main() {

}

//Controllers - file

//serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-type", "applicatioan/json")
	json.NewEncoder(w).Encode(course)
}

func getOneCourse() (w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "applicatioan/json")

	//grab id from request
	params := mux.Vars(r)

	//loops through course, find matching id and return the response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
		}
	}
	json.NewEncoder(w).Encode("No course found ")
	return
}
