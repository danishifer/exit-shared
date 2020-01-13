package exit_shared

import "fmt"

type FrameworkHTTPError struct {
	Code    int
	Message string
}

func (e *FrameworkHTTPError) Error() string {
	return fmt.Sprintf("framework error [%d]: \"%s\"", e.Code, e.Message)
}
