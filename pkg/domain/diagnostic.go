package domain

import "time"

type Diagnostic struct {
	ID      uint64
	Version uint32

	CreatedAt time.Time
	UpdatedAt time.Time

	DefinedNumber string
	SKU           string

	Items  []Item
	Images []Image
}
