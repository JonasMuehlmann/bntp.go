// Copyright © 2021-2022 Jonas Muehlmann
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the "Software"),
// to deal in the Software without restriction, including without limitation
// the rights to use, copy, modify, merge, publish, distribute, sublicense,
// and/or sell copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included
// in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
// DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
// TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE
// OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package main

import (
	"os"
	"text/template"

	repository "github.com/JonasMuehlmann/bntp.go/repository/sqlite3"
	"github.com/JonasMuehlmann/bntp.go/tools"
	"github.com/JonasMuehlmann/goaoi"
)

type Entity struct {
	EntityName string
	Struct     any
}

var entities = []Entity{
	{"Document", repository.Document{}},
	{"Bookmark", repository.Bookmark{}},
	{"Tag", repository.Tag{}},
}

type Database struct {
	DatabaseName string
}

var databases = []Database{
	{"mssql"},
	{"mysql"},
	{"psql"},
	{"sqlite3"},
}

type Configuration struct {
	EntityName     string
	DatabaseName   string
	StructFields   []tools.StructField
	RelationFields []tools.StructField
}

var structNameVarTemplaterFragment = `{{$StructName := print (UppercaseBeginning .DatabaseName) (UppercaseBeginning .EntityName) "Repository" -}}
{{$EntityName := UppercaseBeginning .EntityName -}}`

var structDefinition = `type {{UppercaseBeginning .DatabaseName}}{{UppercaseBeginning .EntityName}}Repository struct {
    db sql.DB
}`

var repositoryHelperTypesFragment = structNameVarTemplaterFragment + `
type {{$EntityName}}Field string

var {{$EntityName}}Fields = struct {
    {{range $field := .StructFields -}}
    {{.FieldName}}  {{$EntityName}}Field
    {{end}}
}{
    {{range $field := .StructFields -}}
    {{.FieldName}}: "{{.LogicalFieldName -}}",
    {{end}}
}

var {{$EntityName}}FieldsList = []{{$EntityName}}Field{
    {{range $field := .StructFields -}}
    {{$EntityName}}Field("{{.FieldName}}"),
    {{end}}
}

var {{$EntityName}}RelationsList = []string{
    {{range $relation := .RelationFields -}}
    "{{.FieldName}}",
    {{end}}
}

type {{$EntityName}}Filter struct {
    {{range $field := .StructFields -}}
    {{.FieldName}} optional.Optional[FilterOperation[{{.FieldType}}]]
    {{end}}
    {{range $relation := .RelationFields -}}
    {{.FieldName}} optional.Optional[UpdateOperation[{{.FieldType}}]]
    {{end}}
}

type {{$EntityName}}Updater struct {
    {{range $field := .StructFields -}}
    {{.FieldName}} optional.Optional[UpdateOperation[{{.FieldType}}]]
    {{end}}
    {{range $relation := .RelationFields -}}
    {{.FieldName}} optional.Optional[UpdateOperation[{{.FieldType}}]]
    {{end}}
}

type {{$StructName}}Hook func(context.Context, {{$StructName}}) error`

// TODO: Add template for type safe config struct per provider for New() methods
// It could embed a generic RepositoryConfig into repository specific configurations e.g. Sqlite3RepositoryConfig

func main() {
	for _, database := range databases {
		for _, entity := range entities {
			tmplRaw, err := os.ReadFile("templates/sql_repositories/" + tools.LowercaseBeginning(entity.EntityName) + "_repository.go.tpl")
			if err != nil {
				panic(err)
			}

			tmpl, err := template.New("structDefinition").Funcs(tools.FullFuncMap).Parse(structDefinition)
			if err != nil {
				panic(err)
			}

			tmpl, err = tmpl.New("repositoryHelperTypes").Funcs(tools.FullFuncMap).Parse(repositoryHelperTypesFragment)
			if err != nil {
				panic(err)
			}

			tmpl, err = tmpl.New(tools.LowercaseBeginning(entity.EntityName) + "_repository").Funcs(tools.FullFuncMap).Parse(structNameVarTemplaterFragment + string(tmplRaw))
			if err != nil {
				panic(err)
			}

			outFile, err := os.Create("repository/" + database.DatabaseName + "/" + tools.LowercaseBeginning(entity.EntityName) + "_repository.go")
			if err != nil {
				panic(err)
			}

			entityStruct := tools.NewStructModel(entity.Struct)

			entityStruct.StructFields, err = goaoi.CopyExceptIfSlice(entityStruct.StructFields, func(s tools.StructField) bool { return s.FieldName == "R" || s.FieldName == "L" })
			if err != nil {
				panic(err)
			}

			var relationStruct tools.Struct

			switch e := entity.Struct.(type) {
			case repository.Bookmark:
				e.R = e.R.NewStruct()
				relationStruct = tools.NewStructModel(*e.R)
			case repository.Document:
				e.R = e.R.NewStruct()
				relationStruct = tools.NewStructModel(*e.R)
			case repository.Tag:
				e.R = e.R.NewStruct()
				relationStruct = tools.NewStructModel(*e.R)

			default:
				panic("Unhandled sql repository type")
			}

			err = tmpl.Execute(outFile, Configuration{
				EntityName:     entity.EntityName,
				DatabaseName:   database.DatabaseName,
				StructFields:   entityStruct.StructFields,
				RelationFields: relationStruct.StructFields,
			})
			if err != nil {
				panic(err)
			}
		}
	}
}
