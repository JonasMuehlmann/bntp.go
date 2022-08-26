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

package libdocuments

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/JonasMuehlmann/bntp.go/internal/helper"
	"github.com/JonasMuehlmann/goaoi"
	"github.com/JonasMuehlmann/goaoi/functional"
)

type DocumentContentEntity string

const (
	DocumentContentEntityTag      DocumentContentEntity = "tag"
	DocumentContentEntityLink     DocumentContentEntity = "link"
	DocumentContentEntityBacklink DocumentContentEntity = "backlink"
)

//******************************************************************//
//                        DocumentSyntaxError                       //
//******************************************************************//

type DocumentSyntaxError struct {
	Inner error
}

func (err DocumentSyntaxError) Error() string {
	return fmt.Sprintf("Error processing document: %v", err.Inner)
}

func (err DocumentSyntaxError) Unwrap() error {
	return err.Inner
}

func (err DocumentSyntaxError) Is(other error) bool {
	switch other.(type) {
	case DocumentSyntaxError:
		return true
	default:
		return false
	}
}

func (err DocumentSyntaxError) As(target any) bool {
	switch target.(type) {
	case DocumentSyntaxError:
		reflect.Indirect(reflect.ValueOf(target)).Set(reflect.ValueOf(err))

		return true
	default:
		return false
	}
}

//******************************************************************//
//                      TagsHeaderNotFoundError                     //
//******************************************************************//

type TagsHeaderNotFoundError struct {
	Inner error
}

func (err TagsHeaderNotFoundError) Error() string {
	return fmt.Sprintf("Error searching for tags header: %v", err.Inner)
}

func (err TagsHeaderNotFoundError) Unwrap() error {
	return err.Inner
}

func (err TagsHeaderNotFoundError) Is(other error) bool {
	switch other.(type) {
	case TagsHeaderNotFoundError:
		return true
	default:
		return false
	}
}

func (err TagsHeaderNotFoundError) As(target any) bool {
	switch target.(type) {
	case TagsHeaderNotFoundError:
		reflect.Indirect(reflect.ValueOf(target)).Set(reflect.ValueOf(err))

		return true
	default:
		return false
	}
}

//******************************************************************//
//                      TagsHeaderNotFoundError                     //
//******************************************************************//

type LinksHeaderNotFoundError struct {
	Inner error
}

func (err LinksHeaderNotFoundError) Error() string {
	return fmt.Sprintf("Error searching for tags header: %v", err.Inner)
}

func (err LinksHeaderNotFoundError) Unwrap() error {
	return err.Inner
}

func (err LinksHeaderNotFoundError) Is(other error) bool {
	switch other.(type) {
	case LinksHeaderNotFoundError:
		return true
	default:
		return false
	}
}

func (err LinksHeaderNotFoundError) As(target any) bool {
	switch target.(type) {
	case LinksHeaderNotFoundError:
		reflect.Indirect(reflect.ValueOf(target)).Set(reflect.ValueOf(err))

		return true
	default:
		return false
	}
}

//******************************************************************//
//                      TagsHeaderNotFoundError                     //
//******************************************************************//

type BacklinksHeaderNotFoundError struct {
	Inner error
}

func (err BacklinksHeaderNotFoundError) Error() string {
	return fmt.Sprintf("Error searching for tags header: %v", err.Inner)
}

func (err BacklinksHeaderNotFoundError) Unwrap() error {
	return err.Inner
}

func (err BacklinksHeaderNotFoundError) Is(other error) bool {
	switch other.(type) {
	case BacklinksHeaderNotFoundError:
		return true
	default:
		return false
	}
}

func (err BacklinksHeaderNotFoundError) As(target any) bool {
	switch target.(type) {
	case BacklinksHeaderNotFoundError:
		reflect.Indirect(reflect.ValueOf(target)).Set(reflect.ValueOf(err))

		return true
	default:
		return false
	}
}

//******************************************************************//
//                     EmptyEntitiesListError                       //
//******************************************************************//

type EmptyEntitiesListError struct {
	Entity DocumentContentEntity
}

func (err EmptyEntitiesListError) Error() string {
	return fmt.Sprintf("The list for entity type %v is empty", err.Entity)
}

func (err EmptyEntitiesListError) Is(other error) bool {
	switch other.(type) {
	case EmptyEntitiesListError:
		return true
	default:
		return false
	}
}

