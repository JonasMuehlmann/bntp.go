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
	"io/ioutil"
	"log"
	"path"
	"regexp"
	"strings"
)

const (
	MainSchema = "sqlite"
	SchemaDir  = "schema"
	FileName   = "bntp"
)

const (
	FORMAT_MYSQL          = "mysql"
	FORMAT_POSTGRES       = "postgres"
	FORMAT_SQLSERVER      = "sqlserver"
	FORMAT_SQLSERVER_TEST = "sqlserver_test"
)

type Converter func(string) string

var (
	Formats    = []string{FORMAT_MYSQL, FORMAT_POSTGRES, FORMAT_SQLSERVER}
	Converters = map[string]Converter{
		FORMAT_MYSQL:          ToMysql,
		FORMAT_POSTGRES:       ToPostgrees,
		FORMAT_SQLSERVER:      ToSqlServer,
		FORMAT_SQLSERVER_TEST: ToSqlServerTest,
	}
)

func GetFileName(format string) string {
	return FileName + "_" + format + ".sql"
}

func main() {
	mainSchemaContentRaw, err := ioutil.ReadFile(path.Join(SchemaDir, GetFileName(MainSchema)))
	mainSchemaContent := string(mainSchemaContentRaw)
	if err != nil {
		log.Fatal(err)
	}

	for format, converter := range Converters {
		newSchema := converter(mainSchemaContent)

		newSchemaFileName := path.Join(SchemaDir, GetFileName(format))

		err := ioutil.WriteFile(newSchemaFileName, []byte(newSchema), 0o644)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func ToMysql(mainSchema string) string {
	schema := mainSchema

	schema = strings.Replace(schema, "PRAGMA foreign_keys = ON;", "", -1)
	schema = strings.Replace(schema, "PRIMARY KEY NOT NULL", "PRIMARY KEY", -1)
	schema = strings.ReplaceAll(schema, "TEXT", "VARCHAR(255)")
	schema = strings.ReplaceAll(schema, "INTEGER", "BIGINT")

	return schema
}

func ToPostgrees(mainSchema string) string {
	schema := mainSchema

	schema = strings.ReplaceAll(schema, "PRAGMA foreign_keys = ON;", "")
	schema = strings.ReplaceAll(schema, "INTEGER", "BIGINT")

	return schema
}

func ToSqlServer(mainSchema string) string {
	schema := mainSchema

	schema = strings.Replace(schema, "PRAGMA foreign_keys = ON;", "", -1)
	schema = strings.Replace(schema, "PRIMARY KEY NOT NULL", "PRIMARY KEY", -1)
	schema = strings.ReplaceAll(schema, "TEXT", "VARCHAR(255)")
	schema = strings.ReplaceAll(schema, "TIMESTAMP", "DATETIME")
	schema = strings.ReplaceAll(schema, "INTEGER", "BIGINT")

	return schema
}

func ToSqlServerTest(mainSchema string) string {
	schema := mainSchema

	schema = strings.Replace(schema, "PRAGMA foreign_keys = ON;", "", -1)
	schema = strings.Replace(schema, "PRIMARY KEY NOT NULL", "PRIMARY KEY", -1)
	schema = strings.ReplaceAll(schema, "TEXT", "VARCHAR(255)")
	schema = strings.ReplaceAll(schema, "TIMESTAMP", "DATETIME")
	removeForeignKeys := regexp.MustCompile(`REFERENCES\s+\w+\(\w+\)`)
	schema = removeForeignKeys.ReplaceAllString(schema, "")

	return schema
}
