package repositories

import (
	"database/sql"
)

type newsOutletRepository struct {
	connection *sql.DB
}

func NewNewsOutletRepository(connection *sql.DB) newsOutletRepository {
	return newsOutletRepository{
		connection: connection,
	}
}
