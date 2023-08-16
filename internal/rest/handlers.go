package rest

import (
	"fmt"
	"net/http"

	"justice/api"
	"justice/internal/model"

	"github.com/gorilla/mux"
	"github.com/labstack/echo/v4"
)

func Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError) // Set the status code to 200 OK
	w.Write([]byte("ER"))
}

func Readiness(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK) // Set the status code to 200 OK
	w.Write([]byte("OK"))
}

func Base(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>This is the homepage. Try /hello and /hello/Sammy\n</h1>")
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello from Docker!\n</h1>")
}

func Name(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["name"]

	fmt.Fprintf(w, "<h1>Hello, %s!\n</h1>", title)
}

// CreateTaskHandler handles the creation of a new task
func CreateTaskHandler(c echo.Context) error {
	var taskRequest api.TaskRequest
	if err := c.Bind(&taskRequest); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	// Convert the taskRequest to a model.Task
	newTask := model.Task{
		Title:       taskRequest.Title,
		Description: taskRequest.Description,
	}

	// Save the new task to the database using model package
	err := model.CreateTask(&newTask)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to create task")
	}

	// Convert the model.Task to a TaskResponse
	taskResponse := api.TaskResponse{
		ID:          newTask.ID,
		Title:       newTask.Title,
		Description: newTask.Description,
	}

	return c.JSON(http.StatusCreated, taskResponse)
}
