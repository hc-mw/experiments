package user

import "gorm.io/gorm"

type UserRepo interface {
	Create(user *User) error
	GetUserByName(username string) (*User, error)
}

type store struct {
	db *gorm.DB
}

func NewUserStore(db *gorm.DB) UserRepo {
	return &store{db}
}

func (s *store) Create(user *User) error {
	return s.db.Create(user).Error
}

func (s *store) GetUserByName(username string) (*User, error) {
	var user User
	err := s.db.Where("username = ?", username).First(&user).Error
	return &user, err
}
