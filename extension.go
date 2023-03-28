package entber

import (
	"embed"

	"entgo.io/ent/entc/gen"
)

//go:embed templates
var templates embed.FS

func (e *extension) Hooks() []gen.Hook {
	return e.hooks
}

func NewExtension(config *Config) *extension {
	e := new(extension)
	e.config = config
	e.hooks = append(e.hooks, e.generate)
	return e
}
