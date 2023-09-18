package model

import (
	"database/sql"
	"log"
	"time"

	"rest-api/config"

	_ "github.com/mattn/go-sqlite3"
)

type Recipe struct {
	ID          string    `json:"id"`
	Title       string    `json:"title" binding:"required"`
	MakingTime  string    `json:"making_time" binding:"required"`
	Serves      string    `json:"serves" binding:"required"`
	Ingredients string    `json:"ingredients" binding:"required"`
	Cost        int       `json:"cost" binding:"required"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

var Db *sql.DB
var err error

func init() {
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatalln(err)
	}

	dropTableSQL := `DROP TABLE IF EXISTS recipes`
	_, err = Db.Exec(dropTableSQL)
	if err != nil {
		log.Fatalln(err)
	}

	createTableSQL := `CREATE TABLE IF NOT EXISTS recipes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		making_time TEXT NOT NULL,
		serves TEXT NOT NULL,
		ingredients TEXT NOT NULL,
		cost INTEGER NOT NULL,
		created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP)`
	_, err = Db.Exec(createTableSQL)
	if err != nil {
		log.Fatalln(err)
	}

	insertSQL1 := `INSERT INTO recipes (title, making_time, serves, ingredients, cost) 
		VALUES ('チキンカレー', '45分', '4人', '玉ねぎ,肉,スパイス', 1000)`
	_, err = Db.Exec(insertSQL1)
	if err != nil {
		log.Fatalln(err)
	}

	insertSQL2 := `INSERT INTO recipes (title, making_time, serves, ingredients, cost) 
		VALUES ('オムライス', '30分', '2人', '玉ねぎ,卵,スパイス,醤油', 700)`
	_, err = Db.Exec(insertSQL2)
	if err != nil {
		log.Fatalln(err)
	}
}

func GetRecipes() (recipes []Recipe) {
	rows, err := Db.Query("SELECT * FROM recipes")
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var recipe Recipe
		err := rows.Scan(&recipe.ID, &recipe.Title, &recipe.MakingTime, &recipe.Serves, &recipe.Ingredients, &recipe.Cost, &recipe.CreatedAt, &recipe.UpdatedAt)
		if err != nil {
			log.Fatalln(err)
		}
		recipes = append(recipes, recipe)
	}
	rows.Close()
	return recipes
}

func GetRecipe(id string) (recipes []Recipe) {
	var recipe Recipe
	row := Db.QueryRow("SELECT * FROM recipes WHERE id = ?", id)
	err := row.Scan(&recipe.ID, &recipe.Title, &recipe.MakingTime, &recipe.Serves, &recipe.Ingredients, &recipe.Cost, &recipe.CreatedAt, &recipe.UpdatedAt)
	if err != nil {
		log.Fatalln(err)
	}
	recipes = append(recipes, recipe)
	return recipes
}

func GetRecipeByTitle(title string) (recipes []Recipe) {
	var recipe Recipe
	row := Db.QueryRow("SELECT * FROM recipes WHERE title = ?", title)
	err := row.Scan(&recipe.ID, &recipe.Title, &recipe.MakingTime, &recipe.Serves, &recipe.Ingredients, &recipe.Cost, &recipe.CreatedAt, &recipe.UpdatedAt)
	if err != nil {
		log.Fatalln(err)
	}
	recipes = append(recipes, recipe)
	return recipes
}

func CreateRecipe(recipe Recipe) (err error) {
	_, err = Db.Exec("INSERT INTO recipes (title, making_time, serves, ingredients, cost) VALUES (?, ?, ?, ?, ?)", recipe.Title, recipe.MakingTime, recipe.Serves, recipe.Ingredients, recipe.Cost)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func UpdateRecipe(recipe Recipe) (err error) {
	_, err = Db.Exec("UPDATE recipes SET title = ?, making_time = ?, serves = ?, ingredients = ?, cost = ? WHERE id = ?", recipe.Title, recipe.MakingTime, recipe.Serves, recipe.Ingredients, recipe.Cost, recipe.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func DeleteRecipe(id string) (err error) {
	_, err = Db.Exec("DELETE FROM recipes WHERE id = ?", id)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
