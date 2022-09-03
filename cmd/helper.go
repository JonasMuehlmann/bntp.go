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

package cmd

import (
	"fmt"
	"reflect"
)

func UnmarshalEntities[TEntity any](cli *Cli, args []string, format string) (entities []*TEntity, err error) {
	tags := make([]*TEntity, len(args))
	for i, arg := range args {
		tags[i] = new(TEntity)

		err := cli.BNTPBackend.Unmarshallers[format].Unmarshall(tags[i], arg)
		if err != nil {
			return tags, EntityMarshallingError{Inner: err}
		}
	}

	return tags, nil
}

//******************************************************************//
//                      EntitymarshallingError                      //
//******************************************************************//

type EntityMarshallingError struct {
	Inner error
}

func (err EntityMarshallingError) Error() string {
	return fmt.Sprintf("Error marshalling/unmarshalling entity: %v", err.Inner)
}

func (err EntityMarshallingError) Unwrap() error {
	return err.Inner
}

func (err EntityMarshallingError) Is(other error) bool {
	switch other.(type) {
	case EntityMarshallingError:
		return true
	default:
		return false
	}
}

func (err EntityMarshallingError) As(target any) bool {

	switch target.(type) {
	case EntityMarshallingError:
		reflect.Indirect(reflect.ValueOf(target)).Set(reflect.ValueOf(err))
		return true
	default:
		return false
	}
}

//******************************************************************//
//                        EmptyIterableError                        //
//******************************************************************//

type ConflictingPositionalArgsAndFlagError struct {
	Flag string
}

func (err ConflictingPositionalArgsAndFlagError) Error() string {
	return "Can not use positional arguments together with the flag " + err.Flag
}

func (err ConflictingPositionalArgsAndFlagError) Is(other error) bool {
	switch other.(type) {
	case ConflictingPositionalArgsAndFlagError:
		return true
	default:
		return false
	}
}

func (err ConflictingPositionalArgsAndFlagError) As(target any) bool {
	switch target.(type) {
	case ConflictingPositionalArgsAndFlagError:
		reflect.Indirect(reflect.ValueOf(target)).Set(reflect.ValueOf(err))
		return true
	default:
		return false
	}
}

//******************************************************************//
//                     Non entity output structs                    //
//******************************************************************//

type DoesExist struct {
	DoesExist bool `json:"doesExist"`
}

type Count struct {
	Count int64 `json:"count"`
}

type NumAffectedRecords struct {
	NumAffectedRecords int64 `json:"numAffectedRecords"`
}
