package models

import (

	// Importing this one because sqlx asks for this one.
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// DB defines the database.
// It uses the package jmoiron/sqlx.
type DB struct {
	*sqlx.DB
}

// Tx is the transaction for our database.
// E.g: inserting, updating, deleting.
type Tx struct {
	Tx *sqlx.Tx
}

// NewDatabase returns the pointer to sqlx.DB struct.
func NewDatabase(dataSourceName string) (*DB, error) {
	sqlDB, dbErr := sqlx.Connect("mysql", dataSourceName)

	if dbErr != nil {
		return nil, dbErr
	}

	return &DB{
		sqlDB,
	}, nil
}

// Create the database if not exists.
func (db *DB) Create(dbName string) {
	db.MustExec("CREATE DATABASE IF NOT EXISTS `" + dbName + "` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;")
	db.MustBegin()
}

// Use selects the given database to operate with.
func (db *DB) Use(dbName string) {
	db.MustExec("USE " + dbName)
}

// Begin starts and returns a new transaction.
func (db *DB) Begin() (*Tx, error) {
	return &Tx{
		Tx: db.DB.MustBegin(),
	}, nil
}

// Migrate the tables.
func (db *DB) Migrate() {
	schemas := schemas()
	for _, schema := range schemas {
		db.MustExec(schema)
	}
	print("Database Migrated Successfully!")
}

func schemas() []string {

	return []string{
		`
		CREATE TABLE IF NOT EXISTS users (
			id bigint NOT NULL AUTO_INCREMENT,
			email varchar(50) NOT NULL,
			password varchar(255) NOT NULL,
			ctime bigint,
			utime bigint,
			PRIMARY KEY (id)
		);`,
		`
		CREATE TABLE IF NOT EXISTS posts (
			id bigint NOT NULL AUTO_INCREMENT,
			user_id bigint NOT NULL,
			body text,
			caption varchar(255) NULL,
			ctime bigint,
			utime bigint,
			PRIMARY KEY (id),
			FOREIGN KEY (user_id) REFERENCES users(id)
		);`,
	}
}
