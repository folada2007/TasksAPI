package entity

import (
	"time"
)

type Task struct {
	ID        int64
	Status    string
	CreatedAt time.Time
	Title     string
	EndAt     time.Time
}
