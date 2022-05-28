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

type FilterOperator int

const (
	FilterEqual FilterOperator = iota
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

type ScalarOperand[T any] struct {
	Operand T
}

type RangeOperand[T any] struct {
	Start T
	End   T
}

type ListOperand[T any] struct {
	Operands []T
}

type CompoundOperand[T any] struct {
	LHS FilterOperation[T]
	RHS FilterOperation[T]
}

type Operand[T any] interface {
	OperandDummyMethod()
}

func (o ScalarOperand[T]) OperandDummyMethod()   {}
func (o RangeOperand[T]) OperandDummyMethod()    {}
func (o ListOperand[T]) OperandDummyMethod()     {}
func (o CompoundOperand[T]) OperandDummyMethod() {}

type FilterOperation[T any] struct {
	Operand  Operand[T]
	Operator FilterOperator
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
