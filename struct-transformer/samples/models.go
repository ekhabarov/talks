package models

type Product struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}
