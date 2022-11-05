package domain

import "time"

type Diagnostic struct {
	ID      int
	Version int

	CreatedAt time.Time
	UpdatedAt time.Time

	DefinedNumber string
	SKU           string

	Items  []Item
	Images []Image
}
