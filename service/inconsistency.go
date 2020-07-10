package service

import (
	"github.com/bdwilliams/go-jsonify/jsonify"
	"github.com/desafios-job/import-data/domain/entity"
	"github.com/desafios-job/import-data/domain/repository"
)

type inconsistencyApp struct {
	repo repository.InconsistencyRepository
}

// InconsistencyAppInterface inconsistency service layer
type InconsistencyAppInterface interface {
	SaveMany(inconsistencies entity.Inconsistencies)
	GetAll() ([]string, error)
	Truncate() error
}

// SaveMany inconsistency service layer
func (i *inconsistencyApp) SaveMany(inconsistencies entity.Inconsistencies) {
	i.repo.SaveMany(inconsistencies)
}

// GetAll inconsistency service layer
func (i *inconsistencyApp) GetAll() ([]string, error) {

	rows, err := i.repo.GetAll()
	if err != nil {
		return nil, err
	}

	return jsonify.Jsonify(rows), nil
}

// Truncate inconsistency service layer
func (i *inconsistencyApp) Truncate() error {
	return i.repo.Truncate()
}

// NewInconsistencyService instance
func NewInconsistencyService(repository repository.InconsistencyRepository) InconsistencyAppInterface {
	return &inconsistencyApp{
		repo: repository,
	}
}
