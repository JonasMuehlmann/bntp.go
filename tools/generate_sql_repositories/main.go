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

	repository "github.com/JonasMuehlmann/bntp.go/model/repositoryite3"

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
    db *sql.DB
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
    {{.FieldName}} optional.Optional[model.FilterOperation[{{Unslice (UnaliasSQLBoilerSlice .FieldType)}}]]
    {{end}}
    {{range $relation := .RelationFields -}}
    {{.FieldName}} optional.Optional[model.FilterOperation[{{Unslice (UnaliasSQLBoilerSlice .FieldType)}}]]
    {{end}}
}

type {{$EntityName}}FilterMapping[T any] struct {
    Field {{$EntityName}}Field
    FilterOperation model.FilterOperation[T]
}

func (filter *{{$EntityName}}Filter) GetSetFilters() *list.List {
    setFilters := list.New()

    {{range $field := .StructFields -}}
    if filter.{{.FieldName}}.HasValue {
    setFilters.PushBack({{$EntityName}}FilterMapping[{{Unslice (UnaliasSQLBoilerSlice .FieldType)}}]{Field: {{$EntityName}}Fields.{{.FieldName}}, FilterOperation: filter.{{.FieldName}}.Wrappee})
    }
    {{end}}

    return setFilters
}

type {{$EntityName}}Updater struct {
    {{range $field := .StructFields -}}
    {{.FieldName}} optional.Optional[model.UpdateOperation[{{.FieldType}}]]
    {{end}}
    {{range $relation := .RelationFields -}}
    {{.FieldName}} optional.Optional[model.UpdateOperation[{{.FieldType}}]]
    {{end}}
}

type {{$EntityName}}UpdaterMapping[T any] struct {
    Field {{$EntityName}}Field
    Updater model.UpdateOperation[T]
}

func (updater *{{$EntityName}}Updater) GetSetUpdaters() *list.List {
    setUpdaters := list.New()

    {{range $field := .StructFields -}}
    if updater.{{.FieldName}}.HasValue {
    setUpdaters.PushBack({{$EntityName}}UpdaterMapping[{{.FieldType}}]{Field: {{$EntityName}}Fields.{{.FieldName}}, Updater: updater.{{.FieldName}}.Wrappee})
    }
    {{end}}

    return setUpdaters
}

func (updater *{{$EntityName}}Updater) ApplyToModel({{LowercaseBeginning $EntityName}}Model *{{$EntityName}}) {
    {{range $field := .StructFields -}}
    if updater.{{.FieldName}}.HasValue {
        model.ApplyUpdater(&(*{{LowercaseBeginning $EntityName}}Model).{{.FieldName}}, updater.{{.FieldName}}.Wrappee)
    }
    {{end}}
}

type {{$StructName}}Hook func(context.Context, {{$StructName}}) error

type queryModSlice{{$EntityName}} []qm.QueryMod

func (s queryModSlice{{$EntityName}}) Apply(q *queries.Query) {
    qm.Apply(q, s...)
}

