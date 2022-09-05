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

type FilterOperator int

// TODO: Implement FilterContains and FilterNotContains
// TODO: Implement FilterEmpty and FilterNotEmpty
const (
	FilterEqual FilterOperator = iota + 1
	FilterNEqual
	FilterGreaterThan
	FilterGreaterThanEqual
	FilterLessThan
	FilterLessThanEqual

	FilterIn
	FilterNotIn

	FilterBetween
	FilterNotBetween

	FilterLike
	FilterNotLike

	FilterOr
	FilterAnd
)

func (o FilterOperator) String() string {
	switch o {
	case FilterEqual:
		return "FilterEqual"
	case FilterNEqual:
		return "FilterNEqual"
	case FilterGreaterThan:
		return "FilterGreaterThan"
	case FilterGreaterThanEqual:
		return "FilterGreaterThanEqual"
	case FilterLessThan:
		return "FilterLessThan"
	case FilterLessThanEqual:
		return "FilterLessThanEqual"
	case FilterIn:
		return "FilterIn"
	case FilterNotIn:
		return "FilterNotIn"
	case FilterBetween:
		return "FilterNotIn"
	case FilterNotBetween:
		return "FilterBetweenEqual"
	case FilterLike:
		return "FilterLike"
	case FilterNotLike:
		return "FilterNotLike"
	case FilterOr:
		return "FilterOr"
	case FilterAnd:
		return "FilterAnd"
	default:
		return ""
	}
}

func FilterOperatorFromString(s string) FilterOperator {
	switch s {
	case "FilterEqual":
		return FilterEqual
	case "FilterNEqual":
		return FilterNEqual
	case "FilterGreaterThan":
		return FilterGreaterThan
	case "FilterGreaterThanEqual":
		return FilterGreaterThanEqual
	case "FilterLessThan":
		return FilterLessThan
	case "FilterLessThanEqual":
		return FilterLessThanEqual
	case "FilterIn":
		return FilterIn
	case "FilterNotIn":
		return FilterNotIn
	case "FilterBetween":
		return FilterBetween
	case "FilterNotBetween":
		return FilterNotBetween
	case "FilterLike":
		return FilterLike
	case "FilterNotLike":
		return FilterNotLike
	case "FilterOr":
		return FilterOr
	case "FilterAnd":
		return FilterAnd
	default:
		return 0
	}
}

func (s FilterOperator) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(s.String())
	buffer.WriteString(`"`)

	return buffer.Bytes(), nil
}

func (s *FilterOperator) UnmarshalJSON(b []byte) error {
	var j string

	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}

	*s = FilterOperatorFromString(j)

	return nil
}

type ScalarOperand[T any] struct {
	Operand T `json:"operand" toml:"operand" yaml:"operand"`
}

type RangeOperand[T any] struct {
	Start T `json:"start" toml:"start" yaml:"start"`
	End   T `json:"end" toml:"end" yaml:"end"`
}

type ListOperand[T any] struct {
	Operands []T `json:"operands" toml:"operands" yaml:"operands"`
}

type CompoundOperand[T any] struct {
	LHS FilterOperation[T] `json:"lhs" toml:"lhs" yaml:"lhs"`
	RHS FilterOperation[T] `json:"rhs" toml:"rhs" yaml:"rhs"`
}

type Operand[T any] interface {
	OperandDummyMethod()
}

// TODO: Find a way to have such marshallers but allow users to support their own through plugins

func (s *FilterOperation[T]) UnmarshalJSON(b []byte) error {
	tmpOperation := map[string]json.RawMessage{}
	tmpOperand := map[string]json.RawMessage{}

	err := json.Unmarshal(b, &tmpOperation)
	if err != nil {
		return err
	}

	err = json.Unmarshal(tmpOperation["operator"], &s.Operator)
	if err != nil {
		return err
	}
	err = json.Unmarshal(tmpOperation["operand"], &tmpOperand)
	if err != nil {
		return err
	}

	if _, ok := tmpOperand["operand"]; ok {
		operand := ScalarOperand[T]{}

		err = json.Unmarshal(tmpOperation["operand"], &operand)
		if err != nil {
			return err
		}

		s.Operand = operand

	} else if _, ok := tmpOperand["start"]; ok {
		operand := RangeOperand[T]{}
		err = json.Unmarshal(tmpOperation["operand"], &operand)
		if err != nil {
			return err
		}

		s.Operand = operand
	} else if _, ok := tmpOperand["operands"]; ok {
		operand := ListOperand[T]{}
		err = json.Unmarshal(tmpOperation["operand"], &operand)
		if err != nil {
			return err
		}

		s.Operand = operand
	} else if _, ok := tmpOperand["lhs"]; ok {
		operand := CompoundOperand[T]{}
		err = json.Unmarshal(tmpOperation["operand"], &operand)
		if err != nil {
			return err
		}

		s.Operand = operand
	}

	return nil
}

func (o ScalarOperand[T]) OperandDummyMethod()   {}
func (o RangeOperand[T]) OperandDummyMethod()    {}
func (o ListOperand[T]) OperandDummyMethod()     {}
func (o CompoundOperand[T]) OperandDummyMethod() {}

type FilterOperation[T any] struct {
	Operand  Operand[T]     `json:"operand" toml:"operand" yaml:"operand"`
	Operator FilterOperator `json:"operator" toml:"operator" yaml:"operator"`
}

func ConvertFilter[TOut any, TIn any](from FilterOperation[TIn], operandConverter func(fromOperator TIn) (TOut, error)) (FilterOperation[TOut], error) {
	switch t := any(from.Operand).(type) {
	//***********************    ScalarOperand    **********************//
	case ScalarOperand[TIn]:
		operandValue, err := operandConverter(t.Operand)
		if err != nil {
			return FilterOperation[TOut]{}, err
		}

		return FilterOperation[TOut]{Operand: ScalarOperand[TOut]{Operand: operandValue}, Operator: from.Operator}, nil
		//***********************    RangeOperand    ***********************//
	case RangeOperand[TIn]:
		operandStartValue, err := operandConverter(t.Start)
		if err != nil {
			return FilterOperation[TOut]{}, err
		}
		operandStopValue, err := operandConverter(t.End)
		if err != nil {
			return FilterOperation[TOut]{}, err
		}

		return FilterOperation[TOut]{Operand: RangeOperand[TOut]{Start: operandStartValue, End: operandStopValue}, Operator: from.Operator}, nil
		//************************    ListOperand    ***********************//
	case ListOperand[TIn]:
		operandValues := make([]TOut, 0, len(t.Operands))
		for _, operandValue := range t.Operands {
			newOperandValue, err := operandConverter(operandValue)
			if err != nil {
				return FilterOperation[TOut]{}, err
			}

			operandValues = append(operandValues, newOperandValue)
		}

		return FilterOperation[TOut]{Operand: ListOperand[TOut]{Operands: operandValues}, Operator: from.Operator}, nil
		//**********************    CompoundOperand    *********************//
	case CompoundOperand[TIn]:
		lhsOperation, err := ConvertFilter(t.LHS, operandConverter)
		if err != nil {
			return FilterOperation[TOut]{}, err
		}

		rhsOperation, err := ConvertFilter(t.RHS, operandConverter)
		if err != nil {
			return FilterOperation[TOut]{}, err
		}

		return FilterOperation[TOut]{
			Operand:  CompoundOperand[TOut]{LHS: lhsOperation, RHS: rhsOperation},
			Operator: from.Operator,
		}, nil
	default:
		panic("Unhandled Operand type")
	}
}
