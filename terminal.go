package exit_shared_mongo

type TerminalStatus int32

type Terminal struct {
	ID           string         `bson:"id" json:"id"`
	Status       TerminalStatus `bson:"status" json:"status"`
	Name         string         `bson:"name" json:"name"`
	FriendlyName string         `bson:"friendlyName" json:"friendly_name"`
}
