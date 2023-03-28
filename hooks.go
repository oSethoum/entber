package entber

import (
	"log"
	"os"
	"path"

	"entgo.io/ent/entc/gen"
)

func (e *extension) generate(next gen.Generator) gen.Generator {
	return gen.GenerateFunc(func(g *gen.Graph) error {
		s := parseTemplate("routes", g)
		err := os.WriteFile(path.Join(g.Target, "routes"), []byte(s), 0666)
		if err != nil {
			log.Fatalln(err)
		}

		for _, schema := range g.Schemas {
			data := data{
				Graph:         g,
				CurrentSchema: schema,
			}
			s := parseTemplate("routes", data)
			err := os.WriteFile(path.Join(g.Target, "routes"), []byte(s), 0666)
			if err != nil {
				log.Fatalln(err)
			}
		}

		return next.Generate(g)
	})
}
