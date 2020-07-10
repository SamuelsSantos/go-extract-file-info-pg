package service

import (
	"github.com/bdwilliams/go-jsonify/jsonify"
	"github.com/desafios-job/import-data/domain/entity"
	"github.com/desafios-job/import-data/domain/repository"
)

type shoppingApp struct {
	repo repository.ShoppingRepository
}

// ShoppingAppInterface shopping service layer
type ShoppingAppInterface interface {
	SaveMany(shoppings entity.Shoppings)
	GetAll() ([]string, error)
	Truncate() error
}

// SaveMany shoppings service layer
func (s *shoppingApp) SaveMany(shoppings entity.Shoppings) {
	s.repo.SaveMany(shoppings)
}

// GetAll shoppings service layer
func (s *shoppingApp) GetAll() ([]string, error) {
	rows, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	return jsonify.Jsonify(rows), nil
}

// GetAll Truncate service layer
func (s *shoppingApp) Truncate() error {
	return s.repo.Truncate()
}

// NewShoppingService instance
func NewShoppingService(repository repository.ShoppingRepository) ShoppingAppInterface {
	return &shoppingApp{
		repo: repository,
	}
}
