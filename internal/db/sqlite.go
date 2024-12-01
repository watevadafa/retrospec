package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path/filepath"
)

var DB *sql.DB

func InitDB(dbPath string) error {
	var err error
	DB, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		return err
	}

	// Set connection pool settings
	DB.SetMaxOpenConns(1) // SQLite only supports one connection
	DB.SetMaxIdleConns(1)

	// Run migrations
	if err := runMigrations(); err != nil {
		return err
	}

	return DB.Ping()
}

func runMigrations() error {
	// Read migration files from migrations directory
	files, err := filepath.Glob("migrations/*.sql")
	if err != nil {
		return err
	}

	// Execute each migration file
	for _, file := range files {
		content, err := os.ReadFile(file)
		if err != nil {
			return err
		}

		// Execute the SQL migration
		_, err = DB.Exec(string(content))
		if err != nil {
			return err
		}
	}

	return nil
}
