package todo

import "gorm.io/gorm"

type TodoRepo interface {
	Create(todo *Todo) error
	GetAllByUserID(userId uint) ([]Todo, error)
}

type todoStore struct {
	db *gorm.DB
}

func NewTodoRepo(db *gorm.DB) TodoRepo {
	return &todoStore{db}
}

func (s *todoStore) Create(todo *Todo) error {
	return s.db.Create(todo).Error
}

func (s *todoStore) GetAllByUserID(userId uint) ([]Todo, error) {
	var todos []Todo
	err := s.db.Where("user_id = ?", userId).Find(&todos).Error
	return todos, err
}
