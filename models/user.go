package models

// User represents the user for the keepmotivat.in application.
type User struct {
	ID           int64
	EmailAddress string `db:"email_address"`
	Password     string `db:"password"`
	Ctime        int64  `db:"ctime"`
	Utime        int64  `db:"utime"`
}

// Users repsents a slice of user.
type Users []User
