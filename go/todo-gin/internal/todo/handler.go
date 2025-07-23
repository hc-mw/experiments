package todo

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	todoService TodoService
}

func NewHandler(ts TodoService) *Handler {
	return &Handler{ts}
}

func (h *Handler) CreateTodo(c *gin.Context) {
	var input struct {
		Title string `json:"title"`
	}
	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID := c.GetUint("userID")
	if err := h.todoService.Create(userID, input.Title); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "todo created"})
}

func (h *Handler) GetTodos(c *gin.Context) {
	userID := c.GetUint("userID")
	todos, err := h.todoService.GetAllByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todos)
}
