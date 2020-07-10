package entity

import (
	"database/sql"
	"fmt"

	"github.com/desafios-job/import-data/domain/validators"
)

// Shopping struct that represents model
type Shopping struct {
	ID                  int64
	CustomerID          sql.NullString
	Private             sql.NullInt64
	Incomplete          sql.NullInt64
	LastDateShop        sql.NullTime
	AvgTicket           sql.NullFloat64
	LastTicketShop      sql.NullFloat64
	MostFrequentedStore sql.NullString
	LastStore           sql.NullString
}

// Shoppings is collection of shopping
type Shoppings []*Shopping

// NewShopping instantiating uninitialized objects
func NewShopping(customerID sql.NullString,
	private sql.NullInt64,
	incomplete sql.NullInt64,
	lastDateShop sql.NullTime,
	avgTicket sql.NullFloat64,
	lastTicketShop sql.NullFloat64,
	mostFrequentedStore sql.NullString,
	lastStore sql.NullString) *Shopping {
	return &Shopping{
		CustomerID:          customerID,
		Private:             private,
		Incomplete:          incomplete,
		LastDateShop:        lastDateShop,
		AvgTicket:           avgTicket,
		LastTicketShop:      lastTicketShop,
		MostFrequentedStore: mostFrequentedStore,
		LastStore:           lastStore,
	}
}

// Validate shopping
func (s *Shopping) Validate() (valid bool, errors []string) {

	const msg = "%s is invalid! [%s:%s]."

	validator := validators.Validator{}

	errors = []string{}

	if s.CustomerID.Valid && !validator.IsValidBrazilianID(s.CustomerID.String) {
		s.CustomerID.Valid = false
		message := fmt.Sprintf(msg, "Customer ID", "customer_id", s.CustomerID.String)
		errors = append(errors, message)
	}

	if s.MostFrequentedStore.Valid && !validator.IsValidCNPJ(s.MostFrequentedStore.String) {
		s.MostFrequentedStore.Valid = false
		message := fmt.Sprintf(msg, "CNPJ", "most_frequented_store", s.CustomerID.String)
		errors = append(errors, message)
	}

	if s.LastStore.Valid && !validator.IsValidCNPJ(s.LastStore.String) {
		s.LastStore.Valid = false
		message := fmt.Sprintf(msg, "CNPJ", "last_store", s.CustomerID.String)
		errors = append(errors, message)
	}

	return len(errors) == 0, errors
}
