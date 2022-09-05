package domain

import "time"

type Diagnostic struct {
	ID        int64
	CreatedAt time.Time
	UpdatedAt time.Time

	DefinedNumber string
	SKU           string

	Items []Item

	Images []Image
}
