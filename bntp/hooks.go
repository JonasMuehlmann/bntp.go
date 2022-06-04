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
	"context"
	"errors"
	"fmt"

	"github.com/JonasMuehlmann/goaoi"
)

var BadHookPointError = errors.New("the hook point does not exist")

type HookExecutionError struct {
	Inner error
}

func (err HookExecutionError) Error() string {
	return fmt.Sprintf("Error Executing hooks: %v", err.Inner)
}

func (err HookExecutionError) Unwrap() error {
	return err.Inner
}

type HookPoint int

const (
	BeforeAddHook HookPoint = 1 << iota
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

	// TODO: Handle these hooks in managers.
	AfterErrorHook
	AfterDeadlineHook
	AfterTimeoutHook
	AfterCancelHook

	_end
)

type Hooks[TEntity any] struct {
	// Outer dimension is a fixed size array because there is a fixed number of hookpoints
	hooks [_end - 1][]func(context.Context, *TEntity) error
}

func (hooks *Hooks[TEntity]) AddHook(point HookPoint, hook func(context.Context, *TEntity) error) error {
	if point < 0 || point > _end {
		return BadHookPointError
	}

	for hp := 0; hp < int(_end); hp <<= 1 {
		hooks.hooks[hp] = append(hooks.hooks[hp], hook)
	}

	return nil
}

func (hooks *Hooks[TEntity]) ClearHooks(point HookPoint) error {
	if point < 0 || point > _end {
		return BadHookPointError
	}

	for hp := 0; hp < int(_end); hp <<= 1 {
		hooks.hooks[hp] = nil
	}

	return nil
}

func (hooks *Hooks[TEntity]) ExecuteHooks(ctx context.Context, point HookPoint, entity *TEntity) error {
	if point < 0 || point > _end {
		return BadHookPointError
	}

	for hp := 0; hp < int(_end); hp <<= 1 {
		err := goaoi.ForeachSlice(hooks.hooks[point], func(hook func(context.Context, *TEntity) error) error { return hook(ctx, entity) })
		if err != nil {
			return err
		}
	}

	return nil
}

func (hooks *Hooks[TEntity]) PartiallySpecializeExecuteHooks(ctx context.Context, point HookPoint) func(entity *TEntity) error {
	return func(entity *TEntity) error { return hooks.ExecuteHooks(ctx, point, entity) }
}

// NOTE: This is not great but finding a better solution is hard with this language.
func (hooks *Hooks[TEntity]) PartiallySpecializeExecuteHooksForNoPointer(ctx context.Context, point HookPoint) func(entity TEntity) error {
	return func(entity TEntity) error { return hooks.ExecuteHooks(ctx, point, &entity) }
}
