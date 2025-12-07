package todo

import "time"

type CreateInput struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	DueDate     *time.Time `json:"due_date"`
}

type UpdateInput struct {
	Title       *string    `json:"title"`
	Description *string    `json:"description"`
	Status      *Status    `json:"status"`
	DueDate     *time.Time `json:"due_date"`
}
