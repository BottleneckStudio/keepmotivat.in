package usecase

import "BottleneckStudio/keepmotivat.in/models"

// CreatePost interface ...
type CreatePost interface {
	CreatePost(*models.Post) error
}

// CreatePostUsecase ...
type CreatePostUsecase struct {
	repo models.PostRepository
}

// NewCreatePostUsecase ...
func NewCreatePostUsecase(repo models.PostRepository) CreatePost {
	return CreatePostUsecase{
		repo: repo,
	}
}

// CreatePost ...
func (uc CreatePostUsecase) CreatePost(post *models.Post) error {
	err := uc.repo.Create(post)
	return err
}
