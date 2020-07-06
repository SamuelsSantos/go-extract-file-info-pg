package service

import (
	"github.com/desafios-job/import-data/domain/entity"
	"github.com/desafios-job/import-data/domain/repository"
)

type inconsistencyApp struct {
	repo repository.InconsistencyRepository
}

// InconsistencyAppInterface inconsistency service layer
type InconsistencyAppInterface interface {
	SaveMany(inconsistencies entity.Inconsistencies)
	GetAll() (entity.Inconsistencies, error)
	Truncate() error
}

// SaveMany inconsistency service layer
func (i *inconsistencyApp) SaveMany(inconsistencies entity.Inconsistencies) {
	i.repo.SaveMany(inconsistencies)
}

// GetAll inconsistency service layer
func (i *inconsistencyApp) GetAll() (entity.Inconsistencies, error) {
	return i.repo.GetAll()
}

// Truncate inconsistency service layer
func (i *inconsistencyApp) Truncate() error {
	return i.repo.Truncate()
}

// NewInconsistencyApp instance
func NewInconsistencyApp(repository repository.InconsistencyRepository) InconsistencyAppInterface {
	return &inconsistencyApp{
		repo: repository,
	}
}
