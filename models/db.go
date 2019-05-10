package models

import (
	"log"
	"os"
)

var dbhost = getenvWithDefault("DB_HOST", "127.0.0.1")

func init() {
	// Initialize connection to db here
	log.Printf("db host is: %s\n", dbhost)
}

func getenvWithDefault(key string, def string) string {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	return v
}