func (err EmptyEntitiesListError) As(target any) bool {
	switch target.(type) {
	case EmptyEntitiesListError:
		reflect.Indirect(reflect.ValueOf(target)).Set(reflect.ValueOf(err))

		return true
	default:
		return false
	}
}

// TODO: Implement context handling.
func AddTags(ctx context.Context, content string, tags []string) (newContent string, err error) {
	if len(tags) == 0 {
		return "", helper.IneffectiveOperationError{Inner: helper.EmptyInputError{}}
	}
	if content == "" {
		return "", helper.IneffectiveOperationError{Inner: helper.NonExistentPrimaryDataError{}}
	}

	err = goaoi.AnyOfSlice(tags, functional.IsZero[string])
	if err == nil {
		err = helper.NilInputError{}

		return
	}

	lines := strings.Split(content, "\n")

	iTagLine, err := findTagsLine(lines)
	if err != nil {
		if errors.Is(err, goaoi.ElementNotFoundError{}) {
			err = DocumentSyntaxError{Inner: TagsHeaderNotFoundError{Inner: err}}
		}

		return
	}

	// No tag line exists, and line is end of file
	if iTagLine == len(lines) {
		lines = append(lines, strings.Join(tags, ","))
	} else {
		// Tag line has some tags, make sure not to break enumeration syntax
		if lines[iTagLine] != "" {
			lines[iTagLine] += ","
		}

		lines[iTagLine] += strings.Join(tags, ",")
	}

	return strings.Join(lines, "\n"), nil
}

func RemoveTags(ctx context.Context, content string, tags []string) (newContent string, err error) {
	if len(tags) == 0 {
		return "", helper.IneffectiveOperationError{Inner: helper.EmptyInputError{}}
	}
	if content == "" {
		return "", helper.IneffectiveOperationError{Inner: helper.NonExistentPrimaryDataError{}}
	}

	err = goaoi.AnyOfSlice(tags, functional.IsZero[string])
	if err == nil {
		err = helper.NilInputError{}

		return
	}

	lines := strings.Split(content, "\n")

	iTagLine, err := findTagsLine(lines)
	if err != nil {
		if errors.Is(err, goaoi.ElementNotFoundError{}) {
			err = DocumentSyntaxError{Inner: TagsHeaderNotFoundError{Inner: err}}
		}

		return
	}

	if iTagLine == len(lines) {
		return "", helper.IneffectiveOperationError{Inner: helper.NonExistentDependencyError{Inner: EmptyEntitiesListError{Entity: DocumentContentEntityTag}}}

	}
	lineTags := strings.Split(lines[iTagLine], ",")

	unary_predicate := func(tag string) bool {
		_, err := goaoi.FindIfSlice(tags, functional.AreEqualPartial(tag))

		return err == nil
	}

	// PERF: Consider using a set here
	lineTags, err = goaoi.TakeIfSlice(lineTags, functional.NegateUnaryPredicate(unary_predicate))
	if err != nil {
		return
	}

	lines[iTagLine] = strings.Join(lineTags, "\n")

	return strings.Join(lines, "\n"), nil
}

func AddLinks(ctx context.Context, content string, links []string) (newContent string, err error) {
	if len(links) == 0 {
		return "", helper.IneffectiveOperationError{Inner: helper.EmptyInputError{}}
	}
	if content == "" {
		return "", helper.IneffectiveOperationError{Inner: helper.NonExistentPrimaryDataError{}}
	}

	err = goaoi.AnyOfSlice(links, functional.IsZero[string])
	if err == nil {
		err = helper.NilInputError{}

		return
	}

	lines := strings.Split(content, "\n")

	iLinksLinesStart, err := findLinksLinesStart(lines)
	if err != nil {
		if errors.Is(err, goaoi.ElementNotFoundError{}) {
			err = DocumentSyntaxError{Inner: LinksHeaderNotFoundError{Inner: err}}
		}

		return
	}

	unary_predicate := func(line string) bool {
		return !strings.HasPrefix(line, "- ")
	}

	iLinksLinesEnd, err := goaoi.FindIfSlice(lines[iLinksLinesStart:], unary_predicate)
	if errors.Is(err, goaoi.EmptyIterableError{}) {
		// No links exist yet, make space for new ones
		linesNew := make([]string, 0, len(lines)+len(links))
		linesNew = append(linesNew, lines...)
		lines = linesNew

		err = nil
	}
	if err != nil && !errors.Is(err, goaoi.ElementNotFoundError{}) {
		return
	}

	var linesAfterLinkLine []string
	var linesBeforeLinksLineEnd []string

	if !errors.Is(err, goaoi.ElementNotFoundError{}) {
		linesAfterLinkLine = lines[iLinksLinesEnd+1:]
		linesBeforeLinksLineEnd = lines[:iLinksLinesEnd+1]
	} else {
		linesBeforeLinksLineEnd = lines
	}

	transformer := func(link string) string {
		return "- (" + link + ")[" + link + "]"
	}

	newLinksLines, err := goaoi.TransformCopySliceUnsafe(links, transformer)
	if err != nil {
		return
	}

	lines = append(linesBeforeLinksLineEnd, append(newLinksLines, linesAfterLinkLine...)...)

	return strings.Join(lines, "\n"), nil
}

