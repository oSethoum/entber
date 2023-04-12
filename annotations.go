package entber

import "encoding/json"

const (
	skipAnootationName      = "entris-skip"
	SkipFieldCreate    uint = iota
	SkipFieldUpdate
	SkipFieldQuery
	SkipFieldType
	SkipSchemaQuery
	SkipSchemaCreate
	SkipSchemaUpdate
	SkipSchemaDelete
	SkipEdgeQuery
	SkipAll
)

func (a *skipAnnotation) Name() string {
	return a.name
}

func (a *skipAnnotation) decode(v interface{}) error {
	buffer, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return json.Unmarshal(buffer, a)
}

func Skip(skips ...uint) *skipAnnotation {
	return &skipAnnotation{
		name:  skipAnootationName,
		Skips: skips,
	}
}
