package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByEmail(email string) (model.User, error)
	CreateUser(user model.User) (model.User, error)
	GetUserTaskCategory() ([]model.UserTaskCategory, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) GetUserByEmail(email string) (model.User, error) {
	user := model.User{}
	r.db.Table("users").Where("email = ?", email).First(&user)

	return user, nil
}

func (r *userRepository) CreateUser(user model.User) (model.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) GetUserTaskCategory() ([]model.UserTaskCategory, error) {
	userTaskCategory := []model.UserTaskCategory{}
	err := r.db.Table("users").Select("users.id AS id, users.fullname AS fullname, users.email AS email, tasks.title AS task, tasks.deadline AS deadline, tasks.priority AS priority, tasks.status AS status, categories.name AS category").Joins("LEFT JOIN tasks ON tasks.user_id = users.id").Joins("LEFT JOIN categories ON categories.id = tasks.category_id").Scan(&userTaskCategory).Error

	return userTaskCategory, err
}
