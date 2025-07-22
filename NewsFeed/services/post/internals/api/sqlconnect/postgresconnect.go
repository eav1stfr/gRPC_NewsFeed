package sqlconnect

import (
	"github.com/jmoiron/sqlx"
	"os"
	"utils"
)

func ConnectDb() (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", os.Getenv("CONNECTION_STRING"))
	if err != nil {
		return nil, utils.ConnectingToDbError
	}

	if err = db.Ping(); err != nil {
		return nil, utils.ConnectingToDbError
	}

	return db, nil
}
