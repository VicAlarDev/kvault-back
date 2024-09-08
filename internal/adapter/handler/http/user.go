package http

import (
	"github.com/gin-gonic/gin"

	"github.com/VicAlarDev/kvault-back/internal/adapter/request"
	"github.com/VicAlarDev/kvault-back/internal/core/domain"
	"github.com/VicAlarDev/kvault-back/internal/core/service"
)

type UserHandler struct {
	svc service.UserService
}

func NewUserHandler(svc service.UserService) *UserHandler {
	return &UserHandler{
		svc: svc,
	}
}

type RegisterRequest struct {
	Name     string `json:"name" binding:"required" example:"John Doe"`
	Email    string `json:"email" binding:"required,email" example:"test@example.com"`
	Password string `json:"password" binding:"required,min=8" example:"12345678"`
}

// Register godoc
//
//	@Summary		Register a new user
//	@Description	create a new user account
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			RegisterRequest	body		RegisterRequest	true	"Register request"
//	@Success		200				{object}	userResponse	"User created"
//	@Failure		400				{object}	errorResponse	"Validation error"
//	@Failure		401				{object}	errorResponse	"Unauthorized error"
//	@Failure		404				{object}	errorResponse	"Data not found error"
//	@Failure		409				{object}	errorResponse	"Data conflict error"
//	@Failure		500				{object}	errorResponse	"Internal server error"
//	@Router			/users [post]
func (uh *UserHandler) Register(c *gin.Context) {
	var req request.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	user := &domain.User{
		Name:     req.Name,
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	user, err := uh.svc.Register(c.Request.Context(), user)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"data": user,
	})
}
