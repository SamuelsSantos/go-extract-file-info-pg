package repository

import "github.com/desafios-job/import-data/domain/entity"

// InconsistencyRepository repository
type InconsistencyRepository interface {
	SaveMany(inconsistencies entity.Inconsistencies)
	GetAll() (entity.Inconsistencies, error)
	Truncate() error
}
