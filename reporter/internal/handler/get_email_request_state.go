package handler

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetEmailRequestState(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		requestID, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email request ID"})
			return
		}

		var state int
		err = db.QueryRow("SELECT state FROM not2_mail_item WHERE mail_request = $1", requestID).Scan(&state)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch email request state", "err": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"id": requestID, "state": state})
	}
}
