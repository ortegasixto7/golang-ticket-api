package models

type Integration struct {
	ID          string
	Name        string
	Description string
	PublicKey   string
	IsEnabled   bool
	CreatedAt   int64
	UpdatedAt   int64
	DeletedAt   int64
}
