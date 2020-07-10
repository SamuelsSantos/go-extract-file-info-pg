package service

import (
	"fmt"
	"strings"

	"github.com/desafios-job/import-data/domain/entity"
	"github.com/desafios-job/import-data/infraestructure/converters"
	"github.com/desafios-job/import-data/infraestructure/files"
	"github.com/desafios-job/import-data/infraestructure/persistence"
	"github.com/klassmann/cpfcnpj"
)

func neowayLayout() *files.Layout {

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

type neowayDW struct {
	FileName             string
	Layout               files.Layout
	InconsistencyService InconsistencyAppInterface
	ShoppingService      ShoppingAppInterface
}

func cleanMaskDocument(value string) string {
	return cpfcnpj.Clean(value)
}

// Clean data, Truncate tables shopping and inconsistency
func (dw *neowayDW) Clean() error {
	er := dw.InconsistencyService.Truncate()
	if er != nil {
		return er
	}

	err := dw.ShoppingService.Truncate()
	if err != nil {
		return err
	}

	return nil
}

// Extract data from file
func (dw *neowayDW) Extract() (*files.Rows, error) {
	return files.ProcessExtract(dw.FileName, &dw.Layout)
}

// Transform file data to business data
func (dw *neowayDW) Transform(rows *files.Rows) (*entity.Shoppings, *entity.Inconsistencies) {
	converter := converters.Converter{}
	shoppings := make(entity.Shoppings, 0)
	inconsistencies := make(entity.Inconsistencies, 0)

	for _, row := range *rows {
		shopping := &entity.Shopping{
			CustomerID:          *converter.NewNullString(cleanMaskDocument(row.CustomerID)),
			Private:             *converter.NewNullInt(row.Private),
			Incomplete:          *converter.NewNullInt(row.Incomplete),
			LastDateShop:        *converter.NewNullTime("2006-01-02", row.LastDateShop),
			AvgTicket:           *converter.NewNullFloat(row.AvgTicket),
			LastTicketShop:      *converter.NewNullFloat(row.AvgTicket),
			MostFrequentedStore: *converter.NewNullString(cleanMaskDocument(row.MostFrequentedStore)),
			LastStore:           *converter.NewNullString(cleanMaskDocument(row.LastStore)),
		}

		if _, err := shopping.Validate(); len(err) != 0 {

			inconsistency := entity.NewInconsistency(
				*converter.NewNullString(dw.FileName),
				*converter.NewNullString(fmt.Sprintf("%s", strings.Join(err, ","))),
			)

			inconsistencies = append(inconsistencies, inconsistency)
		} else {
			shoppings = append(shoppings, shopping)
		}

	}

	return &shoppings, &inconsistencies
}

// SaveShoppings save handled shoppings
func (dw *neowayDW) SaveShoppings(shoppings *entity.Shoppings) {
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
			dw.ShoppingService.SaveMany((*shoppings)[start:end])
			start = end
			end += Denominador
			handleData <- "done"
		}
	}

	dw.ShoppingService.SaveMany((*shoppings)[size-mod:])

	handleData := make(chan string)
	go loopData(handleData)

	for num := range handleData {
		fmt.Println(num)
	}
}

// SaveInconsistencies save handled inconsistencies
func (dw *neowayDW) SaveInconsistencies(inconsistencies *entity.Inconsistencies) {
	dw.InconsistencyService.SaveMany(*inconsistencies)
}

// NewNeowayDW instantiating uninitialized objects
func NewNeowayDW(filename string, repositories persistence.Repositories) DW {
	return &neowayDW{
		FileName:             filename,
		Layout:               *neowayLayout(),
		InconsistencyService: NewInconsistencyService(repositories.Inconsistency),
		ShoppingService:      NewShoppingService(repositories.Shopping),
	}
}
