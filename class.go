package exit_shared

type Class struct {
	ID   string `bson:"id" json:"id"`
	Name string `bson:"name" json:"name"`
}

type ClassMetadata struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (t *Class) Metadata() ClassMetadata {
	return ClassMetadata{
		ID:   t.ID,
		Name: t.Name,
	}
}
