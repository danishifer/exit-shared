package instituteService

import exit "exit-shared"

type GetTerminalByCardSerialRequest struct {
	CardSerial string `json:"card_serial"`
}

type GetTerminalByCardSerialResponse struct {
	Terminal exit.Terminal `json:"terminal"`
}

func (r *GetTerminalByCardSerialRequest) IsValid() bool {
	if r.CardSerial == "" {
		return false
	}
	return true
}
