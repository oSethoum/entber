package entber

import (
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"entgo.io/ent/entc/load"
)

type extension struct {
	entc.DefaultExtension
	hooks []gen.Hook
	data  data
}

type Driver = string

const (
	SQLite     Driver = "sqlite"
	MySQL      Driver = "mysql"
	PostgreSQL Driver = "postgres"
)

type option = func(*extension)

type data struct {
	*gen.Graph
	DBConfig      *DBConfig
	TSConfig      *TSConfig
	FiberConfig   *FiberConfig
	CurrentSchema *load.Schema
}

type DBConfig struct {
	Path   string
	Driver string
	Dsn    string
}

type TSConfig struct {
	TypesPath string
	ApiPath   string
}

type FiberConfig struct {
	HandlersPath string
	RoutesPath   string
}

type comparable interface{ ~string | ~int | ~float32 }
