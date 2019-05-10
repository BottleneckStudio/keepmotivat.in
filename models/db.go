package models

import "os"

var dbhost = getenvWithDefault("DB_HOST", "127.0.0.1")

func init() {
	// Initialize connection to db here
}

func getenvWithDefault(key string, def string) string {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	return v
}
