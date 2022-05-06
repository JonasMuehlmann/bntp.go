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

import (
	"errors"
)

var BadHookPointError = errors.New("The hook point does not exist")

type hookPoint int

const (
	BeforeAddHook hookPoint = iota
	AfterAddHook

	BeforeSelectHook
	AfterSelectHook

	BeforeUpdateHook
	AfterUpdateHook

	BeforeDeleteHook
	AfterDeleteHook

	BeforeUpsertHook
	AfterUpsertHook

	BeforeAnyHook
	AfterAnyHook

	AfterError
	AfterDeadline
	AfterTimeout
	AfterCancel

	_end
)

type Hooks[TEntityHook any] struct {
	// Outer dimension is a fixed size array because there is a fixed number of hookpoints
	hooks [_end - 1][]TEntityHook
}

func (hooks Hooks[TEntityHook]) AddHook(point hookPoint, hook TEntityHook) error {
	if point > 0 && point < _end {
		hooks.hooks[point] = append(hooks.hooks[point], hook)
	}

	return BadHookPointError
}

func (hooks Hooks[TEntityHook]) ClearHooks(point hookPoint) error {
	if point > 0 && point < _end {
		hooks.hooks[point] = nil
	}

	return BadHookPointError
}
