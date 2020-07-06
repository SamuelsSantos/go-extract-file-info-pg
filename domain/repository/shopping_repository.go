package repository

import (
	"database/sql"

	"github.com/desafios-job/import-data/domain/entity"
)

// ShoppingRepository repository
type ShoppingRepository interface {
	SaveMany(shoppings entity.Shoppings)
	GetAll() (*sql.Rows, error)
	Truncate() error
}
