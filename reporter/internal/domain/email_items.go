package domain

// EmailItem Struct to represent the email item
type EmailItem struct {
	ID          int    `json:"id"`
	Receiver    string `json:"receiver"`
	SendDate    string `json:"send_date"`
	MetaData    string `json:"meta_data"`
	MessageID   string `json:"message_id"`
	State       int    `json:"state"`
	InsertTime  string `json:"insert_time"`
	UpdateTime  string `json:"update_time"`
	Active      bool   `json:"active"`
	MailRequest int    `json:"mail_request"`
}
