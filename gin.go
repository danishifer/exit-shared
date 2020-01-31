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

// ERRORS
type ErrorBadRequest struct{}

func (e *ErrorBadRequest) Error() string {
	return fmt.Sprint("bad request")
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
		return &ErrorBadRequest{}
	}

	return nil
}
