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

package subcommands

import (
	"fmt"
)

type ParameterConversionError struct {
	Parameter  any
	Value      any
	TargetType string
}

func (err ParameterConversionError) Error() string {
	return fmt.Sprintf("Failed to convert parameter %v of value %v to type %v", err.Parameter, err.Value, err.TargetType)
}

type IncompleteCompoundParameterError struct {
	Parameter     any
	MissingFields []string
}

func (err IncompleteCompoundParameterError) Error() string {
	return fmt.Sprintf("Compund parameter %v is missing value for field(s) %q", err.Parameter, err.MissingFields)
}

type InvalidParameterValueError struct {
	Parameter     any
	BadValue      any
	AllowedValues []any
}

func (err InvalidParameterValueError) Error() string {
	return fmt.Sprintf("Received invalid value %v in parameter %v, allowed values are %p", err.BadValue, err.Parameter, err.AllowedValues)
}
