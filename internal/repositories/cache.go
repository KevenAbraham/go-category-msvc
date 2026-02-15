package repositories

import "github.com/KevenAbraham/go-category-msvc/internal/entities"

type cache struct {
	db []*entities.Category
}

func NewCache() *cache {
	return &cache{
		db: make([]*entities.Category, 0),
	}
}

func(r *cache) Save(category *entities.Category) error {
	r.db = append(r.db, category)

	return nil
}

func(r *cache) List() ([]*entities.Category, error) {
	return r.db, nil
}

