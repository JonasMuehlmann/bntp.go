// Code generated by SQLBoiler 4.10.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testBookmarkTypes(t *testing.T) {
	

	query := BookmarkTypes()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testBookmarkTypesDelete(t *testing.T) {
	

	seed := randomize.NewSeed()
	var err error
	o := &BookmarkType{}
	if err = randomize.Struct(seed, o, bookmarkTypeDBTypes, true, bookmarkTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookmarkType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := BookmarkTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBookmarkTypesQueryDeleteAll(t *testing.T) {
	

	seed := randomize.NewSeed()
	var err error
	o := &BookmarkType{}
	if err = randomize.Struct(seed, o, bookmarkTypeDBTypes, true, bookmarkTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookmarkType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := BookmarkTypes().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := BookmarkTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBookmarkTypesSliceDeleteAll(t *testing.T) {
	

	seed := randomize.NewSeed()
	var err error
	o := &BookmarkType{}
	if err = randomize.Struct(seed, o, bookmarkTypeDBTypes, true, bookmarkTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookmarkType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := BookmarkTypeSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := BookmarkTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBookmarkTypesExists(t *testing.T) {
	

	seed := randomize.NewSeed()
	var err error
	o := &BookmarkType{}
	if err = randomize.Struct(seed, o, bookmarkTypeDBTypes, true, bookmarkTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookmarkType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := BookmarkTypeExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if BookmarkType exists: %s", err)
	}
	if !e {
		t.Errorf("Expected BookmarkTypeExists to return true, but got false.")
	}
}

func testBookmarkTypesFind(t *testing.T) {
	

	seed := randomize.NewSeed()
	var err error
	o := &BookmarkType{}
	if err = randomize.Struct(seed, o, bookmarkTypeDBTypes, true, bookmarkTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookmarkType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	bookmarkTypeFound, err := FindBookmarkType(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if bookmarkTypeFound == nil {
		t.Error("want a record, got nil")
	}
}

func testBookmarkTypesBind(t *testing.T) {
	

	seed := randomize.NewSeed()
	var err error
	o := &BookmarkType{}
	if err = randomize.Struct(seed, o, bookmarkTypeDBTypes, true, bookmarkTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookmarkType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = BookmarkTypes().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testBookmarkTypesOne(t *testing.T) {
	

	seed := randomize.NewSeed()
	var err error
	o := &BookmarkType{}
	if err = randomize.Struct(seed, o, bookmarkTypeDBTypes, true, bookmarkTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookmarkType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := BookmarkTypes().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testBookmarkTypesAll(t *testing.T) {
	

	seed := randomize.NewSeed()
	var err error
	bookmarkTypeOne := &BookmarkType{}
	bookmarkTypeTwo := &BookmarkType{}
	if err = randomize.Struct(seed, bookmarkTypeOne, bookmarkTypeDBTypes, false, bookmarkTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookmarkType struct: %s", err)
	}
	if err = randomize.Struct(seed, bookmarkTypeTwo, bookmarkTypeDBTypes, false, bookmarkTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookmarkType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = bookmarkTypeOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = bookmarkTypeTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := BookmarkTypes().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testBookmarkTypesCount(t *testing.T) {
	

	var err error
	seed := randomize.NewSeed()
	bookmarkTypeOne := &BookmarkType{}
	bookmarkTypeTwo := &BookmarkType{}
	if err = randomize.Struct(seed, bookmarkTypeOne, bookmarkTypeDBTypes, false, bookmarkTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookmarkType struct: %s", err)
	}
	if err = randomize.Struct(seed, bookmarkTypeTwo, bookmarkTypeDBTypes, false, bookmarkTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookmarkType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = bookmarkTypeOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = bookmarkTypeTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BookmarkTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func bookmarkTypeBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *BookmarkType) error {
	*o = BookmarkType{}
	return nil
}

func bookmarkTypeAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *BookmarkType) error {
	*o = BookmarkType{}
	return nil
}

func bookmarkTypeAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *BookmarkType) error {
	*o = BookmarkType{}
	return nil
}

func bookmarkTypeBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *BookmarkType) error {
	*o = BookmarkType{}
	return nil
}

func bookmarkTypeAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *BookmarkType) error {
	*o = BookmarkType{}
	return nil
}

func bookmarkTypeBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *BookmarkType) error {
	*o = BookmarkType{}
	return nil
}

func bookmarkTypeAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *BookmarkType) error {
	*o = BookmarkType{}
	return nil
}

func bookmarkTypeBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *BookmarkType) error {
	*o = BookmarkType{}
	return nil
}

func bookmarkTypeAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *BookmarkType) error {
	*o = BookmarkType{}
	return nil
}

func testBookmarkTypesHooks(t *testing.T) {
	

	var err error

	ctx := context.Background()
	empty := &BookmarkType{}
	o := &BookmarkType{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, bookmarkTypeDBTypes, false); err != nil {
		t.Errorf("Unable to randomize BookmarkType object: %s", err)
	}

	AddBookmarkTypeHook(boil.BeforeInsertHook, bookmarkTypeBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	bookmarkTypeBeforeInsertHooks = []BookmarkTypeHook{}

	AddBookmarkTypeHook(boil.AfterInsertHook, bookmarkTypeAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	bookmarkTypeAfterInsertHooks = []BookmarkTypeHook{}

	AddBookmarkTypeHook(boil.AfterSelectHook, bookmarkTypeAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	bookmarkTypeAfterSelectHooks = []BookmarkTypeHook{}

	AddBookmarkTypeHook(boil.BeforeUpdateHook, bookmarkTypeBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	bookmarkTypeBeforeUpdateHooks = []BookmarkTypeHook{}

	AddBookmarkTypeHook(boil.AfterUpdateHook, bookmarkTypeAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	bookmarkTypeAfterUpdateHooks = []BookmarkTypeHook{}

	AddBookmarkTypeHook(boil.BeforeDeleteHook, bookmarkTypeBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	bookmarkTypeBeforeDeleteHooks = []BookmarkTypeHook{}

	AddBookmarkTypeHook(boil.AfterDeleteHook, bookmarkTypeAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	bookmarkTypeAfterDeleteHooks = []BookmarkTypeHook{}

	AddBookmarkTypeHook(boil.BeforeUpsertHook, bookmarkTypeBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	bookmarkTypeBeforeUpsertHooks = []BookmarkTypeHook{}

	AddBookmarkTypeHook(boil.AfterUpsertHook, bookmarkTypeAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	bookmarkTypeAfterUpsertHooks = []BookmarkTypeHook{}
}

func testBookmarkTypesInsert(t *testing.T) {
	

	seed := randomize.NewSeed()
	var err error
	o := &BookmarkType{}
	if err = randomize.Struct(seed, o, bookmarkTypeDBTypes, true, bookmarkTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookmarkType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BookmarkTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBookmarkTypesInsertWhitelist(t *testing.T) {
	

	seed := randomize.NewSeed()
	var err error
	o := &BookmarkType{}
	if err = randomize.Struct(seed, o, bookmarkTypeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize BookmarkType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(bookmarkTypeColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := BookmarkTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBookmarkTypeToManyBookmarks(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a BookmarkType
	var b, c Bookmark

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, bookmarkTypeDBTypes, true, bookmarkTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookmarkType struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, bookmarkDBTypes, false, bookmarkColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, bookmarkDBTypes, false, bookmarkColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&b.BookmarkTypeID, a.ID)
	queries.Assign(&c.BookmarkTypeID, a.ID)
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.Bookmarks().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if queries.Equal(v.BookmarkTypeID, b.BookmarkTypeID) {
			bFound = true
		}
		if queries.Equal(v.BookmarkTypeID, c.BookmarkTypeID) {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := BookmarkTypeSlice{&a}
	if err = a.L.LoadBookmarks(ctx, tx, false, (*[]*BookmarkType)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Bookmarks); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Bookmarks = nil
	if err = a.L.LoadBookmarks(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Bookmarks); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testBookmarkTypeToManyAddOpBookmarks(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a BookmarkType
	var b, c, d, e Bookmark

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, bookmarkTypeDBTypes, false, strmangle.SetComplement(bookmarkTypePrimaryKeyColumns, bookmarkTypeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Bookmark{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, bookmarkDBTypes, false, strmangle.SetComplement(bookmarkPrimaryKeyColumns, bookmarkColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Bookmark{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddBookmarks(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if !queries.Equal(a.ID, first.BookmarkTypeID) {
			t.Error("foreign key was wrong value", a.ID, first.BookmarkTypeID)
		}
		if !queries.Equal(a.ID, second.BookmarkTypeID) {
			t.Error("foreign key was wrong value", a.ID, second.BookmarkTypeID)
		}

		if first.R.BookmarkType != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.BookmarkType != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Bookmarks[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Bookmarks[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Bookmarks().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testBookmarkTypeToManySetOpBookmarks(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a BookmarkType
	var b, c, d, e Bookmark

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, bookmarkTypeDBTypes, false, strmangle.SetComplement(bookmarkTypePrimaryKeyColumns, bookmarkTypeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Bookmark{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, bookmarkDBTypes, false, strmangle.SetComplement(bookmarkPrimaryKeyColumns, bookmarkColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.SetBookmarks(ctx, tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Bookmarks().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetBookmarks(ctx, tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Bookmarks().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.BookmarkTypeID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.BookmarkTypeID) {
		t.Error("want c's foreign key value to be nil")
	}
	if !queries.Equal(a.ID, d.BookmarkTypeID) {
		t.Error("foreign key was wrong value", a.ID, d.BookmarkTypeID)
	}
	if !queries.Equal(a.ID, e.BookmarkTypeID) {
		t.Error("foreign key was wrong value", a.ID, e.BookmarkTypeID)
	}

	if b.R.BookmarkType != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.BookmarkType != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.BookmarkType != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.BookmarkType != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.Bookmarks[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.Bookmarks[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testBookmarkTypeToManyRemoveOpBookmarks(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a BookmarkType
	var b, c, d, e Bookmark

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, bookmarkTypeDBTypes, false, strmangle.SetComplement(bookmarkTypePrimaryKeyColumns, bookmarkTypeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Bookmark{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, bookmarkDBTypes, false, strmangle.SetComplement(bookmarkPrimaryKeyColumns, bookmarkColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.AddBookmarks(ctx, tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Bookmarks().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveBookmarks(ctx, tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Bookmarks().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.BookmarkTypeID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.BookmarkTypeID) {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.BookmarkType != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.BookmarkType != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.BookmarkType != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.BookmarkType != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.Bookmarks) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.Bookmarks[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.Bookmarks[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testBookmarkTypesReload(t *testing.T) {
	

	seed := randomize.NewSeed()
	var err error
	o := &BookmarkType{}
	if err = randomize.Struct(seed, o, bookmarkTypeDBTypes, true, bookmarkTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookmarkType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testBookmarkTypesReloadAll(t *testing.T) {
	

	seed := randomize.NewSeed()
	var err error
	o := &BookmarkType{}
	if err = randomize.Struct(seed, o, bookmarkTypeDBTypes, true, bookmarkTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookmarkType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := BookmarkTypeSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testBookmarkTypesSelect(t *testing.T) {
	

	seed := randomize.NewSeed()
	var err error
	o := &BookmarkType{}
	if err = randomize.Struct(seed, o, bookmarkTypeDBTypes, true, bookmarkTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookmarkType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := BookmarkTypes().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	bookmarkTypeDBTypes = map[string]string{`ID`: `INTEGER`, `Type`: `TEXT`}
	_                   = bytes.MinRead
)

func testBookmarkTypesUpdate(t *testing.T) {
	

	if 0 == len(bookmarkTypePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(bookmarkTypeAllColumns) == len(bookmarkTypePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &BookmarkType{}
	if err = randomize.Struct(seed, o, bookmarkTypeDBTypes, true, bookmarkTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookmarkType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BookmarkTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, bookmarkTypeDBTypes, true, bookmarkTypePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize BookmarkType struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testBookmarkTypesSliceUpdateAll(t *testing.T) {
	

	if len(bookmarkTypeAllColumns) == len(bookmarkTypePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &BookmarkType{}
	if err = randomize.Struct(seed, o, bookmarkTypeDBTypes, true, bookmarkTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookmarkType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BookmarkTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, bookmarkTypeDBTypes, true, bookmarkTypePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize BookmarkType struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(bookmarkTypeAllColumns, bookmarkTypePrimaryKeyColumns) {
		fields = bookmarkTypeAllColumns
	} else {
		fields = strmangle.SetComplement(
			bookmarkTypeAllColumns,
			bookmarkTypePrimaryKeyColumns,
		)
		fields = strmangle.SetComplement(fields, bookmarkTypeGeneratedColumns)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := BookmarkTypeSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testBookmarkTypesUpsert(t *testing.T) {
	
	if len(bookmarkTypeAllColumns) == len(bookmarkTypePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := BookmarkType{}
	if err = randomize.Struct(seed, &o, bookmarkTypeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize BookmarkType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert BookmarkType: %s", err)
	}

	count, err := BookmarkTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, bookmarkTypeDBTypes, false, bookmarkTypePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize BookmarkType struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert BookmarkType: %s", err)
	}

	count, err = BookmarkTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
