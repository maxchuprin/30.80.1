package storage

import "30.80.1/pkg/storage/model"

type Interface interface {
	Tasks() ([]model.Task, error)
	NewTask(model.Task) (int, error)
	TasksByAuthor(int) ([]model.Task, error)
	TaskById(int) (model.Task, error)
	UpdateTask(int, model.Task) error
	DeleteTask(int) error
}
