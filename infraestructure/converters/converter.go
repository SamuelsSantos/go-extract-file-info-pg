package converters

import (
	"database/sql"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/klassmann/cpfcnpj"
)

func isEmpty(value string) bool {
	return len(value) == 0 || value == "NULL"
}

// ReplaceDecimalSeparator replace from ',' to '.'
func ReplaceDecimalSeparator(value string) string {
	return strings.Replace(value, ",", ".", -1)
}

// CleanMaskDocument remove non numeric characters
func CleanMaskDocument(value string) string {
	return cpfcnpj.Clean(value)
}

// NewNullString convert string to NullString
func NewNullString(value string) sql.NullString {
	if isEmpty(value) {
		return sql.NullString{Valid: false}
	}
	return sql.NullString{
		String: value,
		Valid:  true,
	}
}

// NewNullInt convert string to NullInt64
func NewNullInt(value string) sql.NullInt64 {
	if isEmpty(value) {
		return sql.NullInt64{}
	}

	var result int64
	var err error

	if result, err = strconv.ParseInt(value, 10, 64); err != nil {
		log.Fatal(err)
	}

	return sql.NullInt64{
		Int64: result,
		Valid: true,
	}
}

// NewNullFloat convert string to NullFloat64
func NewNullFloat(value string) sql.NullFloat64 {
	if isEmpty(value) {
		return sql.NullFloat64{}
	}

	var result float64
	var err error

	value = ReplaceDecimalSeparator(value)
	if result, err = strconv.ParseFloat(value, 64); err != nil {
		log.Fatal(err)
		return sql.NullFloat64{}
	}

	return sql.NullFloat64{
		Float64: result,
		Valid:   true,
	}
}

// NewNullTime convert string to NullTime
func NewNullTime(value string) sql.NullTime {

	const layoutISO = "2006-01-02"

	if isEmpty(value) {
		return sql.NullTime{Valid: false}
	}

	var result time.Time
	var err error

	if result, err = time.Parse(layoutISO, value); err != nil {
		log.Fatal(err)
		return sql.NullTime{}
	}

	return sql.NullTime{
		Time:  result,
		Valid: true,
	}
}
