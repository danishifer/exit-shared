package exit_shared_mongo

import "time"

type Permit struct {
	ID           string         `bson:"id"`
	Status       PermitStatus   `bson:"status"`
	Type         string         `bson:"type"`
	AssignedTo   []UserID       `bson:"assignedTo"`
	RequestedOn  time.Time      `bson:"requestedOn"`
	RequestedBy  UserID         `bson:"requestedBy"`
	Users        []UserID       `bson:"users"`
	Schedule     PermitSchedule `bson:"schedule"`
	Note         string         `bson:"note"`
	ApprovedOn   time.Time      `bson:"approvedOn,omitempty"`
	ApprovedBy   UserID         `bson:"approvedBy,omitempty"`
	ApprovedNote string         `bson:"approvedNote,omitempty"`
	RejectedOn   time.Time      `bson:"rejectedOn,omitempty"`
	RejectedBy   UserID         `bson:"rejectedBy,omitempty"`
	RejectedNote string         `bson:"rejectedNote,omitempty"`
}

type PermitSchedule struct {
	Repeat        string    `bson:"repeat"`
	RepeatEndDate time.Time `bson:"repeatEndDate,omitempty"`
	StartDate     time.Time `bson:"startDate"`
	EndDate       time.Time `bson:"endDate"`
}

type PermitStatus int32

const (
	PermitStatusPending  PermitStatus = 0
	PermitStatusApproved PermitStatus = 1
	PermitStatusRejected PermitStatus = 2
)

type PermitForActionRequest struct {
	UserID UserID    `json:"user_id"`
	Type   string    `json:"type"`
	Date   time.Time `json:"date"`
}

func (request *PermitForActionRequest) IsValid() bool {
	if request.UserID == "" || request.Type == "" {
		return false
	}
	return true
}

func (permit *Permit) IsValidForAction(r PermitForActionRequest) bool {
	// TODO: implement permit validation
	return true
}
