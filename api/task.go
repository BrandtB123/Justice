package api

// TaskRequest defines the structure for creating a new task
type TaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// TaskResponse defines the structure for the response when a task is created
type TaskResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
