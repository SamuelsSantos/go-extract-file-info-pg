package files

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

// ProcessExtract process to extract data from file
func ProcessExtract(filename string, layout *Layout) (*Rows, error) {

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	if layout.IgnoreHeader {
		scanner.Scan()
	}

	var rows Rows
	for scanner.Scan() {
		values := strings.Fields(scanner.Text()) // Fields splits the string s around each instance of one or more consecutive white space characters
		if len(values) != len(layout.Fields) {
			return nil, errors.New("Invalid file layout")
		}

		rows = append(rows, ExtractedDataModel{
			CustomerID:          values[layout.Fields["customer_id"]],
			Private:             values[layout.Fields["private"]],
			Incomplete:          values[layout.Fields["incomplete"]],
			LastDateShop:        values[layout.Fields["last_shop"]],
			AvgTicket:           values[layout.Fields["avg_ticket"]],
			LastTicketShop:      values[layout.Fields["last_ticket_shop"]],
			MostFrequentedStore: values[layout.Fields["most_frequented_store"]],
			LastStore:           values[layout.Fields["last_store"]],
		})
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return &rows, nil
}
