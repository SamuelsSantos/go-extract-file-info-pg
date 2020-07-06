package entity

import "database/sql"

// Inconsistency struct that represents model
type Inconsistency struct {
	ID           int64
	FileName     sql.NullString
	ErrorMessage sql.NullString
}

// Inconsistencies is collection of inconsistency
type Inconsistencies []*Inconsistency
