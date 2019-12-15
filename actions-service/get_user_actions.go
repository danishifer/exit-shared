package actionsService

import (
	exit "exit-shared"
	"time"
)

type GetStudentActionsRequest struct {
	UserID   exit.UserID `json:"user_id"`
	FromDate *time.Time  `json:"from_date"`
	ToDate   *time.Time  `json:"to_date"`
}

type GetStudentActionsResponse struct {
	Actions []exit.Action `json:"actions"`
}

func (r *GetStudentActionsRequest) IsValid() bool {
	if r.UserID == "" || r.FromDate == nil || r.ToDate == nil {
		return false
	}
	return true
}
