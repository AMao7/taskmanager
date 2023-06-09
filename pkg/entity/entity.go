package entity

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	ID        uint      `gorm:"primary_key" validate:"required,gt=0"`
	Title     string    `gorm:"type:varchar(100);not null" validate:"required,min=3,max=100"`
	Content   string    `gorm:"type:varchar(1000);not null" validate:"required,min=10,max=100"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (t *Task) UpdateFrom(other Task) {
	// Update fields from 'other' to 't'
	t.Title = other.Title
	t.Content = other.Content
	// continue for all fields that can be updated
}

// TaskStore is an interface with five methods, each corresponding to a different operation you might want to perform on tasks:

type TaskStore interface {
	GetAll() ([]Task, error)
	GetByID(id uint) (*Task, error)
	Create(task *Task) error
	Update(task *Task) error
	Delete(id uint) error
}

type GormTaskStore struct {
	db *gorm.DB
}

func NewGormTaskStore(x *gorm.DB) GormTaskStore {
	return GormTaskStore{db: x}
}

func (x GormTaskStore) Create(task *Task) error {
	return x.db.Create(task).Error
}

func (store GormTaskStore) GetAll() ([]Task, error) {
	var tasks []Task
	err := store.db.Find(&tasks).Error
	return tasks, err
}

func (store GormTaskStore) GetByID(id uint) (*Task, error) {
	var task Task
	err := store.db.First(&task, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &task, nil
}

func (store GormTaskStore) Update(task *Task) error {
	return store.db.Save(task).Error
}

func (store GormTaskStore) Delete(id uint) error {
	return store.db.Delete(&Task{}, id).Error
}
