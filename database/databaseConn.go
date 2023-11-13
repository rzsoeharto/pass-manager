package database

import (
	"database/sql"
	"pass-manager/logger"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() *sql.DB {
	db, dbErr := sql.Open("sqlite3", "vault.db")

	if dbErr != nil {
		logger.Logger.Println(dbErr)
		return nil
	}

	return db
}
