package actionsService

import exit "exit-shared"

type CreateStudentActionWithRequest struct {
	UserID     exit.UserID `json:"user_id"`
	Type       string      `json:"type"`
	DeviceID   string      `json:"device_id"`
	TerminalID string      `json:"terminal_id"`
}

type CreateActionWithCardResponse struct {
	Status exit.ActionStatus `json:"status"`
	Action exit.Action       `json:"action"`
}

func (r *CreateStudentActionWithRequest) IsValid() bool {
	if r.UserID == "" || r.Type == "" || r.DeviceID == "" {
		return false
	}
	return true
}
