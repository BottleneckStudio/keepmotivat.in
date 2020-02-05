package usecase

import "BottleneckStudio/keepmotivat.in/models"

// GetUser interface ...
type GetUser interface {
	GetUser(models.Query) (*models.User, error)
	GetAllUsers(models.Query) (models.Users, error)
}

// GetUserUsecase ...
type GetUserUsecase struct {
	repo models.UserRepository
}

// NewGetUserUsecase ...
func NewGetUserUsecase(repo models.UserRepository) GetUser {
	return GetUserUsecase{
		repo: repo,
	}
}

// GetUser ...
func (uc GetUserUsecase) GetUser(query models.Query) (*models.User, error) {
	user, err := uc.repo.Get(query)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetAllUsers ...
func (uc GetUserUsecase) GetAllUsers(query models.Query) (models.Users, error) {
	users, err := uc.repo.GetAll(query)
	if err != nil {
		return nil, err
	}
	return users, nil
}
