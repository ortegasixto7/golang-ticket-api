package generate

import "time"

type Request struct {
	Title     string     `json:"title"`
	ExpiresAt *time.Time `json:"expires_at"`
}
