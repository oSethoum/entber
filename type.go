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
	SQLite     Driver = "sqlite3"
	MySQL      Driver = "mysql"
	PostgreSQL Driver = "postgres"
)

type option = func(*extension)

type data struct {
	*gen.Graph
	DBConfig    *DBConfig
	TSConfig    *TSConfig
	FiberConfig *FiberConfig
	WithFiber   bool
	AppConfig   *AppConfig

	CurrentSchema *load.Schema
}

type DBConfig struct {
	Path   string
	Driver string
	Dsn    string
}

type TSConfig struct {
	ApiPath   string
	TypesOnly bool
}

type FiberConfig struct {
}

type AppConfig struct {
	Path string
}

type comparable interface {
	~string | ~int | ~float32 | ~uint
}

var go_ts = map[string]string{
	"time.Time": "Date",
	"bool":      "boolean",
	"int":       "number",
	"uint":      "number",
	"float":     "number",
	"enum":      "string",
	"string":    "string",
	"any":       "any",
	"other":     "any",
	"json":      "any",
}
