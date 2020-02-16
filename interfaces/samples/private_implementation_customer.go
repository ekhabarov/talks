package customer

import "database/sql"

type CustomerRepo interface { // exported interface
	Get(int) *Customer
}

type repo struct { // unexported implementation
	db *sql.DB
}

func NewRepo(db *sql.DB) CustomerRepo {
	return repo{db: db}
}

func (repo) Get(id int) *Customer { return db.Select() }
