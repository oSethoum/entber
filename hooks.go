package entber

import (
	"path"

	"entgo.io/ent/entc/gen"
)

func (e *extension) generate(next gen.Generator) gen.Generator {
	return gen.GenerateFunc(func(g *gen.Graph) error {
		e.data.Graph = g
		destination := rootDir()

		if e.data.AppConfig != nil {
			destination = path.Join(destination, e.data.AppConfig.Path)
		}

		s := parseTemplate("ent/input", e.data)
		writeFile(path.Join(destination, "ent/input.go"), s)

		s = parseTemplate("ent/query", e.data)
		writeFile(path.Join(destination, "ent/query.go"), s)

		s = parseTemplate("ent/errors", e.data)
		writeFile(path.Join(destination, "ent/errors.go"), s)

		if e.data.WithFiber {
			s = parseTemplate("fiber/routes/routes", e.data)
			writeFile(path.Join(destination, "routes/routes.go"), s)

			s = parseTemplate("fiber/handlers/helper", e.data)
			writeFile(path.Join(destination, "handlers/helper.go"), s)

			s = parseTemplate("fiber/handlers/ws", e.data)
			writeFile(path.Join(destination, "handlers/ws.go"), s)

			s = parseTemplate("fiber/utils/utils", e.data)
			writeFile(path.Join(destination, "utils/utils.go"), s)

			for _, schema := range g.Schemas {
				e.data.CurrentSchema = schema
				s := parseTemplate("fiber/handlers/handler", e.data)
				writeFile(path.Join(destination, "handlers", snake(plural(schema.Name))+".go"), s)
			}
		}

		if e.data.DBConfig != nil {
			s := parseTemplate("ent/db", e.data)
			writeFile(path.Join(destination, e.data.DBConfig.Path, "db.go"), s)
		}

		if e.data.TSConfig != nil {
			s := parseTemplate("ts/api", e.data)
			writeFile(path.Join(e.data.TSConfig.ApiPath, "api.ts"), s)
			s = parseTemplate("ts/types", e.data)
			writeFile(path.Join(e.data.TSConfig.ApiPath, "types.ts"), s)
		}

		return next.Generate(g)
	})
}
