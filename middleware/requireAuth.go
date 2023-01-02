package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/deuterium-cloud/go-web-jwt/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func RequireAuth(context *gin.Context) {

	// Get the header
	tokenString := context.GetHeader("X-TOKEN-JWT")
	// token, err := context.Cookie("Authorization")

	if tokenString == "" {
		context.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Check the expiration time
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			context.AbortWithStatus(http.StatusUnauthorized)
		}

		// Find the user with token sub
		user := findUserByUsername(claims["sub"].(string))
		if user == nil {
			context.AbortWithStatus(http.StatusUnauthorized)
		}

		// Attach to req
		context.Set("user", user)

		// Continue
		context.Next()

	} else {
		context.AbortWithStatus(http.StatusUnauthorized)
	}

}

func findUserByUsername(username string) *models.User {
	for index, user := range models.Users {
		if user.Username == username {
			return &models.Users[index]
		}
	}
	return nil
}
