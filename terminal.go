package exit_shared

import "time"

type TerminalStatus int32

type Terminal struct {
	ID           string          `bson:"id" json:"id"`
	Status       TerminalStatus  `bson:"status" json:"status"`
	Type         string          `bson:"type" json:"type"`
	Name         string          `bson:"name" json:"name"`
	FriendlyName string          `bson:"friendlyName" json:"friendly_name"`
	LastActive   *time.Time      `bson:"lastActive" json:"last_active"`
	Soc          *SocTerminal    `bson:"soc,omitempty" json:"soc,omitempty"`
	Card         *CardTerminal   `bson:"card,omitempty" json:"card,omitempty"`
	Actions      map[string]bool `bson:"actions" json:"actions"`
}

type CardTerminal struct {
	Serial string `bson:"serial" json:"serial"`
	Number string `bson:"number" json:"number"`
}

type SocTerminal struct {
	Version string `bson:"version" json:"version"`
}
