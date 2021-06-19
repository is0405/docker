package model

type Recipes struct {
	ID          int    `db:"id" json:"id"`
	Title       string `db:"title" json:"title"`
	MakingTime  string `db:"making_time" json:"making_time"`
	Serves      string `db:"serves" json:"serves"`
	Ingridients string `db:"ingredients" json:"ingredients"`
	Cost        int    `db:"cost" json:"cost"`
	CreatedAt   string `db:"created_at" json:"created_at"`
	UpdatedAt   string `db:"updated_at" json:"updated_at"`
}
