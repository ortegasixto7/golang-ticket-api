package models

type Ticket struct {
	ID        string
	Title     string
	Code      string
	CreatedAt int64
	UpdatedAt int64
	ExpiredAt int64
}
