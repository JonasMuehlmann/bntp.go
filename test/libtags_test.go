package test

import (
	"testing"

	"github.com/JonasMuehlmann/bntp.go/internal/libtags"
	"github.com/stretchr/testify/assert"
)

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
