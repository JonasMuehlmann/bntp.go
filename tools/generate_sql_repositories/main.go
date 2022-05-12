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

	"github.com/JonasMuehlmann/bntp.go/tools"
)

type Entity struct {
	EntityName string
}

var entities = []Entity{
	{"Document"},
	{"Bookmark"},
	{"Tag"},
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
	EntityName   string
	DatabaseName string
}

var structNameVarTemplaterFragment = `{{$StructName := print (UppercaseBeginning .DatabaseName) (UppercaseBeginning .EntityName) "Repository" -}}`

var structDefinition = `type {{UppercaseBeginning .DatabaseName}}{{UppercaseBeginning .EntityName}}Repository struct {
    db sql.Db
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

			tmpl, err = tmpl.New(tools.LowercaseBeginning(entity.EntityName) + "_repository").Funcs(tools.FullFuncMap).Parse(structNameVarTemplaterFragment + string(tmplRaw))
			if err != nil {
				panic(err)
			}

			outFile, err := os.Create("repository/" + database.DatabaseName + "/" + tools.LowercaseBeginning(entity.EntityName) + "_repository.go")
			if err != nil {
				panic(err)
			}

			err = tmpl.Execute(outFile, Configuration{
				EntityName:   entity.EntityName,
				DatabaseName: database.DatabaseName,
			})
			if err != nil {
				panic(err)
			}
		}
	}
}
