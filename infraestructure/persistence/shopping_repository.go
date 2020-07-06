package persistence

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/desafios-job/import-data/domain/entity"
)

// ShoppingRepo struct
type ShoppingRepo struct {
	db *sql.DB
}

// NewShoppingRepository new repository
func NewShoppingRepository(db *sql.DB) *ShoppingRepo {
	return &ShoppingRepo{db}
}

// SaveMany shoppings
func (r *ShoppingRepo) SaveMany(shoppings entity.Shoppings) {

	sqlInsert := `
	INSERT INTO public.shopping
		(customer_id, private, incomplete, last_shop, avg_ticket, last_ticket_shop, most_frequented_store, last_store)
	VALUES %s`

	argumentIndexes := GetStatementArgsIndex(8, len(shoppings))
	argumentValues := getStatementShoppingValues(shoppings)

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

// GetAll shoppings
func (r *ShoppingRepo) GetAll() (*sql.Rows, error) {
	return FindAll(r.db, "select * from shopping")
}

// Truncate clean database and restar identity
func (r *ShoppingRepo) Truncate() error {
	return TruncateTable(r.db, "shopping")
}

func getStatementShoppingValues(shoppings []*entity.Shopping) ArgumentValues {

	values := []interface{}{}
	for _, shopping := range shoppings {
		values = append(values, shopping.CustomerID)
		values = append(values, shopping.Private)
		values = append(values, shopping.Incomplete)
		values = append(values, shopping.LastDateShop)
		values = append(values, shopping.AvgTicket)
		values = append(values, shopping.LastTicketShop)
		values = append(values, shopping.MostFrequentedStore)
		values = append(values, shopping.LastStore)
	}

	return values
}
