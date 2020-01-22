package usecase

import "BottleneckStudio/keepmotivat.in/models"

// CreateUser interface ...
type CreateUser interface {
	CreateUser(models.User) error
}

// CreateUserUsecase ...
type CreateUserUsecase struct {
	repo models.UserRepository
}

// NewCreateUserUsecase ...
func NewCreateUserUsecase(repo models.UserRepository) CreateUser {
	return CreateUserUsecase{
		repo: repo,
	}
}

// CreateUser ...
func (uc CreateUserUsecase) CreateUser(user models.User) error {
	err := uc.repo.Create(user)
	return err
}
