package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Models for Database
// Course Model
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

// Author Model
type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

// Fake Database
var courses []Course

// Middleware / Helper Function
func (course *Course) isEmpty() bool {
	return course.CourseId == "" && course.CourseName == ""
}

func main() {

}

// Controllers
// The Serve Home Route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to APIs in GoLang</h1>"))
}

// Get All Courses
func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all Courses")
	// Explicitly setting the Content-Type to inform the client that JSON data is being returned.
	w.Header().Set("Content-Type", "application/json")
	// Encode the courses variable to JSON and write it to the response
	json.NewEncoder(w).Encode(courses)
}
