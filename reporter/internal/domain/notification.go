package domain

import "time"

type Notification struct {
	ID           int          `json:"id"`
	Active       bool         `json:"active"`
	InsertTime   time.Time    `json:"inserttime"`
	UpdateTime   time.Time    `json:"updatetime"`
	Version      int          `json:"version"`
	MessageID    string       `json:"message_id"`
	MetaData     string       `json:"meta_data"`
	Receiver     string       `json:"receiver"`
	SendDate     time.Time    `json:"send_date"`
	State        int          `json:"state"`
	MailRequest  int          `json:"mail_request"`
	EmailItem    EmailItem    `json:"-"`
	EmailRequest EmailRequest `json:"-"`
}
