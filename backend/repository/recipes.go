package repository

import (
	"database/sql"

	"github.com/is0405/recipes/model"
	"github.com/jmoiron/sqlx"
)

func AddRecipe(db *sqlx.DB, mr *model.Recipes) (sql.Result, error) {
	return db.Exec(`
INSERT INTO recipes (title, making_time, serves, ingredients, cost)
VALUES (?, ?, ?, ?, ?)
`, mr.Title, mr.MakingTime, mr.Serves, mr.Ingridients, mr.Cost)
}
