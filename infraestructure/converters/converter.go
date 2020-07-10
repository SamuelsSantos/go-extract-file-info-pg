package converters

import (
	"database/sql"
	"log"
	"strconv"
	"strings"
	"time"
)

// Converter struct
type Converter struct{}

// TransformingProcess interface converter string to sql.DataType
type TransformingProcess interface {
	NewNullString() *sql.NullString
	NewNullInt() *sql.NullInt64
	NewNullFloat() *sql.NullFloat64
	NewNullTime() *sql.NullTime
}

func isEmpty(value string) bool {
	return len(value) == 0 || value == "NULL"
}

func replaceDecimalSeparator(value string) string {
	return strings.Replace(value, ",", ".", -1)
}

// NewNullString convert string to NullString
func (c *Converter) NewNullString(value string) *sql.NullString {
	if isEmpty(value) {
		return &sql.NullString{Valid: false}
	}
	return &sql.NullString{
		String: value,
		Valid:  true,
	}
}

// NewNullInt convert string to NullInt64
func (c *Converter) NewNullInt(value string) *sql.NullInt64 {
	if isEmpty(value) {
		return &sql.NullInt64{}
	}

	var result int64
	var err error

	if result, err = strconv.ParseInt(value, 10, 64); err != nil {
		log.Fatal(err)
		return &sql.NullInt64{}
	}

	return &sql.NullInt64{
		Int64: result,
		Valid: true,
	}
}

// NewNullFloat convert string to NullFloat64
func (c *Converter) NewNullFloat(value string) *sql.NullFloat64 {
	if isEmpty(value) {
		return &sql.NullFloat64{}
	}

	var result float64
	var err error

	value = replaceDecimalSeparator(value)
	if result, err = strconv.ParseFloat(value, 64); err != nil {
		log.Fatal(err)
		return &sql.NullFloat64{}
	}

	return &sql.NullFloat64{
		Float64: result,
		Valid:   true,
	}
}

// NewNullTime parses a formatted string and returns the time value it represents.
// The layout defines the format by showing how the reference time, defined to be
// Mon Jan 2 15:04:05 -0700 MST 2006
func (c *Converter) NewNullTime(layout, value string) *sql.NullTime {

	if isEmpty(value) {
		return &sql.NullTime{Valid: false}
	}

	var result time.Time
	var err error

	if result, err = time.Parse(layout, value); err != nil {
		log.Fatal(err)
		return &sql.NullTime{}
	}

	return &sql.NullTime{
		Time:  result,
		Valid: true,
	}
}
