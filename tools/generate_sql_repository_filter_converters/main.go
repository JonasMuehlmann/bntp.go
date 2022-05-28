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

type Entities struct {
	Document string
	Bookmark string
	Tag      string
}

var entities = Entities{
	Document: "Document",
	Bookmark: "Bookmark",
	Tag:      "Tag",
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
	Entities     Entities
	DatabaseName string
}

func main() {
	for _, database := range databases {
		tmplRaw, err := os.ReadFile("templates/sql_repositories/filter_converter.go.tpl")
		if err != nil {
			panic(err)
		}

		tmpl, err := template.New(tools.LowercaseBeginning(database.DatabaseName) + "_filter_converter").Funcs(tools.FullFuncMap).Parse(string(tmplRaw))
		if err != nil {
			panic(err)
		}

		outFile, err := os.Create("model/repository/" + database.DatabaseName + "/filter_converter.go")
		if err != nil {
			panic(err)
		}

		err = tmpl.Execute(outFile, Configuration{
			Entities:     entities,
			DatabaseName: database.DatabaseName,
		})
		if err != nil {
			panic(err)
		}
	}
}
