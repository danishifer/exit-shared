package actionsService

import exit "exit-shared-mongo"

type CreateActionRequest struct {
	UserID   exit.UserID `json:"user_id"`
	Type     string      `json:"type"`
	DeviceID string      `json:"device_id"`
}

type CreateActionResponse struct {
	Status exit.ActionStatus `json:"status"`
	Action exit.Action       `json:"action"`
}

func (r *CreateActionRequest) IsValid() bool {
	if r.UserID == "" || r.Type == "" || r.DeviceID == "" {
		return false
	}
	return true
}
