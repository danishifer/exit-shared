package exit_shared

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Action struct {
	ID         string       `bson:"id" json:"id"`
	UserID     UserID       `bson:"userId" json:"user_id"`
	Status     ActionStatus `bson:"status" json:"status"`
	Type       string       `bson:"type" json:"type"`
	Date       *time.Time   `bson:"date" json:"date"`
	DeviceID   string       `bson:"deviceId" json:"device_id"`
	TerminalID string       `bson:"terminalId" json:"terminal_id"`
	PermitID   *string      `bson:"permitId,omitempty" json:"permit_id,omitempty"`
}

type ActionIntent struct {
	UserID UserID    `json:"user_id"`
	Type   string    `json:"type"`
	Time   time.Time `json:"time"`
}

const (
	ActionStatusPending  = 0
	ActionStatusApproved = 1
	ActionStatusRejected = 2
)

func (action *Action) ToActionIntent() ActionIntent {
	return ActionIntent{
		UserID: action.UserID,
		Type:   action.Type,
		Time:   *action.Date,
	}
}

func (action *Action) Approve(collection *mongo.Collection, permitID string) (*mongo.UpdateResult, error) {
	filter := bson.D{{"id", action.ID}}
	update := bson.D{
		{"$set", bson.D{
			{"permitId", permitID},
			{"status", ActionStatusApproved},
		}},
	}
	return collection.UpdateOne(context.TODO(), filter, update)
}

func (action *Action) Reject(collection *mongo.Collection) (*mongo.UpdateResult, error) {
	filter := bson.D{{"id", action.ID}}
	update := bson.D{
		{"$set", bson.D{
			{"status", ActionStatusRejected},
		}},
	}
	return collection.UpdateOne(context.TODO(), filter, update)
}
