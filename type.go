package entber

import (
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"entgo.io/ent/entc/load"
)

type extension struct {
	entc.DefaultExtension
	hooks  []gen.Hook
	config *Config
}

type data struct {
	*gen.Graph
	CurrentSchema *load.Schema
}

type Config struct {
	HandlersPath string
	RoutesPath   string
}

type comparable interface{ ~string | ~int | ~float32 }
