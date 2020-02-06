package exit_shared

import "time"

type Device struct {
	ID             string       `bson:"id" json:"id"`
	Type           string       `bson:"type" json:"type"`
	Status         DeviceStatus `bson:"status" json:"status"`
	UserID         UserID       `bson:"userId,omitempty" json:"user_id"`
	CardSerial     string       `bson:"cardSerial,omitempty" json:"card_serial,omitempty"`
	CardNumber     string       `bson:"cardNumber,omitempty" json:"card_number,omitempty"`
	DeviceName     string       `bson:"deviceName,omitempty" json:"device_name,omitempty"`
	DeviceInfo     string       `bson:"deviceInfo,omitempty" json:"device_info,omitempty"`
	ActivationCode string       `bson:"activationCode,omitempty" json:"activation_code,omitempty"`
	APIKey         string       `bson:"apiKey,omitempty" json:"api_key,omitempty"`
	ActivatedOn    *time.Time   `bson:"activatedOn,omitempty" json:"activated_on,omitempty"`
	DeactivatedOn  *time.Time   `bson:"deactivatedOn,omitempty" json:"deactivated_on,omitempty"`
}

const (
	DeviceStatusPendingActivation = 0
	DeviceStatusActivated         = 1
	DeviceStatusDeactivated       = 2
)

type DeviceMetadata struct {
	ID          string       `json:"id"`
	Type        string       `json:"type"`
	Status      DeviceStatus `json:"status"`
	UserID      UserID       `json:"user_id"`
	CardSerial  string       `json:"card_serial,omitempty"`
	CardNumber  string       `json:"card_number,omitempty"`
	DeviceName  string       `json:"device_name,omitempty"`
	DeviceInfo  string       `json:"device_info,omitempty"`
	ActivatedOn *time.Time   `json:"activated_on,omitempty"`
}
