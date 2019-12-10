package exit_shared_mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Action struct {
	ID       string       `bson:"id"`
	UserID   UserID       `bson:"userId"`
	Status   ActionStatus `bson:"status"`
	Type     string       `bson:"type"`
	DeviceID string       `bson:"device"`
	Date     time.Time    `bson:"date"`
	PermitID string       `bson:"permitId,omitempty"`
}

type ActionIntent struct {
	UserID UserID    `json:"user_id"`
	Type   string    `json:"type"`
	Date   time.Time `json:"date"`
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
		Date:   action.Date,
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
