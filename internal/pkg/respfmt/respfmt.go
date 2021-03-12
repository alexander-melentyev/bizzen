package respfmt

import (
	"os"

	"github.com/gin-gonic/gin"
)

// Fmt - common response format.
type Fmt struct {
	TotalCount  uint64      `json:"totalCount,omitempty"`
	ServiceName string      `json:"serviceName,omitempty"`
	UserMsg     string      `json:"userMessage,omitempty"`
	DevMsg      string      `json:"developerMessage,omitempty"`
	Data        interface{} `json:"data,omitempty"`
}

// Err -.
func Err(c *gin.Context, statusCode int, userMsg string, devMsg error) {
	c.JSON(statusCode, Fmt{
		ServiceName: os.Getenv("SERVICE_NAME"),
		UserMsg:     userMsg,
		DevMsg:      devMsg.Error(),
	})
}
