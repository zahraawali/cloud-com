package handler

import (
	"database/sql"
	"net/http"
	"reporter/internal/domain"

	"github.com/gin-gonic/gin"
)

func GetAllEmailRequests(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := db.Query("SELECT id, subject, body, COALESCE(text, '0'), mail_config, reply_address FROM not2_mail_request")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch email requests"})
			return
		}
		defer rows.Close()

		emailRequests := make([]domain.EmailRequest, 0)
		for rows.Next() {
			var request domain.EmailRequest
			if err := rows.Scan(&request.ID, &request.Subject, &request.Body, &request.Text, &request.MailConfig, &request.ReplyAddress); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch email requests", "err": err.Error()})
				return
			}
			emailRequests = append(emailRequests, request)
		}

		c.JSON(http.StatusOK, emailRequests)
	}
}
