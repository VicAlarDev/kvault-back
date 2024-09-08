package http

import (
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/VicAlarDev/kvault-back/internal/core/domain"
	"github.com/VicAlarDev/kvault-back/internal/core/port"
)

const (
	// authorizationHeaderKey is the key for authorization header in the request
	authorizationHeaderKey = "authorization"
	// authorizationType is the accepted authorization type
	authorizationType = "bearer"
	// authorizationPayloadKey is the key for authorization payload in the context
	authorizationPayloadKey = "authorization_payload"
)

// authMiddleware is a middleware to check if the user is authenticated
func authMiddleware(token port.TokenService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)

		isEmpty := len(authorizationHeader) == 0
		if isEmpty {
			err := domain.ErrEmptyAuthorizationHeader
			handleAbort(ctx, err)
			return
		}

		fields := strings.Fields(authorizationHeader)
		isValid := len(fields) == 2
		if !isValid {
			err := domain.ErrInvalidAuthorizationHeader
			handleAbort(ctx, err)
			return
		}

		currentAuthorizationType := strings.ToLower(fields[0])
		if currentAuthorizationType != authorizationType {
			err := domain.ErrInvalidAuthorizationType
			handleAbort(ctx, err)
			return
		}

		accessToken := fields[1]
		payload, err := token.Validate(accessToken)
		if err != nil {
			handleAbort(ctx, err)
			return
		}

		ctx.Set(authorizationPayloadKey, payload)
		ctx.Next()
	}
}
