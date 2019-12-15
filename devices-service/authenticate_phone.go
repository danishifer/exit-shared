package devicesService

import exit "exit-shared"

type AuthenticatePhoneRequest struct {
	UserID exit.UserID `json:"user_id" header:"x-exit-user-id" binding:"required"`
	APIKey string      `json:"api_key" header:"x-exit-api-key" validate:"len=16" binding:"required"`
}

type AuthenticatePhoneResponse struct {
	DeviceID      string `json:"device_id"`
	Authenticated bool   `json:"authenticated"`
}

func (r *AuthenticatePhoneRequest) IsValid() bool {
	if r.UserID == "" || len(r.APIKey) != 16 {
		return false
	}
	return true
}
