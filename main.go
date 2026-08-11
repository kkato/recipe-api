package main

import (
	"log"
	"net/http"

	"github.com/kkato/recipe-api/internal/database"
	"github.com/kkato/recipe-api/internal/handler"
	"github.com/kkato/recipe-api/internal/repository"
)

func main() {
	db, err := database.New("recipe.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := repository.NewRecipeRepository(db)
	h := handler.NewRecipeHandler(repo)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /recipes", h.Create)
	mux.HandleFunc("GET /recipes", h.GetAll)
	mux.HandleFunc("GET /recipes/{id}", h.GetByID)
	mux.HandleFunc("PATCH /recipes/{id}", h.Update)
	mux.HandleFunc("DELETE /recipes/{id}", h.Delete)

	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
