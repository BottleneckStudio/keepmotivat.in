package models

import "os"

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
