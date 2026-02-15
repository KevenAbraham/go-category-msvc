package usecase

import (
	"github.com/KevenAbraham/go-category-msvc/internal/entities"
	"github.com/KevenAbraham/go-category-msvc/internal/repositories"
)

type createUsecase struct {
	repository repositories.ICategoryRepository
}

func NewCreateUsecase(repo repositories.ICategoryRepository) *createUsecase {
	return &createUsecase{
		repository: repo,
	}
}

func (u *createUsecase) Create(name string) error {
	category, err := entities.NewCategory(name)
	if err != nil {
		return err
	}

	err = u.repository.Save(category)
	if err != nil {
		return err
	}

	return nil
}
