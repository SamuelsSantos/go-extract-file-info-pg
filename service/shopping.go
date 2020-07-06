package service

import (
	"github.com/desafios-job/import-data/domain/entity"
	"github.com/desafios-job/import-data/domain/repository"
)

type shoppingApp struct {
	repo repository.ShoppingRepository
}

// ShoppingAppInterface shopping service layer
type ShoppingAppInterface interface {
	SaveMany(shoppings entity.Shoppings)
	GetAll() (entity.Shoppings, error)
	Truncate() error
}

// SaveMany shoppings service layer
func (s *shoppingApp) SaveMany(shoppings entity.Shoppings) {
	s.repo.SaveMany(shoppings)
}

// GetAll shoppings service layer
func (s *shoppingApp) GetAll() (entity.Shoppings, error) {
	return s.repo.GetAll()
}

// GetAll Truncate service layer
func (s *shoppingApp) Truncate() error {
	return s.repo.Truncate()
}

// NewShoppingApp instance
func NewShoppingApp(repository repository.ShoppingRepository) ShoppingAppInterface {
	return &shoppingApp{
		repo: repository,
	}
}
