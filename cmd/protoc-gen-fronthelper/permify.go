package main

import (
	"path"
	"reflect"
	"slices"
	"strings"
	"text/template"

	pgs "github.com/lyft/protoc-gen-star/v2"
	pgsgo "github.com/lyft/protoc-gen-star/v2/lang/go"
)

// PermifyPlugin
type PermifyModule struct {
	*pgs.ModuleBase
	ctx pgsgo.Context
	tpl *template.Template
}

var fns = template.FuncMap{
	"last": func(x int, a interface{}) bool {
		return x == reflect.ValueOf(a).Len()-1
	},
}

// Permify returns an initialized PermifyPlugin
func Permify() *PermifyModule {
	return &PermifyModule{ModuleBase: &pgs.ModuleBase{}}
}

func (p *PermifyModule) InitContext(c pgs.BuildContext) {
	p.ModuleBase.InitContext(c)
	p.ctx = pgsgo.InitContext(c.Parameters())

	tpl := template.New("permify").Funcs(fns)

	p.tpl = template.Must(tpl.Parse(permsTpl))
}

// Name satisfies the generator.Plugin interface.
func (p *PermifyModule) Name() string { return "permify" }

func (p *PermifyModule) Execute(targets map[string]pgs.File, pkgs map[string]pgs.Package) []pgs.Artifact {
	visited := map[string][]pgs.File{}
	for _, t := range targets {
		key := t.File().InputPath().Dir().String()
		if _, ok := visited[key]; ok {
			visited[key] = append(visited[key], t)
			continue
		}

		visited[key] = []pgs.File{t}
	}

	data := struct {
		FS          []pgs.File
		Permissions map[string]map[string]*Perm
	}{
		FS:          []pgs.File{},
		Permissions: map[string]map[string]*Perm{},
	}
	for _, fs := range visited {
		slices.SortFunc(fs, func(a, b pgs.File) int {
			return strings.Compare(a.File().InputPath().String(), b.File().InputPath().String())
		})

		data.FS = append(data.FS, fs...)

		for _, f := range fs {
			for _, s := range f.Services() {
				sName := string(s.Name())

				p.Debugf("Service: %s (%s)", sName, sName)

				for _, m := range s.Methods() {
					mName := string(m.Name())
					mName = strings.TrimPrefix(mName, "services.")

					comment := m.SourceCodeInfo().LeadingComments()
					comment = strings.TrimLeft(comment, " ")
					if !strings.HasPrefix(comment, "@perm") {
						continue
					}
					comment = strings.TrimRight(comment, "\n")

					perm, err := p.parseComment(sName, mName, comment)
					if err != nil {
						p.Failf("failed to parse comment in %s method %s (comment: '%s'), error: %w", f.InputPath(), mName, comment, err)
						return nil
					}
					if perm == nil {
						p.Failf("failed to parse comment in %s method %s (comment: '%s')", f.InputPath(), mName, comment)
						return nil
					}

					if perm.Name != mName {
						continue
					}

					if perm.Name == "SuperUser" || perm.Name == "Any" {
						continue
					}

					if _, ok := data.Permissions[sName]; !ok {
						data.Permissions[sName] = map[string]*Perm{}
					}
					if _, ok := data.Permissions[sName][perm.Name]; !ok {
						data.Permissions[sName][perm.Name] = perm
						p.Debugf("Permission added: %q - %+v\n", mName, perm)
					} else {
						p.Debugf("Permission already in list: %q - %+v\n", mName, perm)
					}
				}
			}
		}
	}

	slices.SortStableFunc(data.FS, func(a, b pgs.File) int {
		return strings.Compare(a.FullyQualifiedName(), b.FullyQualifiedName())
	})

	p.AddGeneratorTemplateFile(path.Join("perms.ts"), p.tpl, data)

	return p.Artifacts()
}

func (p *PermifyModule) parseComment(service string, method string, comment string) (*Perm, error) {
	comment = strings.TrimPrefix(comment, "@perm: ")
	comment = strings.TrimPrefix(comment, "@perm")

	perm := &Perm{
		Name: method,
	}

	if comment == "" {
		return perm, nil
	}

	split := strings.Split(comment, ";")

	for i := 0; i < len(split); i++ {
		k, v, _ := strings.Cut(split[i], "=")
		if v == "" {
			continue
		}

		switch strings.ToLower(k) {
		case "name":
			perm.Name = v
			continue
		case "attrs":
			for _, v := range strings.Split(v, "|") {
				attrSplit := strings.Split(v, "/")
				if len(attrSplit) <= 1 {
					p.Fail("Invalid attrs value found: ", v)
				}

				attrType := attrSplit[1]
				validValue := ""
				validList := strings.Split(attrSplit[1], ":")
				if len(validList) > 1 {
					attrType = validList[0]
					validValue = strings.Join(validList[1:], ":")
				}

				perm.Attrs = append(perm.Attrs, Attr{
					Key:   attrSplit[0],
					Type:  attrType,
					Valid: validValue,
				})
			}
			continue
		}
	}

	return perm, nil
}

const permsTpl = `// Code generated by protoc-gen-fronthelper. DO NOT EDIT.
{{- range $f := .FS }}
// source: {{ $f.File.InputPath }}
{{- end }}

export type Perms =
    | 'CanBeSuper'
    | 'SuperUser'
    | 'TODOService.TODOMethod'
{{- range $sName, $service := $.Permissions -}}
	{{- range $i, $perm := $service }}
	| '{{ $sName }}.{{ $perm.Name }}'
		{{- end }}
{{- end -}};
`

type Perm struct {
	Name  string
	Attrs []Attr
}

type Attr struct {
	Key     string
	Type    string
	Valid   string
	Default string
}
