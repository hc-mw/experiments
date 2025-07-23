package user

import (
	"errors"

	"github.com/hardikchoksi151/todo-gin/pkg/hash"
)

type UserService interface {
	Register(username, password string) error
	Authenticate(username, password string) (*User, error)
}

type userService struct {
	repo UserRepo
}

func NewUserService(userRepo UserRepo) UserService {
	return &userService{userRepo}
}

func (s *userService) Register(username, password string) error {

	hashedPassword, err := hash.Generate(password)
	if err != nil {
		return err
	}

	return s.repo.Create(&User{
		Username: username,
		Password: hashedPassword,
	})
}

func (s *userService) Authenticate(username, password string) (*User, error) {
	user, err := s.repo.GetUserByName(username)
	if err != nil {
		return nil, err
	}

	if !hash.Compare(user.Password, password) {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}
