package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hardikchoksi151/todo-gin/config"
	"github.com/hardikchoksi151/todo-gin/pkg/jwt"
)

type Handler struct {
	userService UserService
	config      config.Config
}

func NewHandler(userService UserService, cnf config.Config) *Handler {
	return &Handler{userService, cnf}
}

// Register handles user registration.
func (h *Handler) Register(c *gin.Context) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.userService.Register(input.Username, input.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "registration successful"})
}

// Login handles user authentication and token generation.
func (h *Handler) Login(c *gin.Context) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := h.userService.Authenticate(input.Username, input.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}
	token, err := jwt.GenerateToken(int(user.ID), h.config.JWTSecret)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "could not generate token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}
