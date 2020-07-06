package files

// Layout file
type Layout struct {
	IgnoreHeader bool
	Fields       map[string]int
}

// ExtractedDataModel is a model data from file
type ExtractedDataModel struct {
	CustomerID          string
	Private             string
	Incomplete          string
	LastDateShop        string
	AvgTicket           string
	LastTicketShop      string
	MostFrequentedStore string
	LastStore           string
}

// Rows is a collection of ExtractedDataModel
type Rows []ExtractedDataModel
