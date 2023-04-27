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

const (
	Deno = "deno"
	Node = "node"
)

type option = func(*extension)

type data struct {
	*gen.Graph
	DBConfig    *DBConfig
	TSConfig    *TSConfig
	FiberConfig *FiberConfig

	CurrentSchema *load.Schema
}

type DBConfig struct {
	Path   string
	Driver string
	Dsn    string
}

type AuthConfig struct {
	Token bool
}

type TSConfig struct {
	ApiPath string
	Runtime string
}

type FiberConfig struct {
	HandlersPath     string
	RoutesPath       string
	MiddlewaresPath  string
	WithUpload       bool
	WithEvents       bool
	WithAuth         bool
	WithoutFiberAuth bool
}

type comparable interface {
	~string | ~int | ~float32 | ~uint
}

var gots = map[string]string{
	"time.Time": "string",
	"bool":      "boolean",
	"int":       "number",
	"uint":      "number",
	"float":     "number",
	"enum":      "string",
	"any":       "any",
	"other":     "any",
	"json":      "any",
}

type skipAnnotation struct {
	name  string
	Skips []uint
}
