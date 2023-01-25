package shared

import (
	"time"
)

type SendEmailResponse struct {
	CreatedAt time.Time `json:"created_at"`
	From      string    `json:"from"`
	ID        string    `json:"id"`
	To        string    `json:"to"`
}
