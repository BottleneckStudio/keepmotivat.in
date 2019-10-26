package models

import (
	"fmt"
)

// PostRepository handles the CRUD for post.
type PostRepository interface {
	Create(*Post) error
	Get(*Query) (*Post, error)
	GetAll(*Query) (*[]Post, error)
	Update(*Post) error
}

// DBPostRepository ...
type DBPostRepository struct {
	DB *DB
}

// NewDBPostRepository ...
func NewDBPostRepository(db *DB) PostRepository {
	return DBPostRepository{db}
}

// Create handles the creation of post.
func (repo DBPostRepository) Create(post *Post) error {
	tx := repo.DB.MustBegin()
	_, err := tx.NamedExec("INSERT INTO posts (body, caption, user_id, ctime, utime) VALUES (:body, :caption, :user_id, :ctime, :utime)", post)
	if err != nil {
		return err
	}
	return nil
}

// Get handles the fetching of a single post record.
func (repo DBPostRepository) Get(query *Query) (*Post, error) {
	post := Post{}

	queryString := "SELECT * FROM posts"

	if query.ByColumn != nil {
		queryString = queryString + fmt.Sprintf(" WHERE %v = %v", query.ByColumn.Column, query.ByColumn.Value)
	}

	err := repo.DB.Get(&post, queryString)
	if err != nil {
		return nil, err
	}

	return &post, nil
}

// GetAll handles the fetching ofpost records.
func (repo DBPostRepository) GetAll(query *Query) (*[]Post, error) {
	posts := []Post{}

	queryString := "SELECT * FROM posts"

	if query.ByColumn != nil {
		queryString = queryString + fmt.Sprintf(" WHERE %v = %v", query.ByColumn.Column, query.ByColumn.Value)
	}

	err := repo.DB.Select(&posts, queryString)
	if err != nil {
		return nil, err
	}

	return &posts, nil
}

// Update handles the updating of post.
func (repo DBPostRepository) Update(post *Post) error {
	tx := repo.DB.MustBegin()
	_, err := tx.NamedExec("UPDATE posts SET body = :body, caption = :caption, user_id = :user_id, ctime = :ctime, utime = :utime WHERE id = :id", post)
	if err != nil {
		return err
	}
	return nil
}
