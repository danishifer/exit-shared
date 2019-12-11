package gatewayService

import devicesService "exit-shared-mongo/devices-service"

type StudentPhoneActivateRequest struct {
	FamilyName     string `json:"family_name"`
	ActivationCode string `json:"activation_code"`
	DeviceName     string `json:"device_name"`
	DeviceInfo     string `json:"device_info"`
}

type StudentPhoneActivateResponse devicesService.ActivatePhoneResponse

func (r *StudentPhoneActivateRequest) IsValid() bool {
	if r.FamilyName == "" || len(r.ActivationCode) != 6 {
		return false
	}
	if r.DeviceName == "" || r.DeviceInfo == "" {
		return false
	}
	return true
}
