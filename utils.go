package entber

import (
	"bytes"
	"log"
	"strings"
	"text/template"

	"entgo.io/ent/entc/gen"
)

func parseTemplate(name string, data any) string {
	in, err := templates.ReadFile("templates/" + name + ".tmpl")
	if err != nil {
		log.Fatalln(err)
	}
	t, err := template.New(name).Funcs(gen.Funcs).Parse(string(in))
	if err != nil {
		log.Fatalln(err)
	}
	out := new(bytes.Buffer)
	err = t.Execute(out, data)
	if err != nil {
		log.Fatalln(err)
	}
	return out.String()
}

func in[T comparable](v T, vs []T) bool {
	for _, v2 := range vs {
		if v == v2 {
			return true
		}
	}
	return false
}

func has_prefixes(s string, px []string) bool {
	for _, p := range px {
		if strings.HasPrefix(s, p) {
			return true
		}
	}
	return false
}
