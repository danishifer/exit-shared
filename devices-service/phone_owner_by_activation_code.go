package devicesService

import exit "exit-shared-mongo"

type PhoneOwnerByActivationCodeRequest struct {
	ActivationCode string `json:"device_id"`
}

type PhoneOwnerByActivationCodeResponse struct {
	UserID exit.UserID `json:"user_id"`
}

func (r *PhoneOwnerByActivationCodeRequest) IsValid() bool {
	if r.ActivationCode == "" {
		return false
	}
	return true
}
