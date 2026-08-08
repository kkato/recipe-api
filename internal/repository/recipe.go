package repository

import (
	"context"
	"database/sql"

	"github.com/kkato/recipe-api/internal/model"
)

type RecipeRepository interface {
	Create(ctx context.Context, recipe model.Recipe) (model.Recipe, error)
	GetAll(ctx context.Context) ([]model.Recipe, error)
	GetByID(ctx context.Context, id int) (model.Recipe, error)
	Update(ctx context.Context, id int, recipe model.Recipe) (model.Recipe, error)
	Delete(ctx context.Context, id int) error
}

type recipeRepository struct {
	db *sql.DB
}

func NewRecipeRepository(db *sql.DB) RecipeRepository {
	return &recipeRepository{db: db}
}

func (r *recipeRepository) Create(ctx context.Context, recipe model.Recipe) (model.Recipe, error) {
	query := "INSERT INTO recipes (title, making_time, serves, ingredients, cost) VALUES (?, ?, ?, ?, ?)"
	result, err := r.db.ExecContext(ctx, query, recipe.Title, recipe.MakingTime, recipe.Serves, recipe.Ingredients, recipe.Cost)
	if err != nil {
		return model.Recipe{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return model.Recipe{}, err
	}

	return r.GetByID(ctx, int(id))
}

func (r *recipeRepository) GetAll(ctx context.Context) ([]model.Recipe, error) {
	query := "SELECT id, title, making_time, serves, ingredients, cost, created_at, updated_at FROM recipes"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	recipes := []model.Recipe{}
	for rows.Next() {
		var recipe model.Recipe
		if err := rows.Scan(&recipe.ID, &recipe.Title, &recipe.MakingTime, &recipe.Serves, &recipe.Ingredients, &recipe.Cost, &recipe.CreatedAt, &recipe.UpdatedAt); err != nil {
			return nil, err
		}
		recipes = append(recipes, recipe)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return recipes, nil
}

func (r *recipeRepository) GetByID(ctx context.Context, id int) (model.Recipe, error) {
	query := "SELECT id, title, making_time, serves, ingredients, cost, created_at, updated_at FROM recipes WHERE id = ?"
	row := r.db.QueryRowContext(ctx, query, id)

	var recipe model.Recipe
	if err := row.Scan(&recipe.ID, &recipe.Title, &recipe.MakingTime, &recipe.Serves, &recipe.Ingredients, &recipe.Cost, &recipe.CreatedAt, &recipe.UpdatedAt); err != nil {
		return model.Recipe{}, err
	}

	return recipe, nil
}

func (r *recipeRepository) Update(ctx context.Context, id int, recipe model.Recipe) (model.Recipe, error) {
	query := "UPDATE recipes SET title = ?, making_time = ?, serves = ?, ingredients = ?, cost = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?"
	result, err := r.db.ExecContext(ctx, query, recipe.Title, recipe.MakingTime, recipe.Serves, recipe.Ingredients, recipe.Cost, id)
	if err != nil {
		return model.Recipe{}, err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return model.Recipe{}, err
	}
	if affected == 0 {
		return model.Recipe{}, sql.ErrNoRows
	}

	return r.GetByID(ctx, id)
}

func (r *recipeRepository) Delete(ctx context.Context, id int) error {
	query := "DELETE FROM recipes WHERE id = ?"
	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
