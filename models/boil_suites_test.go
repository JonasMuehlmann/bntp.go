// Code generated by SQLBoiler 4.10.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("BookmarkTypes", testBookmarkTypes)
	t.Run("Bookmarks", testBookmarks)
	t.Run("DocumentTypes", testDocumentTypes)
	t.Run("Documents", testDocuments)
	t.Run("Tags", testTags)
}

func TestSoftDelete(t *testing.T) {}

func TestQuerySoftDeleteAll(t *testing.T) {}

func TestSliceSoftDeleteAll(t *testing.T) {}

func TestDelete(t *testing.T) {
	t.Run("BookmarkTypes", testBookmarkTypesDelete)
	t.Run("Bookmarks", testBookmarksDelete)
	t.Run("DocumentTypes", testDocumentTypesDelete)
	t.Run("Documents", testDocumentsDelete)
	t.Run("Tags", testTagsDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("BookmarkTypes", testBookmarkTypesQueryDeleteAll)
	t.Run("Bookmarks", testBookmarksQueryDeleteAll)
	t.Run("DocumentTypes", testDocumentTypesQueryDeleteAll)
	t.Run("Documents", testDocumentsQueryDeleteAll)
	t.Run("Tags", testTagsQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("BookmarkTypes", testBookmarkTypesSliceDeleteAll)
	t.Run("Bookmarks", testBookmarksSliceDeleteAll)
	t.Run("DocumentTypes", testDocumentTypesSliceDeleteAll)
	t.Run("Documents", testDocumentsSliceDeleteAll)
	t.Run("Tags", testTagsSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("BookmarkTypes", testBookmarkTypesExists)
	t.Run("Bookmarks", testBookmarksExists)
	t.Run("DocumentTypes", testDocumentTypesExists)
	t.Run("Documents", testDocumentsExists)
	t.Run("Tags", testTagsExists)
}

func TestFind(t *testing.T) {
	t.Run("BookmarkTypes", testBookmarkTypesFind)
	t.Run("Bookmarks", testBookmarksFind)
	t.Run("DocumentTypes", testDocumentTypesFind)
	t.Run("Documents", testDocumentsFind)
	t.Run("Tags", testTagsFind)
}

func TestBind(t *testing.T) {
	t.Run("BookmarkTypes", testBookmarkTypesBind)
	t.Run("Bookmarks", testBookmarksBind)
	t.Run("DocumentTypes", testDocumentTypesBind)
	t.Run("Documents", testDocumentsBind)
	t.Run("Tags", testTagsBind)
}

func TestOne(t *testing.T) {
	t.Run("BookmarkTypes", testBookmarkTypesOne)
	t.Run("Bookmarks", testBookmarksOne)
	t.Run("DocumentTypes", testDocumentTypesOne)
	t.Run("Documents", testDocumentsOne)
	t.Run("Tags", testTagsOne)
}

func TestAll(t *testing.T) {
	t.Run("BookmarkTypes", testBookmarkTypesAll)
	t.Run("Bookmarks", testBookmarksAll)
	t.Run("DocumentTypes", testDocumentTypesAll)
	t.Run("Documents", testDocumentsAll)
	t.Run("Tags", testTagsAll)
}

func TestCount(t *testing.T) {
	t.Run("BookmarkTypes", testBookmarkTypesCount)
	t.Run("Bookmarks", testBookmarksCount)
	t.Run("DocumentTypes", testDocumentTypesCount)
	t.Run("Documents", testDocumentsCount)
	t.Run("Tags", testTagsCount)
}

func TestHooks(t *testing.T) {
	t.Run("BookmarkTypes", testBookmarkTypesHooks)
	t.Run("Bookmarks", testBookmarksHooks)
	t.Run("DocumentTypes", testDocumentTypesHooks)
	t.Run("Documents", testDocumentsHooks)
	t.Run("Tags", testTagsHooks)
}

func TestInsert(t *testing.T) {
	t.Run("BookmarkTypes", testBookmarkTypesInsert)
	t.Run("BookmarkTypes", testBookmarkTypesInsertWhitelist)
	t.Run("Bookmarks", testBookmarksInsert)
	t.Run("Bookmarks", testBookmarksInsertWhitelist)
	t.Run("DocumentTypes", testDocumentTypesInsert)
	t.Run("DocumentTypes", testDocumentTypesInsertWhitelist)
	t.Run("Documents", testDocumentsInsert)
	t.Run("Documents", testDocumentsInsertWhitelist)
	t.Run("Tags", testTagsInsert)
	t.Run("Tags", testTagsInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("BookmarkToBookmarkTypeUsingBookmarkType", testBookmarkToOneBookmarkTypeUsingBookmarkType)
	t.Run("DocumentToDocumentTypeUsingDocumentType", testDocumentToOneDocumentTypeUsingDocumentType)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("BookmarkTypeToBookmarks", testBookmarkTypeToManyBookmarks)
	t.Run("BookmarkToTags", testBookmarkToManyTags)
	t.Run("DocumentTypeToDocuments", testDocumentTypeToManyDocuments)
	t.Run("DocumentToTags", testDocumentToManyTags)
	t.Run("DocumentToSourceDocuments", testDocumentToManySourceDocuments)
	t.Run("DocumentToDestinationDocuments", testDocumentToManyDestinationDocuments)
	t.Run("TagToBookmarks", testTagToManyBookmarks)
	t.Run("TagToDocuments", testTagToManyDocuments)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("BookmarkToBookmarkTypeUsingBookmarks", testBookmarkToOneSetOpBookmarkTypeUsingBookmarkType)
	t.Run("DocumentToDocumentTypeUsingDocuments", testDocumentToOneSetOpDocumentTypeUsingDocumentType)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {
	t.Run("BookmarkToBookmarkTypeUsingBookmarks", testBookmarkToOneRemoveOpBookmarkTypeUsingBookmarkType)
}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("BookmarkTypeToBookmarks", testBookmarkTypeToManyAddOpBookmarks)
	t.Run("BookmarkToTags", testBookmarkToManyAddOpTags)
	t.Run("DocumentTypeToDocuments", testDocumentTypeToManyAddOpDocuments)
	t.Run("DocumentToTags", testDocumentToManyAddOpTags)
	t.Run("DocumentToSourceDocuments", testDocumentToManyAddOpSourceDocuments)
	t.Run("DocumentToDestinationDocuments", testDocumentToManyAddOpDestinationDocuments)
	t.Run("TagToBookmarks", testTagToManyAddOpBookmarks)
	t.Run("TagToDocuments", testTagToManyAddOpDocuments)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {
	t.Run("BookmarkTypeToBookmarks", testBookmarkTypeToManySetOpBookmarks)
	t.Run("BookmarkToTags", testBookmarkToManySetOpTags)
	t.Run("DocumentToTags", testDocumentToManySetOpTags)
	t.Run("DocumentToSourceDocuments", testDocumentToManySetOpSourceDocuments)
	t.Run("DocumentToDestinationDocuments", testDocumentToManySetOpDestinationDocuments)
	t.Run("TagToBookmarks", testTagToManySetOpBookmarks)
	t.Run("TagToDocuments", testTagToManySetOpDocuments)
}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {
	t.Run("BookmarkTypeToBookmarks", testBookmarkTypeToManyRemoveOpBookmarks)
	t.Run("BookmarkToTags", testBookmarkToManyRemoveOpTags)
	t.Run("DocumentToTags", testDocumentToManyRemoveOpTags)
	t.Run("DocumentToSourceDocuments", testDocumentToManyRemoveOpSourceDocuments)
	t.Run("DocumentToDestinationDocuments", testDocumentToManyRemoveOpDestinationDocuments)
	t.Run("TagToBookmarks", testTagToManyRemoveOpBookmarks)
	t.Run("TagToDocuments", testTagToManyRemoveOpDocuments)
}

func TestReload(t *testing.T) {
	t.Run("BookmarkTypes", testBookmarkTypesReload)
	t.Run("Bookmarks", testBookmarksReload)
	t.Run("DocumentTypes", testDocumentTypesReload)
	t.Run("Documents", testDocumentsReload)
	t.Run("Tags", testTagsReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("BookmarkTypes", testBookmarkTypesReloadAll)
	t.Run("Bookmarks", testBookmarksReloadAll)
	t.Run("DocumentTypes", testDocumentTypesReloadAll)
	t.Run("Documents", testDocumentsReloadAll)
	t.Run("Tags", testTagsReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("BookmarkTypes", testBookmarkTypesSelect)
	t.Run("Bookmarks", testBookmarksSelect)
	t.Run("DocumentTypes", testDocumentTypesSelect)
	t.Run("Documents", testDocumentsSelect)
	t.Run("Tags", testTagsSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("BookmarkTypes", testBookmarkTypesUpdate)
	t.Run("Bookmarks", testBookmarksUpdate)
	t.Run("DocumentTypes", testDocumentTypesUpdate)
	t.Run("Documents", testDocumentsUpdate)
	t.Run("Tags", testTagsUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("BookmarkTypes", testBookmarkTypesSliceUpdateAll)
	t.Run("Bookmarks", testBookmarksSliceUpdateAll)
	t.Run("DocumentTypes", testDocumentTypesSliceUpdateAll)
	t.Run("Documents", testDocumentsSliceUpdateAll)
	t.Run("Tags", testTagsSliceUpdateAll)
}
