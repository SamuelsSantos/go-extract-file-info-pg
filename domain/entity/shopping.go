package entity

import (
	"database/sql"
	"fmt"

	"github.com/desafios-job/import-data/domain/validators"
)

// Shopping struct that represents model
type Shopping struct {
	ID                  int64           `json:"-"`
	CustomerID          sql.NullString  `json:"customer_id"`
	Private             sql.NullInt64   `json:"private"`
	Incomplete          sql.NullInt64   `json:"incomplete"`
	LastDateShop        sql.NullTime    `json:"last_shop"`
	AvgTicket           sql.NullFloat64 `json:"avg_ticket"`
	LastTicketShop      sql.NullFloat64 `json:"last_ticket_shop"`
	MostFrequentedStore sql.NullString  `json:"most_frequented_store"`
	LastStore           sql.NullString  `json:"last_store"`
}

// Shoppings is collection of shopping
type Shoppings []*Shopping

// Validate shopping
func (s *Shopping) Validate() (valid bool, errors []string) {

	const msg = "%s is invalid! [%s:%s]."

	errors = []string{}

	if s.CustomerID.Valid && !validators.IsValidBrazilianID(s.CustomerID.String) {
		s.CustomerID.Valid = false
		message := fmt.Sprintf(msg, "Customer ID", "customer_id", s.CustomerID.String)
		errors = append(errors, message)
	}

	if s.MostFrequentedStore.Valid && !validators.IsValidCNPJ(s.MostFrequentedStore.String) {
		s.MostFrequentedStore.Valid = false
		message := fmt.Sprintf(msg, "CNPJ", "most_frequented_store", s.CustomerID.String)
		errors = append(errors, message)
	}

	if s.LastStore.Valid && !validators.IsValidCNPJ(s.LastStore.String) {
		s.LastStore.Valid = false
		message := fmt.Sprintf(msg, "CNPJ", "last_store", s.CustomerID.String)
		errors = append(errors, message)
	}

	return len(errors) == 0, errors
}
