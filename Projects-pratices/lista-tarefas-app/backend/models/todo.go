package models

import (
    "time"
)

type Todo struct {
    ID        int       `json:"id"`
    Title     string    `json:"title"`
    Completed bool      `json:"completed"`
    CreatedAt time.Time `json:"createdAt"`
}

type TodoRequest struct {
    Title string `json:"title"`
}