package controller

import (
	"net/http"
	"time"

	"github.com/kkato/recipe-api/app/model"

	"github.com/gin-gonic/gin"
)

var err error

func getRecipes(c *gin.Context) {
	recipes := model.GetRecipes()
	c.JSON(http.StatusOK, gin.H{"recipes": recipes})
}

func getRecipe(c *gin.Context) {
	recipes := model.GetRecipe(c.Param("id"))
	c.JSON(http.StatusOK, gin.H{"message": "Recipe details by id", "recipe": recipes})
}

func createRecipe(c *gin.Context) {
	var recipe model.Recipe
	err = c.ShouldBindJSON(&recipe)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "Recipe creation failed!", "required": "title, making_time, serves, ingredients, cost"})
		return
	}
	err = model.CreateRecipe(recipe)
	createdRecipe := model.GetRecipeByTitle(recipe.Title)
	c.JSON(http.StatusOK, gin.H{"message": "Recipe successfully created!", "recipe": createdRecipe})
}

func updateRecipe(c *gin.Context) {
	var recipe model.Recipe
	err = c.ShouldBindJSON(&recipe)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "Recipe update failed!", "required": "title, making_time, serves, ingredients, cost"})
		return
	}
	recipe.ID = c.Param("id")
	recipe.UpdatedAt = time.Now()
	model.UpdateRecipe(recipe)
	c.JSON(http.StatusOK, gin.H{"message": "Recipe successfully updated!", "recipe": recipe})
}

func deleteRecipe(c *gin.Context) {
	err = model.DeleteRecipe(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "Recipe not found!"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Recipe successfully deleted!"})
	}

}
