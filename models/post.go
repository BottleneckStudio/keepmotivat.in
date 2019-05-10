package models

// Post represents one post in our feed
type Post struct {
	ID      int64
	Body    string `db:"body"`
	Caption string `db:"caption"`
	UserID  int64  `db:"user_id"`
	Ctime   int64  `db:"ctime"`
	Utime   int64  `db:"utime"`
}
