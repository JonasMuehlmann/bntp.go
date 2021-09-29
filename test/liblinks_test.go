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

// ###############
// # ListLinks() #
// ###############
func TestListLinksNoneExist(t *testing.T) {
	db, err := GetDB(t)
	assert.NoError(t, err)

	_, err = liblinks.ListLinks(db, "Foo")
	assert.Error(t, err)
}

func TestListLinksOne(t *testing.T) {
	db, err := GetDB(t)
	assert.NoError(t, err)

	err = libdocuments.AddDocument(db, nil, "Foo", "Bar")
	assert.NoError(t, err)

	err = libdocuments.AddDocument(db, nil, "Foo2", "Bar")
	assert.NoError(t, err)

	err = liblinks.AddLink(db, nil, "Foo", "Foo2")
	assert.NoError(t, err)

	links, err := liblinks.ListLinks(db, "Foo")
	assert.NoError(t, err)
	assert.Len(t, links, 1)
	assert.Equal(t, "Foo2", links[0])
}

func TestListLinksMany(t *testing.T) {
	db, err := GetDB(t)
	assert.NoError(t, err)

	err = libdocuments.AddDocument(db, nil, "Foo", "Bar")
	assert.NoError(t, err)

	err = libdocuments.AddDocument(db, nil, "Foo2", "Bar")
	assert.NoError(t, err)

	err = libdocuments.AddDocument(db, nil, "Foo3", "Bar")
	assert.NoError(t, err)

	err = libdocuments.AddDocument(db, nil, "Foo4", "Bar")
	assert.NoError(t, err)

	err = liblinks.AddLink(db, nil, "Foo", "Foo2")
	assert.NoError(t, err)

	err = liblinks.AddLink(db, nil, "Foo", "Foo3")
	assert.NoError(t, err)

	err = liblinks.AddLink(db, nil, "Foo", "Foo4")
	assert.NoError(t, err)

	links, err := liblinks.ListLinks(db, "Foo")
	assert.NoError(t, err)
	assert.Len(t, links, 3)
	assert.Equal(t, []string{"Foo2", "Foo3", "Foo4"}, links)
}
