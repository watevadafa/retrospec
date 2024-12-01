package config

import "os"

func GetDBPath() string {
	return os.Getenv("DB_PATH")
}
