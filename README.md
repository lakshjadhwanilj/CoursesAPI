# Courses API

This project is a simple REST API built with Go, using the Gorilla Mux router for handling HTTP requests. It provides basic CRUD operations for managing courses.

## Features

- List all courses
- Get a single course by ID
- Create a new course
- Update an existing course
- Delete a course

## Installing Courses API

To install Courses API, follow these steps:

1. Clone the repository to your local machine
2. Navigate to the cloned directory
3. Install the dependencies by running:

```sh
go mod tidy
```

## Endpoints
- `GET /` - Welcome message
- `GET /courses` - Returns a list of all courses
- `GET /course/{id}` - Returns a single course by ID
- `POST /course` - Creates a new course
- `PUT /course/{id}` - Updates an existing course by ID
- `DELETE /course/{id}` - Deletes a course by ID

