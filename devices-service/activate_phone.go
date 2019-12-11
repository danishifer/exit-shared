package devicesService

import exit "exit-shared-mongo"

type ActivatePhoneRequest struct {
	ActivationCode string `json:"activation_code"`
	DeviceName     string `json:"device_name"`
	DeviceInfo     string `json:"device_info"`
}

type ActivatePhoneResponse struct {
	Error  string             `json:"error,omitempty"`
	Status *exit.DeviceStatus `json:"status,omitempty"`
	Device *exit.Device       `json:"device,omitempty"`
}

func (r *ActivatePhoneRequest) IsValid() bool {
	if r.ActivationCode == "" || r.DeviceName == "" || r.DeviceInfo == "" {
		return false
	}
	return true
}
