package api

import (
	"net/http"
	"os"
	"time"

	"github.com/deuterium-cloud/go-web-jwt/middleware"
	"github.com/deuterium-cloud/go-web-jwt/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

func RouteUsers(router *gin.Engine) {
	router.GET("/users", middleware.RequireAuth, middleware.RequireAdmin, GetUsers)
	router.POST("/users/signup", Signup)
	router.POST("/users/login", Login)
}

func GetUsers(context *gin.Context) {
	context.JSON(http.StatusOK, models.Users)
}

func Signup(context *gin.Context) {

	// Extract New User from request
	var body models.UserRequest
	if err := context.Bind(&body); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Cannot bind body"})
		return
	}

	// Check if New User already exists
	userFromDB := findUserByUsername(body.Username)
	if userFromDB != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"})
		return
	}

	// Create new User
	hash, _ := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	newUser := models.User{
		Username:     body.Username,
		HashPassword: string(hash),
		Roles:        []string{"USER"},
	}

	models.Users = append(models.Users, newUser)

	// Send Response
	context.JSON(http.StatusCreated, gin.H{"message": "User successfully created"})
}

func Login(context *gin.Context) {

	// Extract Logindata from request
	var body models.UserRequest
	if err := context.Bind(&body); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Cannot bind body"})
		return
	}

	// Check if User exists
	userFromDB := findUserByUsername(body.Username)
	if userFromDB == nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Bad Username or Password"})
		return
	}

	// Check Password
	var err error = bcrypt.CompareHashAndPassword([]byte(userFromDB.HashPassword), []byte(body.Password))

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Bad Username or Password"})
		return
	}

	// Generate JWT

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":      userFromDB.Username,
		"roles":    userFromDB.Roles,
		"exp":      time.Now().Add(time.Hour).Unix(),
		"token_id": uuid.NewV4().String(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		context.JSON(500, gin.H{"error": "Error creating JWT token"})
		return
	}

	// Send Response

	context.SetSameSite(http.SameSiteLaxMode)
	context.SetCookie("Authorization", tokenString, 3600, "", "", false, true)
	context.Header("X-TOKEN-JWT", tokenString)
	context.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func findUserByUsername(username string) *models.User {
	for index, user := range models.Users {
		if user.Username == username {
			return &models.Users[index]
		}
	}
	return nil
}
