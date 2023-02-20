package shared

type Email struct {
	Bcc     *string `json:"bcc,omitempty"`
	Cc      *string `json:"cc,omitempty"`
	From    string  `json:"from"`
	HTML    *string `json:"html,omitempty"`
	ReplyTo *string `json:"reply_to,omitempty"`
	Subject string  `json:"subject"`
	Text    *string `json:"text,omitempty"`
	To      string  `json:"to"`
}
