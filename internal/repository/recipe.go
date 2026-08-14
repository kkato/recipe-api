package repository

import (
	"context"
	"database/sql"
	"errors"
	"strconv"

	"github.com/kkato/recipe-api/internal/model"
)

var ErrRecipeNotFound = errors.New("recipe not found")

type RecipeRepository interface {
	GetAll(ctx context.Context) ([]model.Recipe, error)
	GetByID(ctx context.Context, id string) (model.Recipe, error)
	GetByTitle(ctx context.Context, title string) (model.Recipe, error)
	Create(ctx context.Context, recipe model.Recipe) (model.Recipe, error)
	Update(ctx context.Context, id string, recipe model.Recipe) (model.Recipe, error)
	Delete(ctx context.Context, id string) error
}

type recipeRepository struct {
	db *sql.DB
}

func NewRecipeRepository(db *sql.DB) RecipeRepository {
	return &recipeRepository{db: db}
}

func (r *recipeRepository) GetAll(ctx context.Context) ([]model.Recipe, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT * FROM recipes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	recipes := make([]model.Recipe, 0)
	for rows.Next() {
		var recipe model.Recipe
		if err := rows.Scan(&recipe.ID, &recipe.Title, &recipe.MakingTime, &recipe.Serves, &recipe.Ingredients, &recipe.Cost, &recipe.CreatedAt, &recipe.UpdatedAt); err != nil {
			return nil, err
		}
		recipes = append(recipes, recipe)
	}
	return recipes, rows.Err()
}

func (r *recipeRepository) GetByID(ctx context.Context, id string) (model.Recipe, error) {
	var recipe model.Recipe
	err := r.db.QueryRowContext(ctx, "SELECT * FROM recipes WHERE id = ?", id).
		Scan(&recipe.ID, &recipe.Title, &recipe.MakingTime, &recipe.Serves, &recipe.Ingredients, &recipe.Cost, &recipe.CreatedAt, &recipe.UpdatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return model.Recipe{}, ErrRecipeNotFound
	}
	if err != nil {
		return model.Recipe{}, err
	}
	return recipe, nil
}

func (r *recipeRepository) GetByTitle(ctx context.Context, title string) (model.Recipe, error) {
	var recipe model.Recipe
	err := r.db.QueryRowContext(ctx, "SELECT * FROM recipes WHERE title = ?", title).
		Scan(&recipe.ID, &recipe.Title, &recipe.MakingTime, &recipe.Serves, &recipe.Ingredients, &recipe.Cost, &recipe.CreatedAt, &recipe.UpdatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return model.Recipe{}, ErrRecipeNotFound
	}
	if err != nil {
		return model.Recipe{}, err
	}
	return recipe, nil
}

func (r *recipeRepository) Create(ctx context.Context, recipe model.Recipe) (model.Recipe, error) {
	res, err := r.db.ExecContext(ctx, "INSERT INTO recipes (title, making_time, serves, ingredients, cost) VALUES (?, ?, ?, ?, ?)",
		recipe.Title, recipe.MakingTime, recipe.Serves, recipe.Ingredients, recipe.Cost)
	if err != nil {
		return model.Recipe{}, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return model.Recipe{}, err
	}

	return r.GetByID(ctx, strconv.FormatInt(id, 10))
}

func (r *recipeRepository) Update(ctx context.Context, id string, recipe model.Recipe) (model.Recipe, error) {
	res, err := r.db.ExecContext(ctx, "UPDATE recipes SET title = ?, making_time = ?, serves = ?, ingredients = ?, cost = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?",
		recipe.Title, recipe.MakingTime, recipe.Serves, recipe.Ingredients, recipe.Cost, id)
	if err != nil {
		return model.Recipe{}, err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return model.Recipe{}, err
	}
	if rows == 0 {
		return model.Recipe{}, ErrRecipeNotFound
	}

	return r.GetByID(ctx, id)
}

func (r *recipeRepository) Delete(ctx context.Context, id string) error {
	res, err := r.db.ExecContext(ctx, "DELETE FROM recipes WHERE id = ?", id)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return ErrRecipeNotFound
	}
	return nil
}
