package models

import "fmt"

// UserRepository handles the CRUD for user.
type UserRepository interface {
	Create(User) error
	Get(Query) (User, error)
	GetAll(Query) (Users, error)
	Update(User) error
}

// DBUserRepository ...
type DBUserRepository struct {
	DB *DB
}

// NewDBUserRepository ...
func NewDBUserRepository(db *DB) UserRepository {
	return DBUserRepository{db}
}

// Create handles the creation of user.
func (repo DBUserRepository) Create(user User) error {
	tx := repo.DB.MustBegin()
	_, err := tx.NamedExec("INSERT INTO users (email, password, ctime) VALUES (:email_address, :password, :ctime)", user)
	if err != nil {
		return err
	}
	return nil
}

// Get handles the fetching of a single user record.
func (repo DBUserRepository) Get(query Query) (User, error) {
	user := User{}

	queryString := "SELECT * FROM users"

	if query.ByColumn != nil {
		queryString = queryString + fmt.Sprintf(" WHERE %v = %v", query.ByColumn.Column, query.ByColumn.Value)
	}

	err := repo.DB.Get(&user, queryString)
	if err != nil {
		return user, err
	}

	return user, nil
}

// GetAll handles the fetching of user records.
func (repo DBUserRepository) GetAll(query Query) (Users, error) {
	users := Users{}

	queryString := "SELECT * FROM users"

	if query.ByColumn != nil {
		queryString = queryString + fmt.Sprintf(" WHERE %v = %v", query.ByColumn.Column, query.ByColumn.Value)
	}

	err := repo.DB.Select(&users, queryString)
	if err != nil {
		return nil, err
	}

	return users, nil
}

// Update handles the updating of user.
func (repo DBUserRepository) Update(user User) error {
	// tx := repo.DB.MustBegin()
	// _, err := tx.NamedExec("UPDATE p SET body = :body, caption = :caption, user_id = :user_id, ctime = :ctime, utime = :utime WHERE id = :id", user)
	// if err != nil {
	// 	return err
	// }
	return nil
}
