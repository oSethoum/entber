package entber

import (
	"path"

	"entgo.io/ent/entc/gen"
)

func (e *extension) generate(next gen.Generator) gen.Generator {
	return gen.GenerateFunc(func(g *gen.Graph) error {
		e.data.Graph = g

		s := parseTemplate("ent/input", e.data)
		writeFile("ent/input.go", s)

		s = parseTemplate("ent/query", e.data)
		writeFile("ent/query.go", s)

		s = parseTemplate("ent/errors", e.data)
		writeFile("ent/errors.go", s)

		if e.data.FiberConfig != nil {
			s = parseTemplate("fiber/config/config", e.data)
			writeFile(path.Join("config", "config.go"), s)

			s = parseTemplate("fiber/routes/init", e.data)
			writeFile(path.Join(e.data.FiberConfig.RoutesPath, "init.go"), s)

			s = parseTemplate("fiber/routes/crud", e.data)
			writeFile(path.Join(e.data.FiberConfig.RoutesPath, "crud.go"), s)

			s = parseTemplate("fiber/handlers/util", e.data)
			writeFile(path.Join(e.data.FiberConfig.HandlersPath, "util.go"), s)

			if e.data.FiberConfig.WithEvents {
				s = parseTemplate("fiber/handlers/ws", e.data)
				writeFile(path.Join(e.data.FiberConfig.HandlersPath, "ws.go"), s)
			}

			if e.data.FiberConfig.WithUpload {
				s = parseTemplate("fiber/handlers/upload", e.data)
				writeFile(path.Join(e.data.FiberConfig.HandlersPath, "upload.go"), s)
			}

			if !e.data.FiberConfig.WithoutFiberAuth {
				s = parseTemplate("fiber/handlers/auth", e.data)
				writeFile(path.Join(e.data.FiberConfig.HandlersPath, "auth.go"), s)

				s = parseTemplate("fiber/middleware/auth", e.data)
				writeFile(path.Join(e.data.FiberConfig.MiddlewarePath, "auth.go"), s)

				s = parseTemplate("fiber/middleware/ws", e.data)
				writeFile(path.Join(e.data.FiberConfig.MiddlewarePath, "ws.go"), s)
			}

			for _, schema := range g.Schemas {
				if skip_schema_query(schema) && skip_schema_create(schema) &&
					skip_schema_update(schema) && skip_schema_delete(schema) {
					continue
				}
				e.data.CurrentSchema = schema
				s := parseTemplate("fiber/handlers/handler", e.data)
				writeFile(path.Join(e.data.FiberConfig.HandlersPath, snake(plural(schema.Name))+".go"), s)
			}
		}

		if e.data.DBConfig != nil {
			s := parseTemplate("ent/db", e.data)
			writeFile(path.Join(e.data.DBConfig.Path, "db.go"), s)
		}

		if e.data.TSConfig != nil {
			s := parseTemplate("ts/api", e.data)
			writeFile(path.Join(e.data.TSConfig.ApiPath, "api.ts"), s)
			s = parseTemplate("ts/realtime", e.data)
			writeFile(path.Join(e.data.TSConfig.ApiPath, "realtime.ts"), s)
			s = parseTemplate("ts/types", e.data)
			writeFile(path.Join(e.data.TSConfig.ApiPath, "types.ts"), s)
		}

		return next.Generate(g)
	})
}
