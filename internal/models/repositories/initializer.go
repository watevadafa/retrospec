package repositories

import (
	"database/sql"
)

var (
	UserRepo *UserRepository
)

func InitializeRepositories(db *sql.DB) {
	UserRepo = NewUserRepository(db)
	// Initialize other repositories here as they're added
}
