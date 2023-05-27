package handler

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"reporter/internal/domain"
	"strconv"
)

func GetEmailRequestStateByConfig(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		configID := c.Param("configID")
		configIDInt, err := strconv.Atoi(configID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid mail config ID"})
			return
		}

		rows, err := db.Query("SELECT m.id, m.receiver, m.send_date, m.meta_data, m.message_id, m.state, m.inserttime, m.updatetime, m.active, m.mail_request "+
			"FROM not2_mail_item m JOIN not2_mail_request r ON m.mail_request = r.id WHERE r.mail_config = $1", configIDInt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch email request states by config"})
			return
		}
		defer rows.Close()

		emailItems := make([]domain.EmailItem, 0)
		for rows.Next() {
			var item domain.EmailItem
			if err := rows.Scan(&item.ID, &item.Receiver, &item.SendDate, &item.MetaData, &item.MessageID, &item.State, &item.InsertTime, &item.UpdateTime, &item.Active, &item.MailRequest); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch email request states by config"})
				return
			}
			emailItems = append(emailItems, item)
		}

		c.JSON(http.StatusOK, emailItems)
	}
}
