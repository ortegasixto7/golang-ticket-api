package generate

import "time"

type Response struct {
	ID        string     `json:"id"`
	Title     string     `json:"title"`
	Code      string     `json:"code"`
	ExpiresAt *time.Time `json:"expires_at"`
}
