package http

import (
	"github.com/gin-gonic/gin"

	"github.com/VicAlarDev/kvault-back/internal/adapter/request"
	"github.com/VicAlarDev/kvault-back/internal/core/port"
)

type AuthHandler struct {
	svc port.AuthService
}

func NewAuthHandler(svc port.AuthService) *AuthHandler {
	return &AuthHandler{
		svc: svc,
	}
}

// Login godoc
//
//	@Summary		Login and get an access token
//	@Description	Logs in a registered user and returns an access token if the credentials are valid.
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			request	body		request.LoginRequest	true	"Login request body"
//	@Success		200		{object}	authResponse			"Succesfully logged in"
//	@Failure		400		{object}	errorResponse			"Validation error"
//	@Failure		401		{object}	errorResponse			"Unauthorized error"
//	@Failure		500		{object}	errorResponse			"Internal server error"
//	@Router			/users/login [post]
func (ah *AuthHandler) Login(ctx *gin.Context) {
	var req request.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		validationError(ctx, err)
		return
	}

	token, err := ah.svc.Login(ctx, req.Email, req.Password)
	if err != nil {
		handleError(ctx, err)
		return
	}

	rsp := newAuthResponse(token)

	handleSuccess(ctx, rsp)
}
