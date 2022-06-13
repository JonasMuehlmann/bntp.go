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

package tools

import (
	"reflect"
	"strings"
	"text/template"
)

func UppercaseBeginning(str string) string {
	return strings.ToUpper(str[:1]) + str[1:]
}

func LowercaseBeginning(str string) string {
	return strings.ToLower(str[:1]) + str[1:]
}

func Pluralize(str string) string {
	return str + "s"
}

func Unslice(str string) string {
	return strings.TrimPrefix(str, "[]")
}

func UnaliasSQLBoilerSlice(str string) string {
	switch str {
	case "TagSlice":
		return "[]*Tag"
	case "DocumentSlice":
		return "[]*Document"
	case "BookmarkSlice":
		return "[]*Bookmark"
	}

	return str
}

type StructField struct {
	FieldName        string
	FieldType        string
	FieldTags        string
	LogicalFieldName string
}
type Struct struct {
	StructName   string
	StructFields []StructField
}

func NewStructModel(target any) Struct {
	entityStruct := Struct{}

	entityStructType := reflect.TypeOf(target)
	entityStruct.StructName = entityStructType.Name()

	for i := 0; i < entityStructType.NumField(); i++ {
		field := entityStructType.Field(i)

		fieldType := field.Type.String()
		fieldType = strings.Replace(fieldType, "main.", "", 1)
		fieldType = strings.Replace(fieldType, "repository.", "", 1)

		entityStruct.StructFields = append(entityStruct.StructFields, StructField{
			FieldName:        field.Name,
			FieldType:        fieldType,
			FieldTags:        "`" + string(field.Tag) + "`",
			LogicalFieldName: strings.Replace(field.Tag.Get("json"), ",omitempty", "", 1),
		})
	}

	return entityStruct
}

var FullFuncMap = template.FuncMap{
	"UppercaseBeginning":    UppercaseBeginning,
	"LowercaseBeginning":    LowercaseBeginning,
	"Pluralize":             Pluralize,
	"Unslice":               Unslice,
	"UnaliasSQLBoilerSlice": UnaliasSQLBoilerSlice,
}
