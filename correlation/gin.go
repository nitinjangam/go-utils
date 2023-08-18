package correlation

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/nitinjangam/go-utils/logger"
)

// TraceMiddleware is a middleware to fetch or generate trace ID
func TraceMiddleware(c *gin.Context) {
	// Check if trace ID exists in the request headers
	traceID := c.GetHeader("X-Trace-ID")

	// If trace ID doesn't exist, generate a new one
	if traceID == "" {
		traceID = uuid.New().String()
	}

	// Create a new logger with the traceID ingested as correlationID
	ctx := logger.New(c.Request.Context(), traceID)

	// Store the trace ID in the request context
	ctx = context.WithValue(c.Request.Context(), "traceID", traceID)

	// Add context with logger instance back to request context
	c.Request = c.Request.WithContext(ctx)

	// Continue processing the request
	c.Next()
}
