package entber

import (
	"fmt"
	"path"
	"strings"

	"entgo.io/ent/entc/gen"
	"entgo.io/ent/entc/load"
	"entgo.io/ent/schema/field"
)

var (
	snake      = gen.Funcs["snake"].(func(string) string)
	plural     = gen.Funcs["plural"].(func(string) string)
	buggyCamel = gen.Funcs["camel"].(func(string) string)
	camel      = func(s string) string { return buggyCamel(snake(s)) }
)

func init() {
	gen.Funcs["camel"] = camel
	gen.Funcs["tag"] = tag
	gen.Funcs["imports"] = imports
	gen.Funcs["null_field_create"] = null_field_create
	gen.Funcs["null_field_update"] = null_field_update
	gen.Funcs["extract_type"] = extract_type
	gen.Funcs["extract_type_info"] = extract_type_info
	gen.Funcs["edge_field"] = edge_field
	gen.Funcs["is_comparable"] = is_comparable
	gen.Funcs["enum_or_edge_filed"] = enum_or_edge_filed
	gen.Funcs["get_name"] = get_name
	gen.Funcs["get_type"] = get_type
	gen.Funcs["get_type_info"] = get_type_info
	gen.Funcs["is_slice"] = is_slice
	gen.Funcs["id_type"] = id_type
	gen.Funcs["order_fields"] = order_fields
	gen.Funcs["select_fields"] = select_fields
	gen.Funcs["dir"] = path.Dir
	gen.Funcs["is_base"] = is_base

}

func is_base(f *load.Field) bool {
	return f.Name == "id" || f.Name == "created_at" || f.Name == "updated_at" || f.Immutable ||
		f.Name == "ID" || f.Name == "createdAt" || f.Name == "updatedAt" || f.Name == "CreatedAt" || f.Name == "UpdatedAt"
}

func tag(f *load.Field) string {
	if f.Tag == "" {
		name := f.Name
		if strings.HasSuffix(name, "ID") {
			name = strings.TrimSuffix(name, "ID")
			name += "Id"
		}
		return fmt.Sprintf("json:\"%s,omitempty\"", name)
	}

	if !strings.Contains(f.Tag, "json") {
		name := camel(f.Name)
		f.Tag = fmt.Sprintf("json:\"%s,omitempty\" %s", name, f.Tag)

	}
	return f.Tag
}

func imports(g *gen.Graph, isInput ...bool) []string {
	imps := []string{}
	for _, s := range g.Schemas {
		for _, f := range s.Fields {
			if len(f.Enums) > 0 && len(isInput) > 0 && isInput[0] {
				imps = append(imps, path.Join(g.Package, strings.Split(f.Info.Ident, ".")[0]))
			}
			if f.Info != nil && len(f.Info.PkgPath) != 0 {
				if !in(f.Info.PkgPath, imps) {
					imps = append(imps, f.Info.PkgPath)
				}
			}
		}
	}

	return imps
}

func null_field_create(f *load.Field) bool {
	return f.Optional || f.Default
}

func null_field_update(f *load.Field) bool {
	return !strings.HasPrefix(extract_type(f), "[]")
}

func extract_type_info(t *field.TypeInfo) string {
	if t.Ident != "" {
		return t.Ident
	}
	return t.Type.String()
}

func extract_type(f *load.Field) string {
	return extract_type_info(f.Info)
}

func edge_field(e *load.Edge) bool {
	return e.Field != ""
}

func is_comparable(f *load.Field) bool {
	return has_prefixes(extract_type(f), []string{
		"string",
		"int",
		"uint",
		"float",
		"time.Time",
	})
}

func enum_or_edge_filed(s *load.Schema, f *load.Field) bool {
	for _, e := range s.Edges {
		if e.Field == f.Name {
			return extract_type(f) == "enum"
		}
	}
	return false
}

func get_name(f *load.Field) string {
	n := camel(f.Name)
	if strings.HasSuffix(n, "ID") {
		n = strings.TrimSuffix(n, "ID") + "Id"
	}
	return n
}

func get_type_info(f *field.TypeInfo) string {
	s := extract_type_info(f)
	t := "any"
	slice := false
	if strings.HasPrefix(s, "[]") {
		slice = true
		s = strings.TrimPrefix(s, "[]")
	}
	for k, v := range go_ts {
		if strings.HasPrefix(s, k) {
			t = v
			break
		}
	}

	if slice {
		return t + "[]"
	}
	return t
}

func get_type(f *load.Field) string {
	if len(f.Enums) > 0 {
		enums := []string{}
		for _, v := range f.Enums {
			enums = append(enums, "\""+v.V+"\"")
		}
		return strings.Join(enums, " | ")
	} else {
		s := extract_type(f)

		t := "any"
		slice := false
		if strings.HasPrefix(s, "[]") {
			slice = true
			s = strings.TrimPrefix(s, "[]")
		}
		for k, v := range go_ts {
			if strings.HasPrefix(s, k) {
				t = v
				break
			}
		}

		if slice {
			return t + "[]"
		}
		return t
	}
}

func is_slice(f *load.Field) bool {
	return strings.HasSuffix(get_type(f), "[]")
}

func id_type(s *load.Schema) string {
	for _, f := range s.Fields {
		if strings.ToLower(f.Name) == "id" {
			return get_type(f)
		}
	}
	return "number"
}

func order_fields(s *load.Schema) string {
	fields := []string{}
	for _, f := range s.Fields {
		if orderable(f) {
			fields = append(fields, snake(get_name(f)))
		}
	}
	return "\"" + strings.Join(fields, "\" | \"") + "\""
}

func select_fields(s *load.Schema) string {
	fields := []string{}
	for _, f := range s.Fields {
		fields = append(fields, snake(get_name(f)))
	}
	return "\"" + strings.Join(fields, "\" | \"") + "\""
}

func orderable(f *load.Field) bool {
	return has_prefixes(extract_type(f), []string{
		"string",
		"int",
		"uint",
		"float",
		"time.Time",
		"bool",
	})
}
