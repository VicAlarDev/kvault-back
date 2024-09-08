package http

import (
	"github.com/VicAlarDev/kvault-back/internal/core/port"
	"github.com/gin-gonic/gin"
	"log"

	"github.com/VicAlarDev/kvault-back/internal/adapter/request"
	"github.com/VicAlarDev/kvault-back/internal/core/domain"
)

type UserHandler struct {
	svc port.UserService
}

func NewUserHandler(svc port.UserService) *UserHandler {
	return &UserHandler{
		svc: svc,
	}
}

// Register godoc
//
//	@Summary		Register a new user
//	@Description	create a new user account
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			RegisterRequest	body		request.RegisterRequest	true	"Register request"
//	@Success		200				{object}	userResponse			"User created"
//	@Failure		400				{object}	errorResponse			"Validation error"
//	@Failure		401				{object}	errorResponse			"Unauthorized error"
//	@Failure		404				{object}	errorResponse			"Data not found error"
//	@Failure		409				{object}	errorResponse			"Data conflict error"
//	@Failure		500				{object}	errorResponse			"Internal server error"
//	@Router			/users [post]
func (uh *UserHandler) Register(c *gin.Context) {
	var req request.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	user := domain.User{
		Name:     req.Name,
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	_, err := uh.svc.Register(c, &user)
	if err != nil {
		handleError(c, err)
		return
	}

	res := newUserResponse(&user)

	log.Printf("User data: %+v", user)

	handleSuccess(c, res)
}
