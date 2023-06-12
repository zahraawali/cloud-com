package handler

import (
	"database/sql"
	"net/http"
	"reporter/internal/domain"

	"github.com/gin-gonic/gin"
)

func GetEmailRequestStateByConfig(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		configIDInt := c.Param("configID")
		// configIDInt, err := strconv.Atoi(configID)
		// if err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid mail config ID"})
		// 	return
		// }

		rows, err := db.Query(`SELECT * FROM not2_mail_item item  
							   JOIN not2_mail_request n2mr on item.mail_request = n2mr.id  
							   JOIN not2_request n2r on n2mr.request_id = n2r.id  
							   WHERE n2mr.mail_config = $1
							   and n2r.not_user = 0;`, configIDInt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to fetch email request states by config",
				"err":   err,
			})
			return
		}
		defer rows.Close()

		var emailItems []domain.Notification
		for rows.Next() {
			var item domain.Notification

			err = rows.Scan(
				&item.ID,
				&item.Active,
				&item.InsertTime,
				&item.UpdateTime,
				&item.Version,
				&item.MessageID,
				&item.MetaData,
				&item.Receiver,
				&item.SendDate,
				&item.State,
				&item.MailRequest,
				&item.EmailItem.ID,
				&item.EmailItem.Active,
				&item.EmailItem.InsertTime,
				&item.EmailItem.UpdateTime,
				&item.EmailItem.Version,
				&item.EmailItem.Body,
				&item.EmailItem.ReplyAddress,
				&item.EmailItem.Subject,
				&item.EmailItem.Text,
				&item.EmailItem.MailConfig,
				&item.EmailItem.RequestID,
				&item.EmailRequest.ID,
				&item.EmailRequest.Active,
				&item.EmailRequest.InsertTime,
				&item.EmailRequest.UpdateTime,
				&item.EmailRequest.Version,
				&item.EmailRequest.JobID,
				&item.EmailRequest.MessageType,
				&item.EmailRequest.NotificationID,
				&item.EmailRequest.NotUser,
			)

			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Failed to parse email request states",
					"err":   err.Error(),
				})
				return
			}

			emailItems = append(emailItems, item)
		}

		c.JSON(http.StatusOK, emailItems)
	}
}
