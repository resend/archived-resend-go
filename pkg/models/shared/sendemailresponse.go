// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

// SendEmailResponse - OK
type SendEmailResponse struct {
	// The date and time the email was sent.
	CreatedAt time.Time `json:"created_at"`
	// The sender email address.
	From string `json:"from"`
	// The unique identifier of the sent email.
	ID string `json:"id"`
	// The recipient email address.
	To string `json:"to"`
}
