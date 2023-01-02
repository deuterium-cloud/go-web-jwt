package main

import (
	"github.com/deuterium-cloud/go-web-jwt/api"
	"github.com/deuterium-cloud/go-web-jwt/initialize"
	"github.com/gin-gonic/gin"
)

func init() {
	initialize.LoadEnvVariables()
}

func main() {
	var router = gin.Default()
	api.RouteAtoms(router)
	api.RouteUsers(router)
	router.Run()
}
