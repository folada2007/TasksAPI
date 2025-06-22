package dto

import "LongTaskAPI/internal/domain/entity"

type AllTasks struct {
	Tasks []entity.Task `json:"tasks"`
}
