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
