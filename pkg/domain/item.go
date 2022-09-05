package domain

type ItemType int

const (
	Product ItemType = iota + 1
	Service
)

type Item struct {
	ItemType
	ID   uint64
	Name string
}
