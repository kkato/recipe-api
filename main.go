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
	router.GET("/recipes", h.GetAll)
	router.GET("/recipes/:id", h.GetByID)
	router.POST("/recipes", h.Create)
	router.PATCH("/recipes/:id", h.Update)
	router.DELETE("/recipes/:id", h.Delete)
	router.Run(":8080")
}
