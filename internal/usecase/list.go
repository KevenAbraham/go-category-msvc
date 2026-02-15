package usecase

import (
	"github.com/KevenAbraham/go-category-msvc/internal/entities"
	"github.com/KevenAbraham/go-category-msvc/internal/repositories"
)

type listUsecase struct {
	repository repositories.ICategoryRepository
}

func NewListUsecase(repo repositories.ICategoryRepository) *listUsecase {
	return &listUsecase{
		repository: repo,
	}
}

func (u *listUsecase) List() ([]*entities.Category, error) {
	categories, err := u.repository.List()
	if err != nil {
		return nil, err
	}

	return categories, nil
}
