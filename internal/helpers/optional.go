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

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

// Optional holds a Wrappe and a flag indicating the existence or abscence of the Wrapee.
type Optional[T any] struct {
	Wrappee  T    `json:"wrapee" db:"wrapee"`
	HasValue bool `json:"has_value" db:"has_value"`
}

// ValueOr returns the Wrapee if the HasValue flag is true, otherwise alternative is returned.
func (optional Optional[T]) ValueOr(alternative T) T {
	if optional.HasValue {
		return optional.Wrappee
	}

	return alternative
}

// OrElse returns the Wrapee if the HasValue flag is true, otherwise the result of alternative() is returned.
func (optional Optional[T]) OrElse(alternative func() T) T {
	if optional.HasValue {
		return optional.Wrappee
	}

	return alternative()
}

// Push sets a Wrappee to val value and the HasValue flag to true.
func (optional *Optional[T]) Push(val T) {
	optional.Wrappee = val
	optional.HasValue = true
}

// Pop sets the HasValue flag to false and returns the Wrappee.
func (optional *Optional[T]) Pop() T {
	optional.HasValue = false

	return optional.Wrappee
}

func (optional *Optional[T]) UnmarshalJSON(input []byte) error {
	// Try to parse key with value of equal type as "Wrapee"
	err := json.Unmarshal(input, &optional.Wrappee)

	if err != nil {
		alias := struct {
			Wrappee  T    `json:"wrapee" db:"wrapee"`
			HasValue bool `json:"has_value" db:"has_value"`
		}{}
		// Try to parse object with keys "wrapee" and "has_value"
		err = json.Unmarshal(input, &alias)
		if err != nil {
			optional.HasValue = false

			return err
		}

		optional.HasValue = true
		optional.Wrappee = alias.Wrappee

		return nil
	}

	optional.HasValue = true

	return nil

}

// Scan implements the database.sql.Scanner interface
func (optional *Optional[T]) Scan(value any) error {
	if value == nil {
		var zero T

		optional.HasValue = false
		optional.Wrappee = zero

		return nil
	}

	switch t := value.(type) {
	case T:
		optional.HasValue = true
		optional.Wrappee = t

		return nil
	}

	return fmt.Errorf("Failed to scan value of type %T into optional of type %T", value, optional.Wrappee)
}

// Value implements the database.sql.driver.Valuer interface
func (optional Optional[T]) Value() (driver.Value, error) {
	if !optional.HasValue {
		return nil, nil
	}

	if driver.IsValue(optional.Wrappee) {
		return optional.Wrappee, nil
	}

	return nil, nil
}

// String implements the fmt.Stringer interface
func (optional Optional[T]) String() string {
	if optional.HasValue {
		return fmt.Sprint(optional.Wrappee)
	}

	return "empty optional"
}
