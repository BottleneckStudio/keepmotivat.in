package usecase

import "BottleneckStudio/keepmotivat.in/models"

// UpdatePost interface ...
type UpdatePost interface {
	UpdatePost(*models.Post) error
}

// UpdatePostUsecase ...
type UpdatePostUsecase struct {
	repo models.PostRepository
}

// NewUpdatePostUsecase ...
func NewUpdatePostUsecase(repo models.PostRepository) UpdatePost {
	return UpdatePostUsecase{
		repo: repo,
	}
}

// UpdatePost ...
func (uc UpdatePostUsecase) UpdatePost(post *models.Post) error {
	err := uc.repo.Update(post)
	return err
}
