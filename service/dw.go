package service

import (
	"github.com/desafios-job/import-data/domain/entity"
	"github.com/desafios-job/import-data/infraestructure/files"
)

// DW Data Warehouse interface
type DW interface {
	Clean() error
	Extract() (*files.Rows, error)
	Transform(rows *files.Rows) (*entity.Shoppings, *entity.Inconsistencies)
	SaveShoppings(shoppings *entity.Shoppings)
	SaveInconsistencies(inconsistencies *entity.Inconsistencies)
}
