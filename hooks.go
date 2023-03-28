package entber

import (
	"path"

	"entgo.io/ent/entc/gen"
)

func (e *extension) generate(next gen.Generator) gen.Generator {
	return gen.GenerateFunc(func(g *gen.Graph) error {

		if e.data.FiberConfig != nil {
			s := parseTemplate("fiber/routes", e.data)
			writeFile(path.Join(e.data.FiberConfig.RoutesPath, "routes.go"), s)

			for _, schema := range g.Schemas {
				data := data{
					Graph:         g,
					CurrentSchema: schema,
				}
				s := parseTemplate("fiber/handler", data)
				writeFile(path.Join(e.data.FiberConfig.HandlersPath, schema.Name+".go"), s)
			}
		}

		if e.data.DBConfig != nil {
			s := parseTemplate("ent/db", e.data)
			writeFile(path.Join(e.data.DBConfig.Path, "db.go"), s)
		}

		if e.data.TSConfig != nil {
			s := parseTemplate("ts/api", e.data)
			writeFile(path.Join("", "db.go"), s)
			s = parseTemplate("ts/types", e.data)
			writeFile(path.Join("", "db.go"), s)
		}

		return next.Generate(g)
	})
}
