package persistence

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/desafios-job/import-data/domain/entity"
)

// InconsistencyRepo struct
type InconsistencyRepo struct {
	db *sql.DB
}

// NewInconsistencyRepository new repository
func NewInconsistencyRepository(db *sql.DB) *InconsistencyRepo {
	return &InconsistencyRepo{db}
}

// SaveMany inconsistency
func (r *InconsistencyRepo) SaveMany(incosistencies entity.Inconsistencies) {

	sqlInsert := `
	INSERT INTO public.inconsistency
		(filename, error_message)
	VALUES %s`

	argumentIndexes := GetStatementArgsIndex(2, len(incosistencies))
	argumentValues := getStatementIncosistencyValues(incosistencies)

	statement := fmt.Sprintf(sqlInsert, strings.Join(argumentIndexes, ","))

	tx, err := r.db.Begin()
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	_, err = r.db.Exec(statement, argumentValues...)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	tx.Commit()

}

// GetAll incosistencies
func (r *InconsistencyRepo) GetAll() (*sql.Rows, error) {
	return FindAll(r.db, "select * from inconsistency")
}

// Truncate clean database and restar identity
func (r *InconsistencyRepo) Truncate() error {
	return TruncateTable(r.db, "inconsistency")
}

func getStatementIncosistencyValues(inconsistencies []*entity.Inconsistency) ArgumentValues {

	values := []interface{}{}
	for _, inconsistency := range inconsistencies {
		values = append(values, inconsistency.FileName)
		values = append(values, inconsistency.ErrorMessage)
	}

	return values
}
