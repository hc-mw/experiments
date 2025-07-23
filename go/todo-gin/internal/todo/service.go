package todo

type TodoService interface {
	Create(userID uint, title string) error
	GetAllByUserID(userID uint) ([]Todo, error)
}

type todoService struct {
	repo TodoRepo
}

func NewTodoService(repo TodoRepo) TodoService {
	return &todoService{repo}
}

func (s *todoService) Create(userID uint, title string) error {
	return s.repo.Create(&Todo{UserID: userID, Title: title})
}

func (s *todoService) GetAllByUserID(userID uint) ([]Todo, error) {
	return s.repo.GetAllByUserID(userID)
}
