package rest

import (
	"net/http"

	"justice/api"
	"justice/internal/model"

	"github.com/labstack/echo/v4"
)

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
