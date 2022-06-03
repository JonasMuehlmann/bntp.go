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

package repository

import (
	"database/sql"
	"time"

	"github.com/JonasMuehlmann/optional.go"
	"github.com/volatiletech/null/v8"
)

func TimeToStr(time time.Time) (string, error) {
	return time.Format("2006-01-02 15:04:05"), nil
}

func StrToTime(time_ string) (time.Time, error) {
	return time.Parse("2006-01-02 15:04:05", time_)
}

func OptionalTimeToNullStr(time optional.Optional[time.Time]) (null.String, error) {
	if time.HasValue {
		return null.StringFrom(time.Wrappee.Format("2006-01-02 15:04:05")), nil
	}

	return null.String{}, nil
}

func NullStrToOptionalTime(time_ null.String) (optional.Optional[time.Time], error) {
	if time_.Valid {
		t, err := time.Parse("2006-01-02 15:04:05", time_.String)

		return optional.Make(t), err
	}

	return optional.Optional[time.Time]{}, nil
}

func OptionalTimeToNullTime(time optional.Optional[time.Time]) (null.Time, error) {
	if time.HasValue {
		return null.TimeFrom(time.Wrappee), nil
	}

	return null.Time{}, nil
}

func NullTimeToOptionalTime(time_ null.Time) (optional.Optional[time.Time], error) {
	if time_.Valid {
		return optional.Make(time_.Time), nil
	}

	return optional.Optional[time.Time]{}, nil
}

func OptionalStringToNullString(str optional.Optional[string]) (null.String, error) {
	if str.HasValue {
		return null.StringFrom(str.Wrappee), nil
	}

	return null.String{}, nil
}

func NullStringToOptionalString(str null.String) (optional.Optional[string], error) {
	if str.Valid {
		return optional.Make(str.String), nil
	}

	return optional.Optional[string]{}, nil
}

func BoolToInt(b bool) (int64, error) {
	if b {
		return 1, nil
	}

	return 0, nil
}

func IntToBool(i int64) (bool, error) {
	return i > 0, nil
}

func MakeDomainToRepositoryEntityConverter[TIn any, TOut any](db *sql.DB, converter func(db *sql.DB, entity *TIn) (*TOut, error)) func(entity *TIn) (*TOut, error) {
	return func(entity *TIn) (*TOut, error) {
		return converter(db, entity)
	}
}
