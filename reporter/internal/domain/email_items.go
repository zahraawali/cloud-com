package domain

// EmailItem Struct to represent the email item
type EmailItem struct {
	ID           int     `json:"id"`
	Active       bool    `json:"active"`
	InsertTime   string  `json:"inserttime"`
	UpdateTime   string  `json:"updatetime"`
	Version      int     `json:"version"`
	Body         int     `json:"body"`
	ReplyAddress string  `json:"reply_address"`
	Subject      string  `json:"subject"`
	Text         *string `json:"text"`
	MailConfig   int     `json:"mail_config"`
	RequestID    int     `json:"request_id"`
}
