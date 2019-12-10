package exit_shared_mongo

type PermitForActionIntentRequest ActionIntent

type PermitForActionIntentResponse struct {
	Approved bool   `json:"approved"`
	Permit   string `json:"permit"`
}

func (request *PermitForActionIntentRequest) IsValid() bool {
	if request.UserID == "" || request.Type == "" {
		return false
	}
	return true
}
