package repositories

import "github.com/KevenAbraham/go-category-msvc/internal/entities"

type ICategoryRepository interface {
	Save(category *entities.Category) error
	List() ([]*entities.Category, error)
}