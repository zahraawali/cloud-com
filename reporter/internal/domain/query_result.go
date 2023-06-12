package domain

type QueryResult struct {
	SQLQuery string       `json:"-"`
	Rows     Notification `json:""`
}
