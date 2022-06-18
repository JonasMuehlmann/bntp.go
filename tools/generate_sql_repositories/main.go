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

	mssqlRepository "github.com/JonasMuehlmann/bntp.go/model/repository/mssql"
	mysqlRepository "github.com/JonasMuehlmann/bntp.go/model/repository/mysql"
	psqlRepository "github.com/JonasMuehlmann/bntp.go/model/repository/psql"
	sqlite3Repository "github.com/JonasMuehlmann/bntp.go/model/repository/sqlite3"

	"github.com/JonasMuehlmann/bntp.go/tools"
	"github.com/JonasMuehlmann/goaoi"
)

type Entity struct {
	EntityName string
	Struct     any
}

var dbEntities = map[string][]Entity{
	"sqlite3": {
		{"Document", sqlite3Repository.Document{}},
		{"Bookmark", sqlite3Repository.Bookmark{}},
		{"Tag", sqlite3Repository.Tag{}},
	},
	"mssql": {
		{"Document", mssqlRepository.Document{}},
		{"Bookmark", mssqlRepository.Bookmark{}},
		{"Tag", mssqlRepository.Tag{}},
	},
	// NOTE: Currently broken on sqlboiler side
	// "mysql": {
	// 	{"Document", mysqlRepository.Document{}},
	// 	{"Bookmark", mysqlRepository.Bookmark{}},
	// 	{"Tag", mysqlRepository.Tag{}},
	// },
	"psql": {
		{"Document", psqlRepository.Document{}},
		{"Bookmark", psqlRepository.Bookmark{}},
		{"Tag", psqlRepository.Tag{}},
	},
}

type Database struct {
	DatabaseName string
}
type Configuration struct {
	EntityName     string
	DatabaseName   string
	StructFields   []tools.StructField
	RelationFields []tools.StructField
}

func main() {
	for database, entities := range dbEntities {
		for _, entity := range entities {
			tmplRaw, err := os.ReadFile("templates/sql_repositories/sql_repository.go.tpl")
			if err != nil {
				panic(err)
			}

			tmpl, err := template.New("structDefinition").Funcs(tools.FullFuncMap).Parse(string(tmplRaw))
			if err != nil {
				panic(err)
			}

			outFile, err := os.Create("model/repository/" + database + "/" + tools.LowercaseBeginning(entity.EntityName) + "_repository.go")
			if err != nil {
				panic(err)
			}

			entityStruct := tools.NewStructModel(entity.Struct)

			entityStruct.StructFields, err = goaoi.CopyExceptIfSlice(entityStruct.StructFields, func(s tools.StructField) bool { return s.FieldName == "R" || s.FieldName == "L" })
			if err != nil {
				panic(err)
			}

			var relationStruct tools.Struct

			// **************************    sqlite3    *************************//
			switch e := entity.Struct.(type) {
			case sqlite3Repository.Bookmark:
				e.R = e.R.NewStruct()
				relationStruct = tools.NewStructModel(*e.R)
			case sqlite3Repository.Document:
				e.R = e.R.NewStruct()
				relationStruct = tools.NewStructModel(*e.R)
			case sqlite3Repository.Tag:
				e.R = e.R.NewStruct()
				relationStruct = tools.NewStructModel(*e.R)

				// ***************************    mssql    **************************//
			case mssqlRepository.Bookmark:
				e.R = e.R.NewStruct()
				relationStruct = tools.NewStructModel(*e.R)
			case mssqlRepository.Document:
				e.R = e.R.NewStruct()
				relationStruct = tools.NewStructModel(*e.R)
			case mssqlRepository.Tag:
				e.R = e.R.NewStruct()
				relationStruct = tools.NewStructModel(*e.R)

				// ***************************    mysql    **************************//
			case mysqlRepository.Bookmark:
				e.R = e.R.NewStruct()
				relationStruct = tools.NewStructModel(*e.R)
			case mysqlRepository.Document:
				e.R = e.R.NewStruct()
				relationStruct = tools.NewStructModel(*e.R)
			case mysqlRepository.Tag:
				e.R = e.R.NewStruct()
				relationStruct = tools.NewStructModel(*e.R)

				// ***************************    psql    ***************************//
			case psqlRepository.Bookmark:
				e.R = e.R.NewStruct()
				relationStruct = tools.NewStructModel(*e.R)
			case psqlRepository.Document:
				e.R = e.R.NewStruct()
				relationStruct = tools.NewStructModel(*e.R)
			case psqlRepository.Tag:
				e.R = e.R.NewStruct()
				relationStruct = tools.NewStructModel(*e.R)

			default:
				panic("Unhandled sql repository type")
			}

			err = tmpl.Execute(outFile, Configuration{
				EntityName:     entity.EntityName,
				DatabaseName:   database,
				StructFields:   entityStruct.StructFields,
				RelationFields: relationStruct.StructFields,
			})
			if err != nil {
				panic(err)
			}
		}
	}
}
