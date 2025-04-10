package models

type Interactor struct {
	ID         string
	Title      string
	PublicKey  string
	PrivateKey string
	IsEnabled  bool
	CreatedAt  int64
}
