package models

// ByID ...
type ByID int

// ByColumn ...
type ByColumn struct {
	Column string
	Value  interface{}
}

// Query ...
type Query struct {
	*ByID
	*ByColumn
}
