package repository

import "LongTaskAPI/internal/domain/entity"

type TaskRepository interface {
	GetByID(id int64) (entity.Task, error)
	Create(task entity.Task) error
	GetAll() ([]entity.Task, error)
	Update(task entity.Task) error
	Delete(id int64) error
}
