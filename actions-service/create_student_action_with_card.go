package actionsService

import exit "exit-shared"

type CreateStudentActionWithCardRequest struct {
	UserID     exit.UserID `json:"user_id"`
	Type       string      `json:"type"`
	DeviceID   string      `json:"device_id"`
	CardSerial string      `json:"card_serial"`
}

type CreateStudentActionWithCardResponse struct {
	Status exit.ActionStatus `json:"status"`
	Action exit.Action       `json:"action"`
}

func (r *CreateStudentActionWithCardRequest) IsValid() bool {
	if r.UserID == "" || r.Type == "" || r.DeviceID == "" {
		return false
	}
	return true
}
