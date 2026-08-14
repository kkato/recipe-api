package handler

import (
	"net/http"

	"github.com/kkato/recipe-api/internal/model"
	"github.com/kkato/recipe-api/internal/repository"

	"github.com/gin-gonic/gin"
)

type RecipeHandler interface {
	GetRecipes(c *gin.Context)
	GetRecipe(c *gin.Context)
	CreateRecipe(c *gin.Context)
	UpdateRecipe(c *gin.Context)
	DeleteRecipe(c *gin.Context)
}

type recipeHandler struct {
	repo repository.RecipeRepository
}

func NewRecipeHandler(repo repository.RecipeRepository) RecipeHandler {
	return &recipeHandler{repo: repo}
}

func (h *recipeHandler) GetRecipes(c *gin.Context) {
	recipes, err := h.repo.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "Failed to fetch recipes"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"recipes": recipes})
}

func (h *recipeHandler) GetRecipe(c *gin.Context) {
	recipe, err := h.repo.GetByID(c.Request.Context(), c.Param("id"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "Recipe not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Recipe details by id", "recipes": []model.Recipe{recipe}})
}

func (h *recipeHandler) CreateRecipe(c *gin.Context) {
	var recipe model.Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "Recipe creation failed!", "required": "title, making_time, serves, ingredients, cost"})
		return
	}
	createdRecipe, err := h.repo.Create(c.Request.Context(), recipe)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "Recipe creation failed!", "required": "title, making_time, serves, ingredients, cost"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Recipe successfully created!", "recipe": []model.Recipe{createdRecipe}})
}

func (h *recipeHandler) UpdateRecipe(c *gin.Context) {
	var recipe model.Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "Recipe update failed!", "required": "title, making_time, serves, ingredients, cost"})
		return
	}

	id := c.Param("id")
	updatedRecipe, err := h.repo.Update(c.Request.Context(), id, recipe)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "No Recipe found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Recipe successfully updated!", "recipe": []model.Recipe{updatedRecipe}})
}

func (h *recipeHandler) DeleteRecipe(c *gin.Context) {
	if err := h.repo.Delete(c.Request.Context(), c.Param("id")); err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "No Recipe found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Recipe successfully removed!"})
}
