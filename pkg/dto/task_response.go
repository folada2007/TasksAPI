package dto

import "time"

type TaskResponse struct {
	Id        int64     `json:"id"`
	Status    string    `json:"status"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	Duration  float64   `json:"duration"`
}
