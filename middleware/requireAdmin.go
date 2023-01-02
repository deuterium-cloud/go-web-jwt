package middleware

import (
	"net/http"

	"github.com/deuterium-cloud/go-web-jwt/models"
	"github.com/gin-gonic/gin"
)

func RequireAdmin(context *gin.Context) {
	user, _ := context.Get("user")

	var userPointer *models.User = user.(*models.User)

	roles := (*userPointer).Roles

	if !contains(roles, "ADMIN") {
		context.AbortWithStatus(http.StatusForbidden)
	}

	context.Next()

}

func contains(roles []string, role string) bool {
	for _, v := range roles {
		if v == role {
			return true
		}
	}

	return false
}
