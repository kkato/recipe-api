package main

import (
	"log"

	"github.com/kkato/recipe-api/internal/database"
	"github.com/kkato/recipe-api/internal/handler"
	"github.com/kkato/recipe-api/internal/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.New()
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	repo := repository.NewRecipeRepository(db)
	h := handler.NewRecipeHandler(repo)

	router := gin.Default()
	router.GET("/recipes", h.GetRecipes)
	router.GET("/recipes/:id", h.GetRecipe)
	router.POST("/recipes", h.CreateRecipe)
	router.PATCH("/recipes/:id", h.UpdateRecipe)
	router.DELETE("/recipes/:id", h.DeleteRecipe)
	router.Run(":8080")
}
