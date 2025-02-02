package middleware

import (
	"net/http"
	"strings"
	"web-article/model"
	"web-article/utils/service"

	"github.com/gin-gonic/gin"
)

type authMiddleware struct {
	jwtService service.JWTService
}

type authHeader struct {
	AuthorizationHeader string `header:"Authorization" binding:"required"`
}

type AuthMiddleware interface {
	RequireToken(roles ...string) gin.HandlerFunc
}

func (a *authMiddleware) RequireToken(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var aH authHeader

		err := c.ShouldBindHeader(&aH)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Header Unautorized"})
			return
		}

		token := strings.Replace(aH.AuthorizationHeader, "Bearer ", "", 1)

		tokenClaim, err := a.jwtService.VerifyToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Token Unautorized"})
			return
		}

		c.Set("user", model.UserLogin{ID: tokenClaim.UserId, Role: tokenClaim.Role})

		validRole := false
		for _, role := range roles {
			if role == tokenClaim.Role {
				validRole = true
				break
			}
		}
		if !validRole {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Forbiden Resourse"})
			return
		}

		c.Next()
	}
}

func NewAuthMiddleware(jwtService service.JWTService) AuthMiddleware {
	return &authMiddleware{jwtService: jwtService}
}