func buildQueryModFilter{{$EntityName}}[T any](filterField {{$EntityName}}Field, filterOperation model.FilterOperation[T]) queryModSlice{{$EntityName}} {
    var newQueryMod queryModSlice{{$EntityName}}

    filterOperator := filterOperation.Operator

    switch filterOperator {
    case model.FilterEqual:
        filterOperand, ok := filterOperation.Operand.(model.ScalarOperand[any])
        if !ok {
            panic("Expected a scalar operand for FilterEqual operator")
        }

        newQueryMod = append(newQueryMod, qm.Where(string(filterField)+" = ?", filterOperand.Operand))
    case model.FilterNEqual:
        filterOperand, ok := filterOperation.Operand.(model.ScalarOperand[any])
        if !ok {
            panic("Expected a scalar operand for FilterNEqual operator")
        }

        newQueryMod = append(newQueryMod, qm.Where(string(filterField)+" != ?", filterOperand.Operand))
    case model.FilterGreaterThan:
        filterOperand, ok := filterOperation.Operand.(model.ScalarOperand[any])
        if !ok {
            panic("Expected a scalar operand for FilterGreaterThan operator")
        }

        newQueryMod = append(newQueryMod, qm.Where(string(filterField)+" > ?", filterOperand.Operand))
    case model.FilterGreaterThanEqual:
        filterOperand, ok := filterOperation.Operand.(model.ScalarOperand[any])
        if !ok {
            panic("Expected a scalar operand for FilterGreaterThanEqual operator")
        }

        newQueryMod = append(newQueryMod, qm.Where(string(filterField)+" >= ?", filterOperand.Operand))
    case model.FilterLessThan:
        filterOperand, ok := filterOperation.Operand.(model.ScalarOperand[any])
        if !ok {
            panic("Expected a scalar operand for FilterLessThan operator")
        }

        newQueryMod = append(newQueryMod, qm.Where(string(filterField)+" < ?", filterOperand.Operand))
    case model.FilterLessThanEqual:
        filterOperand, ok := filterOperation.Operand.(model.ScalarOperand[any])
        if !ok {
            panic("Expected a scalar operand for FilterLessThanEqual operator")
        }

        newQueryMod = append(newQueryMod, qm.Where(string(filterField)+" <= ?", filterOperand.Operand))
    case model.FilterIn:
        filterOperand, ok := filterOperation.Operand.(model.ListOperand[any])
        if !ok {
            panic("Expected a list operand for FilterIn operator")
        }

        newQueryMod = append(newQueryMod, qm.WhereIn(string(filterField)+" IN (?)", filterOperand.Operands))
    case model.FilterNotIn:
        filterOperand, ok := filterOperation.Operand.(model.ListOperand[any])
        if !ok {
            panic("Expected a list operand for FilterNotIn operator")
        }

        newQueryMod = append(newQueryMod, qm.WhereNotIn(string(filterField)+" IN (?)", filterOperand.Operands))
    case model.FilterBetween:
        filterOperand, ok := filterOperation.Operand.(model.RangeOperand[any])
        if !ok {
            panic("Expected a scalar operand for FilterBetween operator")
        }

        newQueryMod = append(newQueryMod, qm.Where(string(filterField)+" BETWEEN ? AND ?", filterOperand.Start, filterOperand.End))
    case model.FilterNotBetween:
        filterOperand, ok := filterOperation.Operand.(model.RangeOperand[any])
        if !ok {
            panic("Expected a scalar operand for FilterNotBetween operator")
        }

        newQueryMod = append(newQueryMod, qm.Where(string(filterField)+" NOT BETWEEN ? AND ?", filterOperand.Start, filterOperand.End))
    case model.FilterLike:
        filterOperand, ok := filterOperation.Operand.(model.ScalarOperand[any])
        if !ok {
            panic("Expected a scalar operand for FilterLike operator")
        }

        newQueryMod = append(newQueryMod, qm.Where(string(filterField)+" LIKE ?", filterOperand.Operand))
    case model.FilterNotLike:
        filterOperand, ok := filterOperation.Operand.(model.ScalarOperand[any])
        if !ok {
            panic("Expected a scalar operand for FilterLike operator")
        }

        newQueryMod = append(newQueryMod, qm.Where(string(filterField)+" NOT LIKE ?", filterOperand.Operand))
    case model.FilterOr:
        filterOperand, ok := filterOperation.Operand.(model.CompoundOperand[any])
        if !ok {
            panic("Expected a scalar operand for FilterOr operator")
        }
        newQueryMod = append(newQueryMod, qm.Expr(buildQueryModFilter{{$EntityName}}(filterField, filterOperand.LHS)))
        newQueryMod = append(newQueryMod, qm.Or2(qm.Expr(buildQueryModFilter{{$EntityName}}(filterField, filterOperand.RHS))))
    case model.FilterAnd:
        filterOperand, ok := filterOperation.Operand.(model.CompoundOperand[any])
        if !ok {
            panic("Expected a scalar operand for FilterAnd operator")
        }

        newQueryMod = append(newQueryMod, qm.Expr(buildQueryModFilter{{$EntityName}}(filterField, filterOperand.LHS)))
        newQueryMod = append(newQueryMod, qm.Expr(buildQueryModFilter{{$EntityName}}(filterField, filterOperand.RHS)))
    default:
        panic("Unhandled FilterOperator")
    }

    return newQueryMod
}

func buildQueryModListFromFilter{{$EntityName}}(setFilters list.List) queryModSlice{{$EntityName}} {
	queryModList := make(queryModSlice{{$EntityName}}, 0, {{len .StructFields}})

	for filter := setFilters.Front(); filter != nil; filter = filter.Next() {
		filterMapping, ok := filter.Value.({{$EntityName}}FilterMapping[any])
		if !ok {
			panic(fmt.Sprintf("Expected type %T but got %T", {{$EntityName}}FilterMapping[any]{}, filter))
		}

        newQueryMod := buildQueryModFilter{{$EntityName}}(filterMapping.Field, filterMapping.FilterOperation)

        queryModList = append(queryModList, newQueryMod...)
	}

	return queryModList
}`

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

			outFile, err := os.Create("model/repository/" + database.DatabaseName + "/" + tools.LowercaseBeginning(entity.EntityName) + "_repository.go")
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
