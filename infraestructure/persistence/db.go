package persistence

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/desafios-job/import-data/domain/repository"
)

// Repositories struct
type Repositories struct {
	Shopping      repository.ShoppingRepository
	Inconsistency repository.InconsistencyRepository
	db            *sql.DB
}

// NewRepositories struct
func NewRepositories(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) (*Repositories, error) {
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
	db, err := sql.Open("postgres", DBURL)
	if err != nil {
		panic(err)
	}

	return &Repositories{
		Shopping:      NewShoppingRepository(db),
		Inconsistency: NewInconsistencyRepository(db),
		db:            db,
	}, nil
}

//Close the  database connection
func (s *Repositories) Close() error {
	return s.db.Close()
}

// ArgumentValues type
type ArgumentValues = []interface{}

// GetStatementArgsIndex create array of fields args index
func GetStatementArgsIndex(numberOfColumns int, numberOfLines int) []string {

	lines := make([]string, numberOfLines)
	columns := make([]string, numberOfColumns)

	for line := range lines {

		start := line*numberOfColumns + 1
		for column := range columns {
			position := start + column
			columns[column] = fmt.Sprintf("$%d", position)
		}
		lines[line] = fmt.Sprintf("(%s)", strings.Join(columns, ","))
	}

	return lines
}

// TruncateTable {table} and restart identity
func TruncateTable(db *sql.DB, table string) error {
	sql := fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY RESTRICT", table)
	_, err := db.Exec(sql)

	if err != nil {
		log.Panic(err)
		return err
	}

	return err
}