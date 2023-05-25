package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Restore the ID of the existing task.

func ErrorHandlingMiddleware(c *gin.Context) {
	// Defer a function that recovers from a panic and sends an error response.
	defer func() {
		if err := recover(); err != nil {
			// Log the error.
			log.Println("Recovered from panic in handler:", err)

			// Send an error response.
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
	}()

	// Call the next middleware/handler function.
	c.Next()
}

func LoggingMiddleware(c *gin.Context) {
	log.Printf("Before request processing: %s %s\n", c.Request.Method, c.Request.URL.Path)
	c.Next() // Pass to the next middleware or final handler
	log.Printf("After request processed: %s\n", c.Writer.Status())
}
