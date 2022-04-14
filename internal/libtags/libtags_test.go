package libtags_test

import (
	"testing"

	"github.com/JonasMuehlmann/bntp.go/internal/helpers"
	"github.com/JonasMuehlmann/bntp.go/internal/libtags"
	"github.com/JonasMuehlmann/bntp.go/test"
	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/maps"
)

// ############
// # AddTag() #
// ############.
func TestAddTag(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libtags.AddTag(db, nil, "Foo")
	assert.NoError(t, err)
}

func TestAddTagTransaction(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	transaction, err := db.Beginx()
	assert.NoError(t, err)

	err = libtags.AddTag(nil, transaction, "Foo")
	assert.NoError(t, err)

	err = transaction.Commit()
	assert.NoError(t, err)
}

func TestAddTagEmptyTag(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libtags.AddTag(db, nil, "")
	assert.NoError(t, err)
}

// ###############
// # RemoveTag() #
// ###############.
func TestRemoveTag(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libtags.AddTag(db, nil, "Foo")
	assert.NoError(t, err)

	err = libtags.DeleteTag(db, nil, "Foo")
	assert.NoError(t, err)
}

func TestRemoveTagTransaction(t *testing.T) {
	db, err := test.GetDB(t)
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
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libtags.AddTag(db, nil, "")
	assert.NoError(t, err)

	err = libtags.AddTag(db, nil, "Bar")
	assert.NoError(t, err)
}

// ###############
// # RenameTag() #
// ###############.
func TestRenameTag(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libtags.AddTag(db, nil, "Foo")
	assert.NoError(t, err)

	err = libtags.RenameTag(db, nil, "Foo", "Bar")
	assert.NoError(t, err)
}

func TestRenameTagTransaction(t *testing.T) {
	db, err := test.GetDB(t)
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
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libtags.AddTag(db, nil, "")
	assert.NoError(t, err)

	err = libtags.RenameTag(db, nil, "XYZ", "Bar")
	assert.NoError(t, err)
}

func TestRenameTagNewEmpty(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libtags.AddTag(db, nil, "")
	assert.NoError(t, err)

	err = libtags.RenameTag(db, nil, "Foo", "")
	assert.NoError(t, err)
}

// ##############
// # ListTags() #
// ##############.
func TestListTagsOneTag(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libtags.AddTag(db, nil, "Foo")
	assert.NoError(t, err)

	tagsAfter, err := libtags.ListTags(db)
	assert.NoError(t, err)
	assert.Len(t, tagsAfter, 1)
}

func TestListTagsManyTags(t *testing.T) {
	db, err := test.GetDB(t)
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
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	tagsBefore, err := libtags.ListTags(db)
	assert.NoError(t, err)
	assert.Len(t, tagsBefore, 0)
}

// #######################
// # ListTagsShortened() #
// #######################.
func TestListTagsShortenedOneTagNoComponents(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libtags.AddTag(db, nil, "Foo")
	assert.NoError(t, err)

	tags, err := libtags.ListTagsShortened(db)
	assert.NoError(t, err)
	assert.Len(t, tags, 1)
	assert.Equal(t, []string{"Foo"}, tags)
}

func TestListTagsShortenedOneTagManyComponents(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libtags.AddTag(db, nil, "X::Y::Z")
	assert.NoError(t, err)

	tags, err := libtags.ListTagsShortened(db)
	assert.NoError(t, err)
	assert.Len(t, tags, 1)
	assert.Equal(t, []string{"Z"}, tags)
}

