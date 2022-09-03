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
	"github.com/JonasMuehlmann/goaoi/functional"
)

// TODO: Add error logging

func GetAddedLinks(old *domain.Document, new *domain.Document) (addedLinkIDs []int64, err error) {
	predicate := func(oldLink int64) bool {
		_, err := goaoi.FindIfSlice(new.LinkedDocumentIDs, functional.AreEqualPartial(oldLink))

		return err == nil
	}

	addedLinkIDs, err = goaoi.TakeIfSlice(old.LinkedDocumentIDs, predicate)

	return
}

func GetRemovedLinks(old *domain.Document, new *domain.Document) (removedLinkIDs []int64, err error) {
	predicate := func(oldLink int64) bool {
		_, err := goaoi.FindIfSlice(new.LinkedDocumentIDs, functional.AreEqualPartial(oldLink))

		return err == nil
	}

	removedLinkIDs, err = goaoi.TakeIfSlice(old.LinkedDocumentIDs, predicate)

	return
}

func GetAddedBacklinks(old *domain.Document, new *domain.Document) (addedBacklinkIDs []int64, err error) {
	predicate := func(oldBacklink int64) bool {
		_, err := goaoi.FindIfSlice(new.BacklinkedDocumentsIDs, functional.AreEqualPartial(oldBacklink))

		return err == nil
	}

	addedBacklinkIDs, err = goaoi.TakeIfSlice(old.BacklinkedDocumentsIDs, predicate)
	return
}

func GetRemovedBacklinks(old *domain.Document, new *domain.Document) (removedBacklinkIDs []int64, err error) {
	predicate := func(oldBacklink int64) bool {
		_, err := goaoi.FindIfSlice(new.BacklinkedDocumentsIDs, functional.AreEqualPartial(oldBacklink))

		return err == nil
	}

	removedBacklinkIDs, err = goaoi.TakeIfSlice(old.BacklinkedDocumentsIDs, predicate)
	return
}

func GetAddedTags(old *domain.Document, new *domain.Document) (addedTagIDs []int64, err error) {
	predicate := func(oldTag int64) bool {
		_, err := goaoi.FindIfSlice(new.TagIDs, functional.AreEqualPartial(oldTag))

		return err == nil
	}

	addedTagIDs, err = goaoi.TakeIfSlice(old.TagIDs, predicate)
	return
}

func GetRemovedTags(old *domain.Document, new *domain.Document) (removedTagIDs []int64, err error) {
	predicate := func(oldTag int64) bool {
		_, err := goaoi.FindIfSlice(new.TagIDs, functional.AreEqualPartial(oldTag))

		return err == nil
	}

	removedTagIDs, err = goaoi.TakeIfSlice(old.TagIDs, predicate)
	return
}
