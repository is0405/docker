package repository

import (
	"database/sql"

	"github.com/is0405/docker/model"
	"github.com/jmoiron/sqlx"
)

func AddRecipe(db *sqlx.DB, mr *model.Recipes) (sql.Result, error) {
	return db.Exec(`
INSERT INTO recipes (title, making_time, serves, ingredients, cost)
VALUES (?, ?, ?, ?, ?)
`, mr.Title, mr.MakingTime, mr.Serves, mr.Ingridients, mr.Cost)
}

func GetRecipe(db *sqlx.DB, id int) (*model.Recipes, error) {
	a := model.Recipes{}
	if err := db.Get(&a, `
SELECT id, title, making_time, serves, ingredients, cost, created_at, updated_at
FROM recipes
WHERE id = ?;
`, id); err != nil {
		return nil, err
	}
	return &a, nil
}

func RemoveRecipe(db *sqlx.DB, id int) (sql.Result, error) {
	return db.Exec(`
DELETE FROM recipes WHERE id = ?;
`, id)
}

func GetRecipesList(db *sqlx.DB) ([]*model.Recipes, error) {
	var recipes []*model.Recipes
	if err := db.Select(&recipes, `
SELECT id, title, making_time, serves, ingredients, cost, created_at, updated_at
FROM recipes
	`); err != nil {
		return nil, err
	}

	return recipes, nil
}

func UpdateRecipe(db *sqlx.DB, mr *model.Recipes) (sql.Result, error) {
	return db.Exec(`
UPDATE recipes SET title = ?, making_time = ?, serves = ?, ingredients = ?, cost = ? WHERE id = ?;
`, mr.Title, mr.MakingTime, mr.Serves, mr.Ingridients, mr.Cost, mr.ID)
}
