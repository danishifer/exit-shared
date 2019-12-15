package instituteService

import exit "exit-shared"

type GetTerminalByIdRequest struct {
	TerminalID string `json:"terminal_id"`
}

type GetTerminalByIdResponse struct {
	Terminal exit.Terminal `json:"terminal"`
}

func (r *GetTerminalByIdRequest) IsValid() bool {
	if r.TerminalID == "" {
		return false
	}
	return true
}
