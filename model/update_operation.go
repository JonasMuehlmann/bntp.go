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

package model

import (
	"bytes"
	"encoding/json"
)

type UpdateOperator int

const (
	UpdateSet UpdateOperator = iota + 1
	UpdateClear

	UpdateAppend
	UpdatePrepend

	UpdateAdd
	UpdateSubtract
)

func (o UpdateOperator) String() string {
	switch o {
	case UpdateSet:
		return "UpdateSet"
	case UpdateClear:
		return "UpdateClear"
	case UpdateAppend:
		return "UpdateAppend"
	case UpdatePrepend:
		return "UpdatePrepend"
	case UpdateAdd:
		return "UpdateAdd"
	case UpdateSubtract:
		return "UpdateSubtract"
	default:
		return ""
	}
}

func UpdateOperatorFromString(s string) UpdateOperator {
	switch s {
	case "UpdateSet":
		return UpdateSet
	case "UpdateClear":
		return UpdateClear
	case "UpdateAppend":
		return UpdateAppend
	case "UpdatePrepend":
		return UpdatePrepend
	case "UpdateAdd":
		return UpdateAdd
	case "UpdateSubtract":
		return UpdateSubtract
	default:
		return 0
	}
}

func (s UpdateOperator) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString("")
	if s != 0 {
		buffer.WriteString(`"`)
		buffer.WriteString(s.String())
		buffer.WriteString(`"`)
	}

	return buffer.Bytes(), nil
}

func (s *UpdateOperator) UnmarshalJSON(b []byte) error {
	var j string

	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}

	*s = UpdateOperatorFromString(j)

	return nil
}

type UpdateOperation[T any] struct {
	Operand  T              `json:"operand" toml:"operand" yaml:"operand"`
	Operator UpdateOperator `json:"operator" toml:"operator" yaml:"operator"`
}

func ApplyUpdater[T any](target *T, updateOperation UpdateOperation[T]) {
	var zero T

	switch updateOperation.Operator {
	case UpdateSet:
		*target = updateOperation.Operand
	case UpdateClear:
		*target = zero
	case UpdateAppend:
		switch t := any(target).(type) {
		case *string:
			operand, _ := any(updateOperation.Operand).(string)
			*t += operand
		default:
			panic("Can only append to string")
		}
		// TODO: Support slices
	case UpdatePrepend:
		switch t := any(target).(type) {
		case *string:
			operand, _ := any(updateOperation.Operand).(string)
			*t = operand + *t
		default:
			panic("Can only prepend to string")
		}
	case UpdateAdd:
		switch t := any(target).(type) {
		case *int:
			operand, _ := any(updateOperation.Operand).(int)
			*t += operand
		case *int8:
			operand, _ := any(updateOperation.Operand).(int8)
			*t += operand
		case *int16:
			operand, _ := any(updateOperation.Operand).(int16)
			*t += operand
		case *int32:
			operand, _ := any(updateOperation.Operand).(int32)
			*t += operand
		case *int64:
			operand, _ := any(updateOperation.Operand).(int64)
			*t += operand
		case *uint:
			operand, _ := any(updateOperation.Operand).(uint)
			*t += operand
		case *uint8:
			operand, _ := any(updateOperation.Operand).(uint8)
			*t += operand
		case *uint16:
			operand, _ := any(updateOperation.Operand).(uint16)
			*t += operand
		case *uint32:
			operand, _ := any(updateOperation.Operand).(uint32)
			*t += operand
		case *uint64:
			operand, _ := any(updateOperation.Operand).(uint64)
			*t += operand
		case *float32:
			operand, _ := any(updateOperation.Operand).(float32)
			*t += operand
		case *float64:
			operand, _ := any(updateOperation.Operand).(float64)
			*t += operand
		default:
			panic("Can only add to numeric types")
		}
	case UpdateSubtract:
		switch t := any(target).(type) {
		case *int:
			operand, _ := any(updateOperation.Operand).(int)
			*t -= operand
		case *int8:
			operand, _ := any(updateOperation.Operand).(int8)
			*t -= operand
		case *int16:
			operand, _ := any(updateOperation.Operand).(int16)
			*t -= operand
		case *int32:
			operand, _ := any(updateOperation.Operand).(int32)
			*t -= operand
		case *int64:
			operand, _ := any(updateOperation.Operand).(int64)
			*t -= operand
		case *uint:
			operand, _ := any(updateOperation.Operand).(uint)
			*t -= operand
		case *uint8:
			operand, _ := any(updateOperation.Operand).(uint8)
			*t -= operand
		case *uint16:
			operand, _ := any(updateOperation.Operand).(uint16)
			*t -= operand
		case *uint32:
			operand, _ := any(updateOperation.Operand).(uint32)
			*t -= operand
		case *uint64:
			operand, _ := any(updateOperation.Operand).(uint64)
			*t -= operand
		case *float32:
			operand, _ := any(updateOperation.Operand).(float32)
			*t -= operand
		case *float64:
			operand, _ := any(updateOperation.Operand).(float64)
			*t -= operand
		default:
			panic("Can only add to numeric types")
		}
	}
}
