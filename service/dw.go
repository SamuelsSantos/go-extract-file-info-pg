package service

import (
	"fmt"
	"strings"

	"github.com/desafios-job/import-data/domain/entity"
	"github.com/desafios-job/import-data/infraestructure/converters"
	"github.com/desafios-job/import-data/infraestructure/files"
)

// NeowayLayout get neoway layout file
func NeowayLayout() *files.Layout {

	return &files.Layout{
		IgnoreHeader: true,
		Fields: map[string]int{
			"customer_id":           0,
			"private":               1,
			"incomplete":            2,
			"last_shop":             3,
			"avg_ticket":            4,
			"last_ticket_shop":      5,
			"most_frequented_store": 6,
			"last_store":            7,
		},
	}
}

// DWNeoway interface
type DWNeoway struct {
	FileName         string
	LayoutFile       files.Layout
	InconsistencyApp InconsistencyAppInterface
	ShoppingApp      ShoppingAppInterface
}

// DW Data Warehouse interface
type DW interface {
	Clean() error
	Extract() (*files.Rows, error)
	Transform() (*entity.Shopping, *entity.Inconsistencies)
	SaveShoppings(shoppings *entity.Shoppings)
	SaveInconsistencies(inconsistencies *entity.Inconsistencies)
}

// Extract data from file
func (dw *DWNeoway) Extract() (*files.Rows, error) {
	return files.ProcessExtract(dw.FileName, &dw.LayoutFile)
}

// Transform file data to business data
func (dw *DWNeoway) Transform(rows files.Rows) (*entity.Shoppings, *entity.Inconsistencies) {

	shoppings := make(entity.Shoppings, 0)
	inconsistencies := make(entity.Inconsistencies, 0)

	for _, row := range rows {
		shopping := &entity.Shopping{
			CustomerID:          converters.NewNullString(converters.CleanMaskDocument(row.CustomerID)),
			Private:             converters.NewNullInt(row.Private),
			Incomplete:          converters.NewNullInt(row.Incomplete),
			LastDateShop:        converters.NewNullTime(row.LastDateShop),
			AvgTicket:           converters.NewNullFloat(row.AvgTicket),
			LastTicketShop:      converters.NewNullFloat(row.AvgTicket),
			MostFrequentedStore: converters.NewNullString(converters.CleanMaskDocument(row.MostFrequentedStore)),
			LastStore:           converters.NewNullString(converters.CleanMaskDocument(row.LastStore)),
		}

		if _, err := shopping.Validate(); len(err) != 0 {

			inconsistency := &entity.Inconsistency{
				FileName:     converters.NewNullString(dw.FileName),
				ErrorMessage: converters.NewNullString(fmt.Sprintf("%s", strings.Join(err, ","))),
			}

			inconsistencies = append(inconsistencies, inconsistency)
		} else {
			shoppings = append(shoppings, shopping)
		}

	}

	return &shoppings, &inconsistencies
}

// SaveShoppings save handled shoppings
func (dw *DWNeoway) SaveShoppings(shoppings *entity.Shoppings) {

	const Denominador = 65000 / 8
	size := len(*shoppings)

	times := size / Denominador
	mod := size % Denominador

	start := 0
	end := Denominador

	data := make([]string, times)

	loopData := func(handleData chan<- string) {
		defer close(handleData)
		for range data {
			fmt.Printf("Start: %v End: %v \n", start, end)
			dw.ShoppingApp.SaveMany((*shoppings)[start:end])
			start = end
			end += Denominador
			handleData <- "done"
		}
	}

	dw.ShoppingApp.SaveMany((*shoppings)[size-mod:])

	handleData := make(chan string)
	go loopData(handleData)

	for num := range handleData {
		fmt.Println(num)
	}
}

// SaveInconsistencies save handled inconsistencies
func (dw *DWNeoway) SaveInconsistencies(inconsistencies *entity.Inconsistencies) {
	dw.InconsistencyApp.SaveMany(*inconsistencies)
}

// Clean data, Truncate tables shopping and inconsistency
func (dw *DWNeoway) Clean() error {
	er := dw.InconsistencyApp.Truncate()
	if er != nil {
		return er
	}

	err := dw.ShoppingApp.Truncate()
	if err != nil {
		return err
	}

	return nil
}
