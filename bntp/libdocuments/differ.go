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
	"github.com/JonasMuehlmann/bntp.go/model/domain"
	"github.com/JonasMuehlmann/goaoi"
)

func GetAddedLinks(old *domain.Document, new *domain.Document) (addedLinks []string, err error) {
	var addedLinksRaw []*domain.Document

	predicate := func(oldLink *domain.Document) bool {
		_, err := goaoi.FindSlice(new.LinkedDocuments, oldLink)

		return err == nil
	}

	transformer := func(rawLink *domain.Document) string {
		return rawLink.Path
	}

	addedLinksRaw, err = goaoi.CopyIfSlice(old.LinkedDocuments, predicate)
	if err != nil {
		return
	}

	addedLinks, err = goaoi.TransformCopySliceUnsafe(addedLinksRaw, transformer)

	return
}

func GetRemovedLinks(old *domain.Document, new *domain.Document) (removedLinks []string, err error) {
	var removedLinksRaw []*domain.Document

	predicate := func(oldLink *domain.Document) bool {
		_, err := goaoi.FindSlice(new.LinkedDocuments, oldLink)

		return err == nil
	}

	transformer := func(rawLink *domain.Document) string {
		return rawLink.Path
	}

	removedLinksRaw, err = goaoi.CopyIfSlice(old.LinkedDocuments, predicate)
	if err != nil {
		return
	}

	removedLinks, err = goaoi.TransformCopySliceUnsafe(removedLinksRaw, transformer)

	return
}

func GetAddedBacklinks(old *domain.Document, new *domain.Document) (addedBacklinks []string, err error) {
	var addedBacklinksRaw []*domain.Document

	predicate := func(oldBacklink *domain.Document) bool {
		_, err := goaoi.FindSlice(new.BacklinkedDocuments, oldBacklink)

		return err == nil
	}

	transformer := func(rawBacklink *domain.Document) string {
		return rawBacklink.Path
	}

	addedBacklinksRaw, err = goaoi.CopyIfSlice(old.BacklinkedDocuments, predicate)
	if err != nil {
		return
	}

	addedBacklinks, err = goaoi.TransformCopySliceUnsafe(addedBacklinksRaw, transformer)

	return
}

func GetRemovedBacklinks(old *domain.Document, new *domain.Document) (removedBacklinks []string, err error) {
	var removedBacklinksRaw []*domain.Document

	predicate := func(oldBacklink *domain.Document) bool {
		_, err := goaoi.FindSlice(new.BacklinkedDocuments, oldBacklink)

		return err == nil
	}

	transformer := func(rawBacklink *domain.Document) string {
		return rawBacklink.Path
	}

	removedBacklinksRaw, err = goaoi.CopyIfSlice(old.BacklinkedDocuments, predicate)
	if err != nil {
		return
	}

	removedBacklinks, err = goaoi.TransformCopySliceUnsafe(removedBacklinksRaw, transformer)

	return
}

func GetAddedTags(old *domain.Document, new *domain.Document) (addedTags []string, err error) {
	var addedTagsRaw []*domain.Tag

	predicate := func(oldTag *domain.Tag) bool {
		_, err := goaoi.FindSlice(new.Tags, oldTag)

		return err == nil
	}

	transformer := func(rawTag *domain.Tag) string {
		return rawTag.Tag
	}

	addedTagsRaw, err = goaoi.CopyIfSlice(old.Tags, predicate)
	if err != nil {
		return
	}

	addedTags, err = goaoi.TransformCopySliceUnsafe(addedTagsRaw, transformer)

	return
}

func GetRemovedTags(old *domain.Document, new *domain.Document) (removedTags []string, err error) {
	var removedTagsRaw []*domain.Tag

	predicate := func(oldTag *domain.Tag) bool {
		_, err := goaoi.FindSlice(new.Tags, oldTag)

		return err == nil
	}

	transformer := func(rawTag *domain.Tag) string {
		return rawTag.Tag
	}

	removedTagsRaw, err = goaoi.CopyIfSlice(old.Tags, predicate)
	if err != nil {
		return
	}

	removedTags, err = goaoi.TransformCopySliceUnsafe(removedTagsRaw, transformer)

	return
}
