package main

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

func main() {
	
}