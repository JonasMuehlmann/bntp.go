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

package marshallers

import (
	"encoding/json"
	"strings"

	"github.com/ghodss/yaml"
	"github.com/gocarina/gocsv"
)

// ******************************************************************//
//                            Interfaces                            //
// ******************************************************************//.
type Marshaller interface {
	Marshall(from any) (to string, err error)
}

type Unmarshaller interface {
	Unmarshall(out any, in string) error
}

// ******************************************************************//
//                               Json                               //
// ******************************************************************//.
type JsonMarshaller struct{}

func (marshaller *JsonMarshaller) Marshall(from any) (to string, err error) {
	res, err := json.Marshal(from)

	return string(res), err
}

type JsonUnmarshaller struct{}

func (unmarshaller *JsonUnmarshaller) Unmarshall(out any, in string) error {
	return json.Unmarshal([]byte(in), out)
}

// ******************************************************************//
//                               Yaml                               //
// ******************************************************************//.
type YamlMarshaller struct{}

func (marshaller *YamlMarshaller) Marshall(from any) (to string, err error) {
	res, err := yaml.Marshal(from)

	return string(res), err
}

type YamlUnmarshaller struct{}

func (unmarshaller *YamlUnmarshaller) Unmarshall(out any, in string) error {
	return yaml.Unmarshal([]byte(in), out)
}

// ******************************************************************//
//                               Csv                                //
// ******************************************************************//.
type CsvMarshaller struct{}

func (marshaller *CsvMarshaller) Marshall(from any) (to string, err error) {
	res := new(strings.Builder)
	err = gocsv.Marshal(from, res)

	return res.String(), err
}

type CsvUnmarshaller struct{}

func (unmarshaller *CsvUnmarshaller) Unmarshall(out any, in string) error {
	return gocsv.Unmarshal(strings.NewReader(in), out)
}
