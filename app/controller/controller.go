package controller

import (
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/recipes", getRecipes)
	router.GET("/recipes/:id", getRecipe)
	router.POST("/recipes", createRecipe)
	router.PATCH("/recipes/:id", updateRecipe)
	router.DELETE("/recipes/:id", deleteRecipe)
	return router
}
