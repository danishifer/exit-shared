package exit_shared

import "time"

type Permit struct {
	ID           string         `bson:"id" json:"id"`
	Status       PermitStatus   `bson:"status" json:"status"`
	Type         string         `bson:"type" json:"type"`
	AssignedTo   []UserID       `bson:"assignedTo" json:"assigned_to"`
	RequestedOn  *time.Time     `bson:"requestedOn" json:"requested_on"`
	RequestedBy  UserID         `bson:"requestedBy" json:"requested_by"`
	Users        []UserID       `bson:"users" json:"users"`
	Schedule     PermitSchedule `bson:"schedule" json:"schedule"`
	Title        string         `bson:"title,omitempty" json:"title,omitempty"`
	Note         string         `bson:"note" json:"note"`
	ApprovedOn   *time.Time     `bson:"approvedOn,omitempty" json:"approved_on"`
	ApprovedBy   UserID         `bson:"approvedBy,omitempty" json:"approved_by"`
	ApprovedNote string         `bson:"approvedNote,omitempty" json:"approved_note"`
	RejectedOn   *time.Time     `bson:"rejectedOn,omitempty" json:"rejected_on"`
	RejectedBy   UserID         `bson:"rejectedBy,omitempty" json:"rejected_by"`
	RejectedNote string         `bson:"rejectedNote,omitempty" json:"rejected_note"`
}

type PermitSchedule struct {
	Repeat        string    `bson:"repeat" json:"repeat"`
	RepeatEndDate time.Time `bson:"repeatEndDate,omitempty" json:"repeat_end_date"`
	StartDate     time.Time `bson:"startDate" json:"start_date"`
	EndDate       time.Time `bson:"endDate" json:"end_date"`
}

type PermitStatus int32

const (
	PermitStatusPending  PermitStatus = 0
	PermitStatusApproved PermitStatus = 1
	PermitStatusRejected PermitStatus = 2
)

func (permit *Permit) IsValidForActionIntent(r ActionIntent) bool {
	// TODO: implement permit validation
	if permit.Type != r.Type {
		return false
	}
	return true
}
