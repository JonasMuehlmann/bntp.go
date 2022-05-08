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

package domain

type FilterOperator int

const (
	FilterEqual FilterOperator = iota
	FilterNEqual
	FilterGreaterThan
	FilterGreaterThanEqual
	FilterLessThan
	FilterLessThanEqual

	FilterIn
	FilterWhereNotIn

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
	Start ScalarOperand[T]
	End   ScalarOperand[T]
}

type ListOperand[T any] struct {
	Operands []ScalarOperand[T]
}

type Operand[T any] interface {
	OperandDummyMethod()
}

func (o ScalarOperand[T]) OperandDummyMethod() {}
func (o RangeOperand[T]) OperandDummyMethod()  {}
func (o ListOperand[T]) OperandDummyMethod()   {}

type FilterOperation[T any] struct {
	Operator FilterOperator
	// TODO: Filter by fields type
	Operand Operand[T]
}
