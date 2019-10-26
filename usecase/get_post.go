package usecase

import "BottleneckStudio/keepmotivat.in/models"

// GetPost interface ...
type GetPost interface {
	GetPost(*models.Query) (*models.Post, error)
	GetAllPosts(*models.Query) (*[]models.Post, error)
}

// GetPostUsecase ...
type GetPostUsecase struct {
	repo models.PostRepository
}

// NewGetPostUsecase ...
func NewGetPostUsecase(repo models.PostRepository) GetPost {
	return GetPostUsecase{
		repo: repo,
	}
}

// GetPost ...
func (uc GetPostUsecase) GetPost(query *models.Query) (*models.Post, error) {
	post := &models.Post{}
	post, err := uc.repo.Get(query)
	if err != nil {
		return nil, err
	}
	return post, nil
}

// GetAllPosts ...
func (uc GetPostUsecase) GetAllPosts(query *models.Query) (*[]models.Post, error) {
	posts := &[]models.Post{}
	posts, err := uc.repo.GetAll(query)
	if err != nil {
		return nil, err
	}
	return posts, nil
}
