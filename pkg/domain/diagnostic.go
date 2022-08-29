package domain

type Diagnostic struct {
	ID            int64
	DefinedNumber string
	SKU           string
	Images        []Image
}
