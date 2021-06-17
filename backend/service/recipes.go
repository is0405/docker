package service

import (
	"github.com/is0405/docker/dbutil"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/is0405/docker/model"
	"github.com/is0405/docker/repository"
	//"fmt"
)

type Recipes struct {
	db *sqlx.DB
}

func NewRecipes(db *sqlx.DB) *Recipes {
	return &Recipes{db}
}

func (a *Recipes) Create(rawRecipe *model.Recipes) (int64, error) {
	var createdId int64
	if err := dbutil.TXHandler(a.db, func(tx *sqlx.Tx) error {
		
		recipes, err := repository.AddRecipe(a.db, rawRecipe)	
		if err != nil {
			return err
		}
		
		if err := tx.Commit(); err != nil {
			return err
		}
		
		id, err := recipes.LastInsertId()
		
		if err != nil {
			return err
		}
		
		createdId = id
		return err
		
	}); err != nil {
		return 0, errors.Wrap(err, "failed recipe insert transaction")
	}
	return createdId, nil
}
