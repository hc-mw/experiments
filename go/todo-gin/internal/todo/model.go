package todo

import (
	"github.com/hardikchoksi151/todo-gin/internal/user"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title     string
	Completed bool
	UserID    uint
	User      user.User
}
