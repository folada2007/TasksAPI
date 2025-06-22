package factory

import (
	"LongTaskAPI/internal/domain/entity"
	"LongTaskAPI/pkg/dto"
)

func ToTask(dto dto.TaskRequestDto) entity.Task {
	return entity.Task{
		Title: dto.Title,
	}
}
