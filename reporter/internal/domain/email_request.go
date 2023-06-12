package domain

// EmailRequest Struct to represent the email request
type EmailRequest struct {
	ID             int    `json:"id"`
	Active         bool   `json:"active"`
	InsertTime     string `json:"inserttime"`
	UpdateTime     string `json:"updatetime"`
	Version        int    `json:"version"`
	JobID          *int   `json:"job_id"`
	MessageType    int    `json:"message_type"`
	NotificationID string `json:"notification_id"`
	NotUser        int    `json:"not_user"`
}
