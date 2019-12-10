package exit_shared_mongo

import "time"

type Device struct {
	ID             string         `bson:"id"`
	Type           string         `bson:"type"`
	Status         DeviceStatus   `bson:"status"`
	UserID         UserID         `bson:"userId,omitempty"`
	CardSerial     string         `bson:"cardSerial,omitempty"`
	CardNumber     string         `bson:"cardNumber,omitempty"`
	Metadata       DeviceMetadata `bson:"metadata,omitempty"`
	ActivationCode string         `bson:"activationCode,omitempty"`
	APIKey         string         `bson:"apiKey,omitempty"`
	ActivatedOn    time.Time      `bson:"activatedOn,omitempty"`
}

type DeviceMetadata struct {
	DeviceName string `bson:"deviceName"`
	DeviceInfo string `bson:"deviceInfo"`
}

const (
	DeviceStatusUnlinked = 0
	DeviceStatusLinked   = 1
)
