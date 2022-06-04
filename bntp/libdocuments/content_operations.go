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
	"strings"

	"github.com/JonasMuehlmann/goaoi"
)

// TODO: Implement context handling.
func AddTags(ctx context.Context, content string, tags []string) (string, error) {
	lines := strings.Split(content, "\n")

	iTagLine, err := findTagsLine(lines)
	if err != nil {
		return "", err
	}

	lines[iTagLine] += strings.Join(tags, ",")

	return strings.Join(lines, "\n"), nil
}

func RemoveTags(ctx context.Context, content string, tags []string) (string, error) {
	lines := strings.Split(content, "\n")

	iTagLine, err := findTagsLine(lines)
	if err != nil {
		return "", err
	}

	lineTags := strings.Split(lines[iTagLine], ",")

	unary_predicate := func(tag string) bool {
		_, err := goaoi.FindSlice(tags, tag)

		return !errors.As(err, &goaoi.ElementNotFoundError{})
	}

	lineTags, err = goaoi.CopyExceptIfSlice(lineTags, unary_predicate)
	if err != nil {
		return "", err
	}

	lines[iTagLine] = strings.Join(lineTags, "\n")

	return strings.Join(lines, "\n"), nil
}

func AddLinks(ctx context.Context, content string, links []string) (string, error) {
	lines := strings.Split(content, "\n")

	iLinksLinesStart, err := findLinksLinesStart(lines)
	if err != nil {
		return "", err
	}

	unary_predicate := func(line string) bool {
		return strings.HasPrefix(line, "- ")
	}

	iLinksLinesEnd, err := goaoi.FindIfSlice(lines[iLinksLinesStart:], unary_predicate)
	if err != nil {
		return "", err
	}

	transformer := func(link string) string {
		return "- ()[" + link + "]"
	}

	newLinksLines, err := goaoi.TransformCopySliceUnsafe(links, transformer)
	if err != nil {
		return "", err
	}

	lines = append(lines[:iLinksLinesStart], append(newLinksLines, lines[:iLinksLinesEnd]...)...)

	return strings.Join(lines, "\n"), nil
}

func RemoveLinks(ctx context.Context, content string, links []string) (string, error) {
	lines := strings.Split(content, "\n")

	iLinksLinesStart, err := findLinksLinesStart(lines)
	if err != nil {
		return "", err
	}

	unary_predicate := func(line string) bool {
		return strings.HasPrefix(line, "- ")
	}

	iLinksLinesEnd, err := goaoi.FindIfSlice(lines[iLinksLinesStart:], unary_predicate)
	if err != nil {
		return "", err
	}

	transformer := func(link string) string {
		return "- ()[" + link + "]"
	}
	linksLinesToRemove, err := goaoi.TransformCopySliceUnsafe(links, transformer)
	if err != nil {
		return "", err
	}

	unary_predicate = func(link string) bool {
		_, err := goaoi.FindSlice(linksLinesToRemove, link)

		return !errors.As(err, &goaoi.ElementNotFoundError{})
	}

	newLinksLines, err := goaoi.CopyExceptIfSlice(links[iLinksLinesStart:iLinksLinesEnd], unary_predicate)
	if err != nil {
		return "", err
	}

	lines = append(lines[:iLinksLinesStart], append(newLinksLines, lines[:iLinksLinesEnd]...)...)

	return strings.Join(lines, "\n"), nil
}

func AddBacklinks(ctx context.Context, content string, backlinks []string) (string, error) {
	lines := strings.Split(content, "\n")

	iBacklinksLinesStart, err := findBacklinksLinesStart(lines)
	if err != nil {
		return "", err
	}

	unary_predicate := func(line string) bool {
		return strings.HasPrefix(line, "- ")
	}

	iBacklinksLinesEnd, err := goaoi.FindIfSlice(lines[iBacklinksLinesStart:], unary_predicate)
	if err != nil {
		return "", err
	}

	transformer := func(link string) string {
		return "- ()[" + link + "]"
	}

	newBacklinksLines, err := goaoi.TransformCopySliceUnsafe(backlinks, transformer)
	if err != nil {
		return "", err
	}

	lines = append(lines[:iBacklinksLinesStart], append(newBacklinksLines, lines[:iBacklinksLinesEnd]...)...)

	return strings.Join(lines, "\n"), nil
}

func RemoveBacklinks(ctx context.Context, content string, backlinks []string) (string, error) {
	lines := strings.Split(content, "\n")

	iBacklinksLinesStart, err := findBacklinksLinesStart(lines)
	if err != nil {
		return "", err
	}

	unary_predicate := func(line string) bool {
		return strings.HasPrefix(line, "- ")
	}

	iBacklinksLinesEnd, err := goaoi.FindIfSlice(lines[iBacklinksLinesStart:], unary_predicate)
	if err != nil {
		return "", err
	}

	transformer := func(link string) string {
		return "- ()[" + link + "]"
	}
	linksLinesToRemove, err := goaoi.TransformCopySliceUnsafe(backlinks, transformer)
	if err != nil {
		return "", err
	}

	unary_predicate = func(link string) bool {
		_, err := goaoi.FindSlice(linksLinesToRemove, link)

		return !errors.As(err, &goaoi.ElementNotFoundError{})
	}

	newBacklinksLines, err := goaoi.CopyExceptIfSlice(backlinks[iBacklinksLinesStart:iBacklinksLinesEnd], unary_predicate)
	if err != nil {
		return "", err
	}

	lines = append(lines[:iBacklinksLinesStart], append(newBacklinksLines, lines[:iBacklinksLinesEnd]...)...)

	return strings.Join(lines, "\n"), nil
}

func findTagsLine(lines []string) (int, error) {
	iTagsLineHeading, err := goaoi.FindSlice(lines, "# Tags")

	return 1 + iTagsLineHeading, err
}

func findLinksLinesStart(lines []string) (int, error) {
	iLinksLineHeading, err := goaoi.FindSlice(lines, "# Links")

	return 1 + iLinksLineHeading, err
}

func findBacklinksLinesStart(lines []string) (int, error) {
	iBacklinksLineHeading, err := goaoi.FindSlice(lines, "# Backlinks")

	return 1 + iBacklinksLineHeading, err
}
