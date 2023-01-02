package api

import (
	"errors"
	"net/http"

	"github.com/deuterium-cloud/go-web-jwt/middleware"
	"github.com/deuterium-cloud/go-web-jwt/models"
	"github.com/gin-gonic/gin"
)

func RouteAtoms(router *gin.Engine) {
	router.GET("/atoms", GetAtoms)
	router.GET("/atoms/:id", GetAtom)
	router.PUT("/atoms/:id", middleware.RequireAuth, UpdateAtom)
	router.POST("/atoms", middleware.RequireAuth, AddAtom)
}

func GetAtoms(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, models.Atoms)
}

func GetAtom(context *gin.Context) {

	id := context.Param("id")
	atom, err := getAtomById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
	} else {
		context.IndentedJSON(http.StatusOK, atom)
	}

}

func UpdateAtom(context *gin.Context) {

	var newAtom models.Atom
	if err := context.BindJSON(&newAtom); err != nil {
		return
	}

	id := context.Param("id")

	if id != newAtom.ID {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "IDs are not the same!"})
		return
	}

	atom, err := getAtomById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	atom.AtomNumber = newAtom.AtomNumber
	atom.Mass = newAtom.Mass
	atom.Name = newAtom.Name
	atom.Symbol = newAtom.Symbol

	context.IndentedJSON(http.StatusOK, atom)
}

func AddAtom(context *gin.Context) {
	var newAtom models.Atom
	if err := context.BindJSON(&newAtom); err != nil {
		return
	}

	models.Atoms = append(models.Atoms, newAtom)
	context.IndentedJSON(http.StatusCreated, newAtom)
}

func getAtomById(id string) (*models.Atom, error) {
	for index, atom := range models.Atoms {
		if atom.ID == id {
			return &models.Atoms[index], nil
		}
	}

	return nil, errors.New("Atom with id=" + id + " not found")
}
