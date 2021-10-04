package main

import "database/sql"

var db *sql.DB

func DbConnection() *sql.DB {
	if db != nil {
		return db
	}

	var err error

	db, err = sql.Open("sqlite3", "data.sqlite3")

	if err != nil {
		panic(err)
	}

	return db
}

func MakeMigrations() error {
	db := DbConnection()

	query := `CREATE TABLE IF NOT EXISTS todos (
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				title VARCHAR(64) NULL,
				description VARCHAR(200),
				is_complete BOOLEAN DEFAULT FALSE,
				created TIMESTAMP DEFAULT DATETIME
	);`

	_, err := db.Exec(query)

	if err != nil {
		return err
	}

	return nil
}
