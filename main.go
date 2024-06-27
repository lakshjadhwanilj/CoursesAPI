package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
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
	// return course.CourseId == "" && course.CourseName == ""
	return course.CourseName == ""
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
	fmt.Println("Get all courses")
	// Explicitly setting the Content-Type to inform the client that JSON data is being returned.
	w.Header().Set("Content-Type", "application/json")
	// Encode the courses variable to JSON and write it to the response
	json.NewEncoder(w).Encode(courses)
}

// Get One Course By ID
func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	// Grab id from request
	params := mux.Vars(r)

	// Looping through courses to find the matching id and returning the response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course found with given id.")
	return
}

// Create One Course
func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one Course")
	w.Header().Set("Content-Type", "application/json")

	// If body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	// If body has an empty JSON
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.isEmpty() {
		json.NewEncoder(w).Encode("No data in the JSON")
		return
	}

	// Generating a unique id & converting to string
	// Append course into courses
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

// Update One Course
func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one Course")
	w.Header().Set("Content-Type", "application/json")

	// Grab id from request
	params := mux.Vars(r)

	// Loop through courses and find matching id
	for index, course := range courses {
		if course.CourseId == params["id"] {
			// Removing the existing value
			courses = append(courses[:index], courses[index+1:]...)

			// Add the new value with id
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
		}
	}

	// If id not found
	json.NewEncoder(w).Encode("Course id not found.")
}
