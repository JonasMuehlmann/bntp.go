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
	"strings"
	"text/template"
)

type Entity struct {
	EntityName string
}

var entities = []Entity{
	{"Document"},
	{"Bookmark"},
	{"Tag"},
}

func main() {
	tmplRaw, err := os.ReadFile("templates/repository.go.tpl")
	if err != nil {
		panic(err)
	}

	tmpl, err := template.New("repository").Funcs(template.FuncMap{
		"UppercaseBeginning": UppercaseBeginning,
		"LowercaseBeginning": LowercaseBeginning,
		"Pluralize":          Pluralize,
	}).Parse(string(tmplRaw))
	if err != nil {
		panic(err)
	}

	for _, entity := range entities {
		outFile, err := os.Create("repository/" + LowercaseBeginning(entity.EntityName) + "_repository.go")
		if err != nil {
			panic(err)
		}

		err = tmpl.Execute(outFile, entity)
		if err != nil {
			panic(err)
		}
	}
}

func UppercaseBeginning(str string) string {
	return strings.ToUpper(str[:1]) + str[1:]
}

func LowercaseBeginning(str string) string {
	return strings.ToLower(str[:1]) + str[1:]
}

func Pluralize(str string) string {
	return str + "s"
}
