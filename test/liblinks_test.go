package test

import (
	"testing"

	"github.com/JonasMuehlmann/bntp.go/internal/libdocuments"
	"github.com/JonasMuehlmann/bntp.go/internal/liblinks"
	"github.com/stretchr/testify/assert"
)

// #############
// # AddLink() #
// #############
func TestAddLink(t *testing.T) {
	db, err := GetDB(t)
	assert.NoError(t, err)

	err = libdocuments.AddDocument(db, nil, "Foo", "Bar")
	assert.NoError(t, err)

	err = libdocuments.AddDocument(db, nil, "Foo2", "Bar2")
	assert.NoError(t, err)

	err = liblinks.AddLink(db, nil, "Foo", "Foo2")
	assert.NoError(t, err)
}

func TestAddLinkSourceDoesNotExist(t *testing.T) {
	db, err := GetDB(t)
	assert.NoError(t, err)

	err = libdocuments.AddDocument(db, nil, "Foo2", "Bar2")
	assert.NoError(t, err)

	err = liblinks.AddLink(db, nil, "Foo", "Foo2")
	assert.Error(t, err)
}

func TestAddLinkDestionationDoesNotExist(t *testing.T) {
	db, err := GetDB(t)
	assert.NoError(t, err)

	err = libdocuments.AddDocument(db, nil, "Foo", "Bar")
	assert.NoError(t, err)

	err = liblinks.AddLink(db, nil, "Foo", "Foo2")
	assert.Error(t, err)
}

func TestAddLinkNoneExist(t *testing.T) {
	db, err := GetDB(t)
	assert.NoError(t, err)

	err = liblinks.AddLink(db, nil, "Foo", "Foo2")
	assert.Error(t, err)
}

func TestAddLinkSelfReference(t *testing.T) {
	db, err := GetDB(t)
	assert.NoError(t, err)

	err = libdocuments.AddDocument(db, nil, "Foo", "Bar")
	assert.NoError(t, err)

	err = liblinks.AddLink(db, nil, "Foo", "Foo")
	assert.Error(t, err)
}

// ################
// # RemoveLink() #
// ################
func TestRemoveLink(t *testing.T) {
	db, err := GetDB(t)
	assert.NoError(t, err)

	err = libdocuments.AddDocument(db, nil, "Foo", "Bar")
	assert.NoError(t, err)

	err = libdocuments.AddDocument(db, nil, "Foo2", "Bar2")
	assert.NoError(t, err)

	err = liblinks.AddLink(db, nil, "Foo", "Foo2")
	assert.NoError(t, err)

	err = liblinks.RemoveLink(db, nil, "Foo", "Foo2")
	assert.NoError(t, err)
}

func TestRemoveLinkSourceDoesNotExist(t *testing.T) {
	db, err := GetDB(t)
	assert.NoError(t, err)

	err = libdocuments.AddDocument(db, nil, "Foo2", "Bar2")
	assert.NoError(t, err)

	err = liblinks.AddLink(db, nil, "Foo", "Foo2")
	assert.Error(t, err)

	err = liblinks.RemoveLink(db, nil, "Bar", "Foo2")
	assert.Error(t, err)
}

func TestRemoveLinkDestionationDoesNotExist(t *testing.T) {
	db, err := GetDB(t)
	assert.NoError(t, err)

	err = libdocuments.AddDocument(db, nil, "Foo", "Bar")
	assert.NoError(t, err)

	err = liblinks.AddLink(db, nil, "Foo", "Foo2")
	assert.Error(t, err)

	err = liblinks.RemoveLink(db, nil, "Foo", "Bar")
	assert.Error(t, err)
}

func TestRemoveLinkNoneExist(t *testing.T) {
	db, err := GetDB(t)
	assert.NoError(t, err)

	err = liblinks.RemoveLink(db, nil, "Foo", "Foo2")
	assert.Error(t, err)
}
