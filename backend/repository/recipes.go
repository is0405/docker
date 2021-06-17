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
`, a.ID, a.Title, a.MakingTime, a.Serves, a.Ingridients, a.Cost, a.CreatedAt, a.UpdatedAt); err != nil {
		return nil, err
	}
	return &a, nil
}
