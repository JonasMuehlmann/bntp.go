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

package bntp

import "github.com/barweiss/go-tuple"

func TupleToSOA2[T1 any, T2 any](aos []tuple.T2[T1, T2]) (res tuple.T2[[]T1, []T2]) {
	res.V1 = make([]T1, 0, len(aos))
	res.V2 = make([]T2, 0, len(aos))

	for _, s := range aos {
		res.V1 = append(res.V1, s.V1)
		res.V2 = append(res.V2, s.V2)
	}

	return res
}

func TupleToAOS2[T1 any, T2 any](soa tuple.T2[[]T1, []T2]) (res []tuple.T2[T1, T2]) {
	res = make([]tuple.T2[T1, T2], 0, len(soa.V1))

	for i := range soa.V1 {
		res = append(res, tuple.T2[T1, T2]{V1: soa.V1[i], V2: soa.V2[i]})
	}

	return res
}
