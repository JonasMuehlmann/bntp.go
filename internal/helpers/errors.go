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

package helpers

import "fmt"

// InvalidIdentifierError indicates that an identifier is not one of the allowed types.
type InvalidIdentifierError struct {
	BadIdentifier any
}

func (err InvalidIdentifierError) Error() string {
	return fmt.Sprintf("Identifier of type %t is invalid, should be string or int", err.BadIdentifier)
}

// DeserializationError indicates that a deserialization failed because of the reason Inner.
// Fitting conditions for this error are, for example, when input data had an invalid structure.
type DeserializationError struct {
	Inner error
}

func (err DeserializationError) Error() string {
	return fmt.Sprintf("Failed deserialization: %v", err.Inner)
}

// SerializationError indicates that a serialization failed because of the reason Inner.
// Fitting conditions for this error are, for example, when data to output was incomplete.
type SerializationError struct {
	Inner error
}

func (err SerializationError) Error() string {
	return fmt.Sprintf("Failed serialization: %v", err.Inner)
}

// IneffectiveOperationError indicates that an operation had no effect because of the reason Inner.
// Fitting conditions for this error are, for example, when an empty dataset is to be processed.
type IneffectiveOperationError struct {
	Inner error
}

func (err IneffectiveOperationError) Error() string {
	return fmt.Sprintf("Operation had no effect: %v", err.Inner)
}

// ImportError that an import failed because of the reason Inner.
// Fitting conditions for this error are, for example, when input data was empty.
type ImportError struct {
	Inner error
}

func (err ImportError) Error() string {
	return fmt.Sprintf("Failed import: %v", err.Inner)
}

// ExportError that an export failed because of the reason Inner.
// Fitting conditions for this error are, for example, when the data provider failed to return a dataset.
type ExportError struct {
	Inner error
}

func (err ExportError) Error() string {
	return fmt.Sprintf("Failed export: %v", err.Inner)
}
