package app

import (
	"database/sql"
	"fmt"
	"log"
	"reporter/internal/handler"
	"reporter/internal/middleware"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func Starter(dsn string, addr string, authServiceBase string) error {
	fmt.Println("application starter called")

	// Connect to the database
	var err error
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Println("Failed to connect to the database:", err)
		return err
	}
	defer db.Close()

	// Initialize the Gin router
	router := gin.Default()

	// middleware
	router.Use(middleware.Auth(authServiceBase))

	// Define the API endpoints
	router.GET("/email-requests", handler.GetAllEmailRequests(db))
	router.GET("/email-requests/:id/state", handler.GetEmailRequestState(db))
	router.GET("/email-requests/config/:configID/state", handler.GetEmailRequestStateByConfig(db))

	// Start the server
	log.Printf("Server is running on %s", addr)
	return router.Run(addr)
}
