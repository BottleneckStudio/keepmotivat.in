package main

import "github.com/BottleneckStudio/keepmotivat.in/models"

const (
	dbName = "keepmotivatin"
	dsn    = "root:@tcp(127.0.0.1:3306)/?charset=utf8mb4"
)

func main() {

	db, err := models.NewDatabase(dsn)
	if err != nil {
		// panic for now.
		panic(err.Error())
	}

	defer db.Close()

	db.Create(dbName)
	db.Use(dbName)
	db.Migrate()

	// Setup routes here later-on
	print("Hello World")
}
