package handler

import (
	"database/sql"
	"net/http"
	"reporter/internal/domain"

	"github.com/gin-gonic/gin"
)

func GetEmailRequestState(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.Param("id")
		// requestID, err := strconv.Atoi(id)
		// if err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email request ID"})
		// 	return
		// }

		// SELECT * FROM not2_mail_item item
		// JOIN not2_mail_request n2mr on n2mr.id = item.mail_request
		// JOIN not2_request n2r on n2mr.request_id = n2r.id
		// where n2r.notification_id = '380a44b7-10e7-4fca-9a7a-d5d5508698e1'
		// and n2r.not_user = 2;

		var state domain.QueryResult
		err := db.QueryRow(`SELECT * FROM not2_mail_item item
						   JOIN not2_mail_request n2mr on n2mr.id = item.mail_request
						   JOIN not2_request n2r on n2mr.request_id = n2r.id
						   where n2r.notification_id = $1
						   and n2r.not_user = 0`, requestID).Scan(
			&state.Rows.ID,
			&state.Rows.Active,
			&state.Rows.InsertTime,
			&state.Rows.UpdateTime,
			&state.Rows.Version,
			&state.Rows.MessageID,
			&state.Rows.MetaData,
			&state.Rows.Receiver,
			&state.Rows.SendDate,
			&state.Rows.State,
			&state.Rows.MailRequest,
			&state.Rows.EmailItem.ID,
			&state.Rows.EmailItem.Active,
			&state.Rows.EmailItem.InsertTime,
			&state.Rows.EmailItem.UpdateTime,
			&state.Rows.EmailItem.Version,
			&state.Rows.EmailItem.Body,
			&state.Rows.EmailItem.ReplyAddress,
			&state.Rows.EmailItem.Subject,
			&state.Rows.EmailItem.Text,
			&state.Rows.EmailItem.MailConfig,
			&state.Rows.EmailItem.RequestID,
			&state.Rows.EmailRequest.ID,
			&state.Rows.EmailRequest.Active,
			&state.Rows.EmailRequest.InsertTime,
			&state.Rows.EmailRequest.UpdateTime,
			&state.Rows.EmailRequest.Version,
			&state.Rows.EmailRequest.JobID,
			&state.Rows.EmailRequest.MessageType,
			&state.Rows.EmailRequest.NotificationID,
			&state.Rows.EmailRequest.NotUser,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch email request state", "err": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"id": requestID, "state": state})
	}
}
