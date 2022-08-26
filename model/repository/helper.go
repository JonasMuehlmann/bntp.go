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
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/JonasMuehlmann/bntp.go/internal/helper"
	"github.com/JonasMuehlmann/optional.go"
	"github.com/volatiletech/null/v8"
)

func TimeToStr(time time.Time) (string, error) {
	return time.Format(helper.DateFormat), nil
}

func StrToTime(time_ string) (time.Time, error) {
	return time.Parse(helper.DateFormat, time_)
}

func OptionalTimeToNullStr(time optional.Optional[time.Time]) (null.String, error) {
	if time.HasValue {
		return null.StringFrom(time.Wrappee.Format(helper.DateFormat)), nil
	}

	return null.String{}, nil
}

func NullStrToOptionalTime(time_ null.String) (optional.Optional[time.Time], error) {
	if time_.Valid {
		t, err := time.Parse(helper.DateFormat, time_.String)

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

func MakeDomainToRepositoryEntityConverter[TIn any, TOut any](ctx context.Context, converter func(ctx context.Context, entity *TIn) (*TOut, error)) func(entity *TIn) (*TOut, error) {
	return func(entity *TIn) (*TOut, error) {
		return converter(ctx, entity)
	}
}

func MakeDomainToRepositoryEntityConverterGeneric[TIn any, TOut any](ctx context.Context, converter func(ctx context.Context, entity *TIn) (any, error)) func(entity *TIn) (*TOut, error) {
	return func(entity *TIn) (*TOut, error) {
		entityRaw, err := converter(ctx, entity)
		if err != nil {
			return nil, err
		}

		entityConcrete, ok := entityRaw.(*TOut)
		if !ok {
			var wantedZeroVal *TOut
			return nil, fmt.Errorf("expected type %T but got %T", wantedZeroVal, entityConcrete)
		}

		return entityConcrete, nil
	}
}

//******************************************************************//
//               ReferenceToNonExistentDependencyError              //
//******************************************************************//

type ReferenceToNonExistentDependencyError struct {
	Inner error
}

func (err ReferenceToNonExistentDependencyError) Error() string {
	return fmt.Sprintf("Error when writing data with reference to non-existent dependency: %v", err.Inner)
}

func (err ReferenceToNonExistentDependencyError) Unwrap() error {
	return err.Inner
}

func (err ReferenceToNonExistentDependencyError) Is(other error) bool {
	switch other.(type) {
	case ReferenceToNonExistentDependencyError:
		return true
	default:
		return false
	}
}

func (err ReferenceToNonExistentDependencyError) As(target any) bool {
	switch target.(type) {
	case ReferenceToNonExistentDependencyError:
		reflect.Indirect(reflect.ValueOf(target)).Set(reflect.ValueOf(err))
		return true
	default:
		return false
	}

}
