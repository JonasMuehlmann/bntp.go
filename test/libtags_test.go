package test

import (
	"testing"

	"github.com/JonasMuehlmann/bntp.go/internal/libtags"
	"github.com/stretchr/testify/assert"
)

// ############
// # AddTag() #
// ############
func TestAddTag(t *testing.T) {
	db, err := GetDB(t)
	assert.NoError(t, err)

	err = libtags.AddTag(db, nil, "Foo")
	assert.NoError(t, err)
}

func TestAddTagTransaction(t *testing.T) {
	db, err := GetDB(t)
	assert.NoError(t, err)

	transaction, err := db.Beginx()
	assert.NoError(t, err)

	err = libtags.AddTag(nil, transaction, "Foo")
	assert.NoError(t, err)

	err = transaction.Commit()
	assert.NoError(t, err)
}

func TestAddTagEmptyTag(t *testing.T) {
	db, err := GetDB(t)
	assert.NoError(t, err)

	err = libtags.AddTag(db, nil, "")
	assert.NoError(t, err)
}

// ###############
// # RemoveTag() #
// ###############
func TestRemoveTag(t *testing.T) {
	db, err := GetDB(t)
	assert.NoError(t, err)

	err = libtags.AddTag(db, nil, "Foo")
	assert.NoError(t, err)

	err = libtags.DeleteTag(db, nil, "Foo")
	assert.NoError(t, err)
}

func TestRemoveTagTransaction(t *testing.T) {
	db, err := GetDB(t)
	assert.NoError(t, err)

	transaction, err := db.Beginx()
	assert.NoError(t, err)

	err = libtags.AddTag(nil, transaction, "Foo")
	assert.NoError(t, err)

	err = libtags.DeleteTag(nil, transaction, "Foo")
	assert.NoError(t, err)

	err = transaction.Commit()
	assert.NoError(t, err)
}

func TestRemoveTagDoesNotExist(t *testing.T) {
	db, err := GetDB(t)
	assert.NoError(t, err)

	err = libtags.AddTag(db, nil, "")
	assert.NoError(t, err)

	err = libtags.AddTag(db, nil, "Bar")
	assert.NoError(t, err)
}

// ###############
// # RenameTag() #
// ###############
func TestRenameTag(t *testing.T) {
	db, err := GetDB(t)
	assert.NoError(t, err)

	err = libtags.AddTag(db, nil, "Foo")
	assert.NoError(t, err)

	err = libtags.RenameTag(db, nil, "Foo", "Bar")
	assert.NoError(t, err)
}

func TestRenameTagTransaction(t *testing.T) {
	db, err := GetDB(t)
	assert.NoError(t, err)

	transaction, err := db.Beginx()
	assert.NoError(t, err)

	err = libtags.AddTag(nil, transaction, "Foo")
	assert.NoError(t, err)

	err = libtags.RenameTag(nil, transaction, "Foo", "Bar")
	assert.NoError(t, err)

	err = transaction.Commit()
	assert.NoError(t, err)
}

func TestRenameTagNoOld(t *testing.T) {
	db, err := GetDB(t)
	assert.NoError(t, err)

	err = libtags.AddTag(db, nil, "")
	assert.NoError(t, err)

	err = libtags.RenameTag(db, nil, "XYZ", "Bar")
	assert.NoError(t, err)
}

func TestRenameTagNewEmpty(t *testing.T) {
	db, err := GetDB(t)
	assert.NoError(t, err)

	err = libtags.AddTag(db, nil, "")
	assert.NoError(t, err)

	err = libtags.RenameTag(db, nil, "Foo", "")
	assert.NoError(t, err)
}

// ##############
// # ListTags() #
// ##############
func TestListTagsOneTag(t *testing.T) {
	db, err := GetDB(t)
	assert.NoError(t, err)

	err = libtags.AddTag(db, nil, "Foo")
	assert.NoError(t, err)

	tagsAfter, err := libtags.ListTags(db)
	assert.NoError(t, err)
	assert.Len(t, tagsAfter, 1)
}

func TestListTagsManyTags(t *testing.T) {
	db, err := GetDB(t)
	assert.NoError(t, err)

	err = libtags.AddTag(db, nil, "Foo1")
	assert.NoError(t, err)

	err = libtags.AddTag(db, nil, "Foo2")
	assert.NoError(t, err)

	err = libtags.AddTag(db, nil, "Foo3")
	assert.NoError(t, err)

	err = libtags.AddTag(db, nil, "Foo4")
	assert.NoError(t, err)

	tagsAfter, err := libtags.ListTags(db)
	assert.NoError(t, err)
	assert.Len(t, tagsAfter, 4)
}

func TestListTagsEmpty(t *testing.T) {
	db, err := GetDB(t)
	assert.NoError(t, err)

	tagsBefore, err := libtags.ListTags(db)
	assert.NoError(t, err)
	assert.Len(t, tagsBefore, 0)
}

// #######################
// # ListTagsShortened() #
// #######################
func TestListTagsShortenedOneTagNoComponents(t *testing.T) {
	db, err := GetDB(t)
	assert.NoError(t, err)

	err = libtags.AddTag(db, nil, "Foo")
	assert.NoError(t, err)

	tags, err := libtags.ListTagsShortened(db)
	assert.NoError(t, err)
	assert.Len(t, tags, 1)
	assert.Equal(t, []string{"Foo"}, tags)
}

func TestListTagsShortenedOneTagManyComponents(t *testing.T) {
	db, err := GetDB(t)
	assert.NoError(t, err)

	err = libtags.AddTag(db, nil, "X::Y::Z")
	assert.NoError(t, err)

	tags, err := libtags.ListTagsShortened(db)
	assert.NoError(t, err)
	assert.Len(t, tags, 1)
	assert.Equal(t, []string{"Z"}, tags)
}

func TestListTagsShortenedManyTags(t *testing.T) {
	db, err := GetDB(t)
	assert.NoError(t, err)

	err = libtags.AddTag(db, nil, "X::Y::Z")
	assert.NoError(t, err)

	err = libtags.AddTag(db, nil, "A::B::C")
	assert.NoError(t, err)

	tags, err := libtags.ListTagsShortened(db)
	assert.NoError(t, err)
	assert.Len(t, tags, 2)
	assert.Equal(t, []string{"Z", "C"}, tags)
}

func TestListTagsShortenedManyTagsAmbiguousComponent(t *testing.T) {
	db, err := GetDB(t)
	assert.NoError(t, err)

	err = libtags.AddTag(db, nil, "X::Y::C")
	assert.NoError(t, err)

	err = libtags.AddTag(db, nil, "A::B::C")
	assert.NoError(t, err)

	tags, err := libtags.ListTagsShortened(db)
	assert.NoError(t, err)
	assert.Len(t, tags, 2)
	assert.Equal(t, []string{"Y::C", "B::C"}, tags)
}
