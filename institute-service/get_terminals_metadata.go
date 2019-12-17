package instituteService

import exit "exit-shared"

type GetTerminalsMetadataRequest struct {
}

type GetTerminalsMetadataResponse struct {
	Terminals []exit.TerminalMetadata `json:"terminals"`
}

func (r *GetTerminalsMetadataRequest) IsValid() bool {
	return true
}
