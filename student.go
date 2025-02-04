package exit_shared

type Student struct {
	ID   UserID `bson:"id" json:"id"`
	Name struct {
		Given  string `bson:"given" json:"given"`
		Middle string `bson:"middle" json:"middle"`
		Family string `bson:"family" json:"family"`
	} `bson:"name" json:"name"`
	Homeroom string `bson:"homeroom" json:"homeroom"`
}
