package services

import (
	"LongTaskAPI/internal/domain/entity"
	"LongTaskAPI/internal/domain/repository"
	"LongTaskAPI/internal/utils"
	"time"
)

const (
	StatusCreated  string = "created"
	StatusRunning  string = "running"
	StatusContinue string = "stopped"
)

type TaskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) *TaskService {
	return &TaskService{
		repo: repo,
	}
}

func (t *TaskService) Create(task entity.Task) (entity.Task, error) {
	task = entity.Task{
		ID:        utils.IdGenerator(),
		CreatedAt: time.Now(),
		Status:    StatusCreated,
		Title:     task.Title,
	}

	err := t.repo.Create(task)
	if err != nil {
		return entity.Task{}, err
	}
	return task, nil
}

func (t *TaskService) GetById(id int64) (entity.Task, error) {
	return t.repo.GetByID(id)
}

func (t *TaskService) GetAll() ([]entity.Task, error) {
	return t.repo.GetAll()
}

func (t *TaskService) StartTask(task entity.Task) (entity.Task, error) {
	task, err := t.repo.GetByID(task.ID)
	if err != nil {
		return entity.Task{}, err
	}
	task.Status = StatusRunning

	err = t.repo.Update(task)
	if err != nil {
		return entity.Task{}, err
	}
	go func(task entity.Task) {

		time.Sleep(120 * time.Second)
		task.Status = StatusContinue
		task.EndAt = time.Now()
		err = t.repo.Update(task)
		if err != nil {
			return
		}
	}(task)
	return task, nil
}

func (t *TaskService) DeleteTask(id int64) error {
	_, err := t.repo.GetByID(id)
	if err != nil {
		return err
	}
	err = t.repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
