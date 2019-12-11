package usersService

import exit "exit-shared-mongo"

type GetStudentByIDRequest struct {
	UserID exit.UserID `json:"user_id"`
}

type GetStudentByIDResponse struct {
	Student exit.Student `json:"student"`
}

func (r *GetStudentByIDRequest) IsValid() bool {
	if r.UserID == "" {
		return false
	}
	return true
}
