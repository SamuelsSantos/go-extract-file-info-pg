package repository

import (
	"database/sql"

	"github.com/desafios-job/import-data/domain/entity"
)

// InconsistencyRepository repository
type InconsistencyRepository interface {
	SaveMany(inconsistencies entity.Inconsistencies)
	GetAll() (*sql.Rows, error)
	Truncate() error
}
