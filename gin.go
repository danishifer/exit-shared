package exit_shared

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Request interface {
	IsValid() bool
}

type Response interface{}

// Errors
type FrameworkHTTPError struct {
	Code    int
	Message string
}

func (e *FrameworkHTTPError) Error() string {
	return fmt.Sprintf("framework error [%d]: \"%s\"", e.Code, e.Message)
}

func DecodeRequest(request Request, c *gin.Context) error {
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "cannot parse body",
		})
		return err
	}

	if !request.IsValid() {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "bad request",
		})
		return &FrameworkHTTPError{
			Code:    http.StatusBadRequest,
			Message: "bad request",
		}
	}

	return nil
}
