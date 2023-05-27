package domain

// EmailRequest Struct to represent the email request
type EmailRequest struct {
	ID           int    `json:"id"`
	Subject      string `json:"subject"`
	Body         string `json:"body"`
	Text         string `json:"text"`
	MailConfig   int    `json:"mail_config"`
	ReplyAddress string `json:"reply_address"`
	// TODO: Add other fields as needed
}