func RemoveLinks(ctx context.Context, content string, links []string) (newContent string, err error) {
	if len(links) == 0 {
		return "", helper.IneffectiveOperationError{Inner: helper.EmptyInputError{}}
	}
	if content == "" {
		return "", helper.IneffectiveOperationError{Inner: helper.NonExistentPrimaryDataError{}}
	}

	err = goaoi.AnyOfSlice(links, functional.IsZero[string])
	if err == nil {
		err = helper.NilInputError{}

		return
	}

	lines := strings.Split(content, "\n")

	iLinksLinesStart, err := findLinksLinesStart(lines)
	if err != nil {
		if errors.Is(err, goaoi.ElementNotFoundError{}) {
			err = DocumentSyntaxError{Inner: LinksHeaderNotFoundError{Inner: err}}
		}

		return
	}

	unary_predicate := func(line string) bool {
		return !strings.HasPrefix(line, "- ")
	}

	iLinksLinesEnd, err := goaoi.FindIfSlice(lines[iLinksLinesStart:], unary_predicate)
	if errors.Is(err, goaoi.ElementNotFoundError{}) || errors.Is(err, goaoi.EmptyIterableError{}) {
		iLinksLinesEnd = len(lines)
	} else if err != nil {
		return
	} else {
		iLinksLinesEnd += iLinksLinesStart
	}

	if iLinksLinesStart == iLinksLinesEnd {
		return "", helper.IneffectiveOperationError{Inner: helper.NonExistentDependencyError{Inner: EmptyEntitiesListError{Entity: DocumentContentEntityLink}}}

	}

	transformer := func(link string) string {
		return "- (" + link + ")[" + link + "]"
	}
	linksLinesToRemove, err := goaoi.TransformCopySliceUnsafe(links, transformer)
	if err != nil {
		return
	}

	unary_predicate = func(link string) bool {
		_, err := goaoi.FindIfSlice(linksLinesToRemove, functional.AreEqualPartial(link))

		return err == nil
	}

	newLinksLines, err := goaoi.TakeIfSlice(lines[iLinksLinesStart:iLinksLinesEnd], functional.NegateUnaryPredicate(unary_predicate))
	if err != nil {
		return
	}
	if len(newLinksLines) == 0 {
		lines = append(lines, "")
	}

	lines = append(lines[:iLinksLinesStart], append(newLinksLines, lines[iLinksLinesEnd:]...)...)

	return strings.Join(lines, "\n"), nil
}

