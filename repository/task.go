package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type TaskRepository interface {
	Store(task *model.Task) error
	Update(id int, task *model.Task) error
	Delete(id int) error
	GetByID(id int) (*model.Task, error)
	GetList() ([]model.Task, error)
	GetTaskCategory(id int) ([]model.TaskCategory, error)
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepo(db *gorm.DB) *taskRepository {
	return &taskRepository{db}
}

func (t *taskRepository) Store(task *model.Task) error {
	err := t.db.Create(&task).Error
	if err != nil {
		return err
	}

	return nil
}

func (t *taskRepository) Update(id int, task *model.Task) error {
	err := t.db.Model(&model.Task{}).Where("id = ?", id).Updates(map[string]interface{}{
		"title":       task.Title,
		"deadline":    task.Deadline,
		"priority":    task.Priority,
		"status":      task.Status,
		"category_id": task.CategoryID,
		"user_id":     task.UserID,
	}).Error

	return err
}

func (t *taskRepository) Delete(id int) error {
	err := t.db.Where("id = ?", id).Delete(&model.Task{}).Error

	return err
}

func (t *taskRepository) GetByID(id int) (*model.Task, error) {
	var task model.Task
	err := t.db.First(&task, id).Error
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (t *taskRepository) GetList() ([]model.Task, error) {
	tasks := []model.Task{}
	err := t.db.Model(&model.Task{}).Scan(&tasks).Error

	return tasks, err
}

func (t *taskRepository) GetTaskCategory(id int) ([]model.TaskCategory, error) {
	taskCategory := []model.TaskCategory{}
	err := t.db.Table("tasks").Where("tasks.id = ?", id).Select("tasks.id AS id, tasks.title AS title, categories.name AS category").Joins("LEFT JOIN categories ON categories.id = tasks.category_id").Scan(&taskCategory).Error

	return taskCategory, err
}
