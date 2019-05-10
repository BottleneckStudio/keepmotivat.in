package models

// Post represents one post in our feed
type Post struct {
	ID      int64
	Body    string
	Caption string
	UserID  int64
	Ctime   int64
	Utime   int64
}
