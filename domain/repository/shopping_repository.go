package repository

import "github.com/desafios-job/import-data/domain/entity"

// ShoppingRepository repository
type ShoppingRepository interface {
	SaveMany(shoppings entity.Shoppings)
	GetAll() (entity.Shoppings, error)
	Truncate() error
}