func TestListTagsShortenedManyTags(t *testing.T) {
	db, err := test.GetDB(t)
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
	db, err := test.GetDB(t)
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

// #############################
// # DeserializeTagHierarchy() #
// #############################.
func TestDeserializeTagHierarchyNoTagsKey(t *testing.T) {
	yml := `
foo:
- bar
- baz
    `
	_, err := libtags.DeserializeTagHierarchy(yml)
	assert.ErrorAs(t, err, &helpers.DeserializationError{})
}

func TestDeserializeTagHierarchyNoTags(t *testing.T) {
	yml := `
tags:
    `
	_, err := libtags.DeserializeTagHierarchy(yml)
	assert.ErrorAs(t, err, &helpers.IneffectiveOperationError{})
}

func TestDeserializeTagHierarchyOnlyTopLevel(t *testing.T) {
	yml := `
tags:
- foo
- bar
- baz
    `
	tagHierarchy, err := libtags.DeserializeTagHierarchy(yml)
	assert.NoError(t, err)

	assert.Equal(t, []string{"tags"}, maps.Keys(tagHierarchy))
	assert.Subset(t, maps.Keys(tagHierarchy["tags"]), []string{"foo", "bar", "baz"})
}

func TestDeserializeTagHierarchyOnePath(t *testing.T) {
	yml := `
tags:
- foo:
    - bar:
        - baz
    `
	tagHierarchy, err := libtags.DeserializeTagHierarchy(yml)
	assert.NoError(t, err)

	assert.Equal(t, []string{"tags"}, maps.Keys(tagHierarchy))
	assert.Equal(t, []string{"foo"}, maps.Keys(tagHierarchy["tags"]))
	assert.Equal(t, []string{"bar"}, maps.Keys(tagHierarchy["tags"]["foo"]))
	assert.Equal(t, []string{"baz"}, maps.Keys(tagHierarchy["tags"]["foo"]["bar"]))
}

func TestDeserializeTagHierarchyTwoPaths(t *testing.T) {
	yml := `
tags:
- foo:
    - bar:
        - baz
- foo2:
    - bar2:
        - baz2
    `
	tagHierarchy, err := libtags.DeserializeTagHierarchy(yml)
	assert.NoError(t, err)

	assert.Equal(t, []string{"tags"}, maps.Keys(tagHierarchy))
	assert.Subset(t, maps.Keys(tagHierarchy["tags"]), []string{"foo", "foo2"})
	assert.Equal(t, []string{"bar"}, maps.Keys(tagHierarchy["tags"]["foo"]))
	assert.Equal(t, []string{"baz"}, maps.Keys(tagHierarchy["tags"]["foo"]["bar"]))

	assert.Equal(t, []string{"bar2"}, maps.Keys(tagHierarchy["tags"]["foo2"]))
	assert.Equal(t, []string{"baz2"}, maps.Keys(tagHierarchy["tags"]["foo2"]["bar2"]))
}

// ###########################
// # SerializeTagHierarchy() #
// ###########################.
func TestSerializeTagHierarchyNoTags(t *testing.T) {
	_, err := libtags.SerializeTagHierarchy(libtags.TagNode{})
	assert.ErrorAs(t, err, &helpers.SerializationError{})
}

func TestSerializeTagHierarchyOnlyTopLevel(t *testing.T) {
	input := `
tags:
- bar
- baz
- foo
    `

	data := libtags.TagNode{
		"tags": libtags.TagNode{
			"bar": nil,
			"baz": nil,
			"foo": nil,
		},
	}

	output, err := libtags.SerializeTagHierarchy(data)
	assert.NoError(t, err)
	assert.YAMLEq(t, input, output)
}

func TestSerializeTagHierarchyOnePath(t *testing.T) {
	input := `
tags:
- foo:
    - bar:
        - baz
    `

	data := libtags.TagNode{
		"tags": libtags.TagNode{
			"foo": libtags.TagNode{
				"bar": libtags.TagNode{
					"baz": nil,
				},
			},
		},
	}

	output, err := libtags.SerializeTagHierarchy(data)
	assert.NoError(t, err)
	assert.YAMLEq(t, input, output)
}

func TestSerializeTagHierarchyTwoPaths(t *testing.T) {
	input := `
tags:
- foo:
    - bar:
        - baz
- foo2:
    - bar2:
        - baz2
    `

	data := libtags.TagNode{
		"tags": libtags.TagNode{
			"foo": libtags.TagNode{
				"bar": libtags.TagNode{
					"baz": nil,
				},
			},
			"foo2": libtags.TagNode{
				"bar2": libtags.TagNode{
					"baz2": nil,
				},
			},
		},
	}

	output, err := libtags.SerializeTagHierarchy(data)
	assert.NoError(t, err)
	assert.YAMLEq(t, input, output)
}
