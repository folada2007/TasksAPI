package in_memory

import (
	"LongTaskAPI/internal/domain/entity"
	"LongTaskAPI/internal/utils"
	"sync"
)

type InMemoryTaskRepo struct {
	tasks map[int64]entity.Task
	mu    sync.RWMutex
}

func NewInMemoryTaskRepo() *InMemoryTaskRepo {
	return &InMemoryTaskRepo{
		tasks: make(map[int64]entity.Task),
	}
}

func (t *InMemoryTaskRepo) GetByID(id int64) (entity.Task, error) {
	t.mu.RLock()
	defer t.mu.RUnlock()

	task, ok := t.tasks[id]
	if !ok {
		return entity.Task{}, utils.ErrorNotFound
	}
	return task, nil
}

func (t *InMemoryTaskRepo) Create(task entity.Task) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.tasks[task.ID] = task
	return nil
}

func (t *InMemoryTaskRepo) GetAll() ([]entity.Task, error) {
	t.mu.RLock()
	defer t.mu.RUnlock()

	var AllTasks []entity.Task

	for _, task := range t.tasks {
		AllTasks = append(AllTasks, task)
	}
	return AllTasks, nil
}

func (t *InMemoryTaskRepo) Update(task entity.Task) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.tasks[task.ID] = task
	return nil
}

func (t *InMemoryTaskRepo) Delete(id int64) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	delete(t.tasks, id)
	return nil
}