func AddBacklinks(ctx context.Context, content string, links []string) (newContent string, err error) {
	if len(links) == 0 {
		return "", helper.IneffectiveOperationError{Inner: helper.EmptyInputError{}}
	}
	if content == "" {
		return "", helper.IneffectiveOperationError{Inner: helper.NonExistentPrimaryDataError{}}
	}

	err = goaoi.AnyOfSlice(links, functional.IsZero[string])
	if err == nil {
		err = helper.NilInputError{}

		return
	}

	lines := strings.Split(content, "\n")

	iBacklinksLinesStart, err := findBacklinksLinesStart(lines)
	if err != nil {
		if errors.Is(err, goaoi.ElementNotFoundError{}) {
			err = DocumentSyntaxError{Inner: BacklinksHeaderNotFoundError{Inner: err}}
		}

		return
	}

	unary_predicate := func(line string) bool {
		return !strings.HasPrefix(line, "- ")
	}

	iBacklinksLinesEnd, err := goaoi.FindIfSlice(lines[iBacklinksLinesStart:], unary_predicate)
	if errors.Is(err, goaoi.EmptyIterableError{}) {
		// No links exist yet, make space for new ones
		linesNew := make([]string, 0, len(lines)+len(links))
		linesNew = append(linesNew, lines...)
		lines = linesNew

		err = nil
	}
	if err != nil && !errors.Is(err, goaoi.ElementNotFoundError{}) {
		return
	}

	var linesAfterBacklinkLine []string
	var linesBeforeBacklinksLineEnd []string

	if !errors.Is(err, goaoi.ElementNotFoundError{}) {
		linesAfterBacklinkLine = lines[iBacklinksLinesEnd+1:]
		linesBeforeBacklinksLineEnd = lines[:iBacklinksLinesEnd+1]
	} else {
		linesBeforeBacklinksLineEnd = lines
	}

	transformer := func(link string) string {
		return "- (" + link + ")[" + link + "]"
	}

	newBacklinksLines, err := goaoi.TransformCopySliceUnsafe(links, transformer)
	if err != nil {
		return
	}

	lines = append(linesBeforeBacklinksLineEnd, append(newBacklinksLines, linesAfterBacklinkLine...)...)

	return strings.Join(lines, "\n"), nil
}

func RemoveBacklinks(ctx context.Context, content string, links []string) (newContent string, err error) {
	if len(links) == 0 {
		return "", helper.IneffectiveOperationError{Inner: helper.EmptyInputError{}}
	}
	if content == "" {
		return "", helper.IneffectiveOperationError{Inner: helper.NonExistentPrimaryDataError{}}
	}

	err = goaoi.AnyOfSlice(links, functional.IsZero[string])
	if err == nil {
		err = helper.NilInputError{}

		return
	}

	lines := strings.Split(content, "\n")

	iBacklinksLinesStart, err := findBacklinksLinesStart(lines)
	if err != nil {
		if errors.Is(err, goaoi.ElementNotFoundError{}) {
			err = DocumentSyntaxError{Inner: BacklinksHeaderNotFoundError{Inner: err}}
		}

		return
	}

	unary_predicate := func(line string) bool {
		return !strings.HasPrefix(line, "- ")
	}

	iBacklinksLinesEnd, err := goaoi.FindIfSlice(lines[iBacklinksLinesStart:], unary_predicate)
	if errors.Is(err, goaoi.ElementNotFoundError{}) || errors.Is(err, goaoi.EmptyIterableError{}) {
		iBacklinksLinesEnd = len(lines)
	} else if err != nil {
		return
	} else {
		iBacklinksLinesEnd += iBacklinksLinesStart
	}

	if iBacklinksLinesStart == iBacklinksLinesEnd {
		return "", helper.IneffectiveOperationError{Inner: helper.NonExistentDependencyError{Inner: EmptyEntitiesListError{Entity: DocumentContentEntityBacklink}}}

	}

	transformer := func(link string) string {
		return "- (" + link + ")[" + link + "]"
	}
	linksLinesToRemove, err := goaoi.TransformCopySliceUnsafe(links, transformer)
	if err != nil {
		return
	}

	unary_predicate = func(link string) bool {
		_, err := goaoi.FindIfSlice(linksLinesToRemove, functional.AreEqualPartial(link))

		return err == nil
	}

	newBacklinksLines, err := goaoi.TakeIfSlice(lines[iBacklinksLinesStart:iBacklinksLinesEnd], functional.NegateUnaryPredicate(unary_predicate))
	if err != nil {
		return
	}
	if len(newBacklinksLines) == 0 {
		lines = append(lines, "")
	}

	lines = append(lines[:iBacklinksLinesStart], append(newBacklinksLines, lines[iBacklinksLinesEnd:]...)...)

	return strings.Join(lines, "\n"), nil
}

func findTagsLine(lines []string) (int, error) {
	iTagsLineHeading, err := goaoi.FindIfSlice(lines, functional.AreEqualPartial("# Tags"))

	return 1 + iTagsLineHeading, err
}

func findLinksLinesStart(lines []string) (int, error) {
	iLinksLineHeading, err := goaoi.FindIfSlice(lines, functional.AreEqualPartial("# Links"))

	return 1 + iLinksLineHeading, err
}

func findBacklinksLinesStart(lines []string) (int, error) {
	iBacklinksLineHeading, err := goaoi.FindIfSlice(lines, functional.AreEqualPartial("# Backlinks"))

	return 1 + iBacklinksLineHeading, err
}
