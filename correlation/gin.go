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

	// Store the trace ID in the request context
	ctx := context.WithValue(c.Request.Context(), "traceID", traceID)
	c.Request = c.Request.WithContext(ctx)

	// Create a new logger with the traceID ingested as correlationID
	ctx1 := logger.New(c.Request.Context(), traceID)
	c.Request = c.Request.WithContext(ctx1)

	// Continue processing the request
	c.Next()
}
