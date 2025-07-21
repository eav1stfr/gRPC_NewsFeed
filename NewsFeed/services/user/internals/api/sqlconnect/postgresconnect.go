package sqlconnect

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"os"
	"utils"
)

func ConnectDB() (*sqlx.DB, error) {
	connectionString := os.Getenv("CONNECTION_STRING")
	db, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		return nil, utils.ConnectingToDbError
	}
	if err = db.Ping(); err != nil {
		return nil, utils.ConnectingToDbError
	}
	return db, nil
}
