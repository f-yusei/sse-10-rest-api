package model

// Store represents a store entity in the business layer.
type Store struct {
	ID             int
	Name           string
	DisplayMessage string
	Bells          []Bell // Associated bells for the store
}
