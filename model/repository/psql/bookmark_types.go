// Code generated by SQLBoiler 4.11.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package repository

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// BookmarkType is an object representing the database table.
type BookmarkType struct {
	ID           int64  `boil:"id" json:"id" toml:"id" yaml:"id"`
	BookmarkType string `boil:"bookmark_type" json:"bookmark_type" toml:"bookmark_type" yaml:"bookmark_type"`

	R *bookmarkTypeR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L bookmarkTypeL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var BookmarkTypeColumns = struct {
	ID           string
	BookmarkType string
}{
	ID:           "id",
	BookmarkType: "bookmark_type",
}

var BookmarkTypeTableColumns = struct {
	ID           string
	BookmarkType string
}{
	ID:           "bookmark_types.id",
	BookmarkType: "bookmark_types.bookmark_type",
}

// Generated where

type whereHelperint64 struct{ field string }

func (w whereHelperint64) EQ(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint64) NEQ(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint64) LT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint64) LTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint64) GT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint64) GTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint64) IN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint64) NIN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

var BookmarkTypeWhere = struct {
	ID           whereHelperint64
	BookmarkType whereHelperstring
}{
	ID:           whereHelperint64{field: "\"bookmark_types\".\"id\""},
	BookmarkType: whereHelperstring{field: "\"bookmark_types\".\"bookmark_type\""},
}

// BookmarkTypeRels is where relationship names are stored.
var BookmarkTypeRels = struct {
	Bookmarks string
}{
	Bookmarks: "Bookmarks",
}

// bookmarkTypeR is where relationships are stored.
type bookmarkTypeR struct {
	Bookmarks BookmarkSlice `boil:"Bookmarks" json:"Bookmarks" toml:"Bookmarks" yaml:"Bookmarks"`
}

// NewStruct creates a new relationship struct
func (*bookmarkTypeR) NewStruct() *bookmarkTypeR {
	return &bookmarkTypeR{}
}

func (r *bookmarkTypeR) GetBookmarks() BookmarkSlice {
	if r == nil {
		return nil
	}
	return r.Bookmarks
}

// bookmarkTypeL is where Load methods for each relationship are stored.
type bookmarkTypeL struct{}

var (
	bookmarkTypeAllColumns            = []string{"id", "bookmark_type"}
	bookmarkTypeColumnsWithoutDefault = []string{"id", "bookmark_type"}
	bookmarkTypeColumnsWithDefault    = []string{}
	bookmarkTypePrimaryKeyColumns     = []string{"id"}
	bookmarkTypeGeneratedColumns      = []string{}
)

type (
	// BookmarkTypeSlice is an alias for a slice of pointers to BookmarkType.
	// This should almost always be used instead of []BookmarkType.
	BookmarkTypeSlice []*BookmarkType
	// BookmarkTypeHook is the signature for custom BookmarkType hook methods
	BookmarkTypeHook func(context.Context, boil.ContextExecutor, *BookmarkType) error

	bookmarkTypeQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	bookmarkTypeType                 = reflect.TypeOf(&BookmarkType{})
	bookmarkTypeMapping              = queries.MakeStructMapping(bookmarkTypeType)
	bookmarkTypePrimaryKeyMapping, _ = queries.BindMapping(bookmarkTypeType, bookmarkTypeMapping, bookmarkTypePrimaryKeyColumns)
	bookmarkTypeInsertCacheMut       sync.RWMutex
	bookmarkTypeInsertCache          = make(map[string]insertCache)
	bookmarkTypeUpdateCacheMut       sync.RWMutex
	bookmarkTypeUpdateCache          = make(map[string]updateCache)
	bookmarkTypeUpsertCacheMut       sync.RWMutex
	bookmarkTypeUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var bookmarkTypeAfterSelectHooks []BookmarkTypeHook

var bookmarkTypeBeforeInsertHooks []BookmarkTypeHook
var bookmarkTypeAfterInsertHooks []BookmarkTypeHook

var bookmarkTypeBeforeUpdateHooks []BookmarkTypeHook
var bookmarkTypeAfterUpdateHooks []BookmarkTypeHook

var bookmarkTypeBeforeDeleteHooks []BookmarkTypeHook
var bookmarkTypeAfterDeleteHooks []BookmarkTypeHook

var bookmarkTypeBeforeUpsertHooks []BookmarkTypeHook
var bookmarkTypeAfterUpsertHooks []BookmarkTypeHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *BookmarkType) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookmarkTypeAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *BookmarkType) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookmarkTypeBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *BookmarkType) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookmarkTypeAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *BookmarkType) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookmarkTypeBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *BookmarkType) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookmarkTypeAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *BookmarkType) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookmarkTypeBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *BookmarkType) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookmarkTypeAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *BookmarkType) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookmarkTypeBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *BookmarkType) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookmarkTypeAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddBookmarkTypeHook registers your hook function for all future operations.
func AddBookmarkTypeHook(hookPoint boil.HookPoint, bookmarkTypeHook BookmarkTypeHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		bookmarkTypeAfterSelectHooks = append(bookmarkTypeAfterSelectHooks, bookmarkTypeHook)
	case boil.BeforeInsertHook:
		bookmarkTypeBeforeInsertHooks = append(bookmarkTypeBeforeInsertHooks, bookmarkTypeHook)
	case boil.AfterInsertHook:
		bookmarkTypeAfterInsertHooks = append(bookmarkTypeAfterInsertHooks, bookmarkTypeHook)
	case boil.BeforeUpdateHook:
		bookmarkTypeBeforeUpdateHooks = append(bookmarkTypeBeforeUpdateHooks, bookmarkTypeHook)
	case boil.AfterUpdateHook:
		bookmarkTypeAfterUpdateHooks = append(bookmarkTypeAfterUpdateHooks, bookmarkTypeHook)
	case boil.BeforeDeleteHook:
		bookmarkTypeBeforeDeleteHooks = append(bookmarkTypeBeforeDeleteHooks, bookmarkTypeHook)
	case boil.AfterDeleteHook:
		bookmarkTypeAfterDeleteHooks = append(bookmarkTypeAfterDeleteHooks, bookmarkTypeHook)
	case boil.BeforeUpsertHook:
		bookmarkTypeBeforeUpsertHooks = append(bookmarkTypeBeforeUpsertHooks, bookmarkTypeHook)
	case boil.AfterUpsertHook:
		bookmarkTypeAfterUpsertHooks = append(bookmarkTypeAfterUpsertHooks, bookmarkTypeHook)
	}
}

// One returns a single bookmarkType record from the query.
func (q bookmarkTypeQuery) One(ctx context.Context, exec boil.ContextExecutor) (*BookmarkType, error) {
	o := &BookmarkType{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "repository: failed to execute a one query for bookmark_types")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all BookmarkType records from the query.
func (q bookmarkTypeQuery) All(ctx context.Context, exec boil.ContextExecutor) (BookmarkTypeSlice, error) {
	var o []*BookmarkType

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "repository: failed to assign all query results to BookmarkType slice")
	}

	if len(bookmarkTypeAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all BookmarkType records in the query.
func (q bookmarkTypeQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "repository: failed to count bookmark_types rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q bookmarkTypeQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "repository: failed to check if bookmark_types exists")
	}

	return count > 0, nil
}

// Bookmarks retrieves all the bookmark's Bookmarks with an executor.
func (o *BookmarkType) Bookmarks(mods ...qm.QueryMod) bookmarkQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"bookmarks\".\"bookmark_type_id\"=?", o.ID),
	)

	return Bookmarks(queryMods...)
}

// LoadBookmarks allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (bookmarkTypeL) LoadBookmarks(ctx context.Context, e boil.ContextExecutor, singular bool, maybeBookmarkType interface{}, mods queries.Applicator) error {
	var slice []*BookmarkType
	var object *BookmarkType

	if singular {
		object = maybeBookmarkType.(*BookmarkType)
	} else {
		slice = *maybeBookmarkType.(*[]*BookmarkType)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &bookmarkTypeR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &bookmarkTypeR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.ID) {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`bookmarks`),
		qm.WhereIn(`bookmarks.bookmark_type_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load bookmarks")
	}

	var resultSlice []*Bookmark
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice bookmarks")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on bookmarks")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for bookmarks")
	}

	if len(bookmarkAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Bookmarks = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &bookmarkR{}
			}
			foreign.R.BookmarkType = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if queries.Equal(local.ID, foreign.BookmarkTypeID) {
				local.R.Bookmarks = append(local.R.Bookmarks, foreign)
				if foreign.R == nil {
					foreign.R = &bookmarkR{}
				}
				foreign.R.BookmarkType = local
				break
			}
		}
	}

	return nil
}

// AddBookmarks adds the given related objects to the existing relationships
// of the bookmark_type, optionally inserting them as new records.
// Appends related to o.R.Bookmarks.
// Sets related.R.BookmarkType appropriately.
func (o *BookmarkType) AddBookmarks(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Bookmark) error {
	var err error
	for _, rel := range related {
		if insert {
			queries.Assign(&rel.BookmarkTypeID, o.ID)
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"bookmarks\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"bookmark_type_id"}),
				strmangle.WhereClause("\"", "\"", 2, bookmarkPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			queries.Assign(&rel.BookmarkTypeID, o.ID)
		}
	}

	if o.R == nil {
		o.R = &bookmarkTypeR{
			Bookmarks: related,
		}
	} else {
		o.R.Bookmarks = append(o.R.Bookmarks, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &bookmarkR{
				BookmarkType: o,
			}
		} else {
			rel.R.BookmarkType = o
		}
	}
	return nil
}

// SetBookmarks removes all previously related items of the
// bookmark_type replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.BookmarkType's Bookmarks accordingly.
// Replaces o.R.Bookmarks with related.
// Sets related.R.BookmarkType's Bookmarks accordingly.
func (o *BookmarkType) SetBookmarks(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Bookmark) error {
	query := "update \"bookmarks\" set \"bookmark_type_id\" = null where \"bookmark_type_id\" = $1"
	values := []interface{}{o.ID}
	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, query)
		fmt.Fprintln(writer, values)
	}
	_, err := exec.ExecContext(ctx, query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}

	if o.R != nil {
		for _, rel := range o.R.Bookmarks {
			queries.SetScanner(&rel.BookmarkTypeID, nil)
			if rel.R == nil {
				continue
			}

			rel.R.BookmarkType = nil
		}
		o.R.Bookmarks = nil
	}

	return o.AddBookmarks(ctx, exec, insert, related...)
}

// RemoveBookmarks relationships from objects passed in.
// Removes related items from R.Bookmarks (uses pointer comparison, removal does not keep order)
// Sets related.R.BookmarkType.
func (o *BookmarkType) RemoveBookmarks(ctx context.Context, exec boil.ContextExecutor, related ...*Bookmark) error {
	if len(related) == 0 {
		return nil
	}

	var err error
	for _, rel := range related {
		queries.SetScanner(&rel.BookmarkTypeID, nil)
		if rel.R != nil {
			rel.R.BookmarkType = nil
		}
		if _, err = rel.Update(ctx, exec, boil.Whitelist("bookmark_type_id")); err != nil {
			return err
		}
	}
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.Bookmarks {
			if rel != ri {
				continue
			}

			ln := len(o.R.Bookmarks)
			if ln > 1 && i < ln-1 {
				o.R.Bookmarks[i] = o.R.Bookmarks[ln-1]
			}
			o.R.Bookmarks = o.R.Bookmarks[:ln-1]
			break
		}
	}

	return nil
}

// BookmarkTypes retrieves all the records using an executor.
func BookmarkTypes(mods ...qm.QueryMod) bookmarkTypeQuery {
	mods = append(mods, qm.From("\"bookmark_types\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"bookmark_types\".*"})
	}

	return bookmarkTypeQuery{q}
}

// FindBookmarkType retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindBookmarkType(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*BookmarkType, error) {
	bookmarkTypeObj := &BookmarkType{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"bookmark_types\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, bookmarkTypeObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "repository: unable to select from bookmark_types")
	}

	if err = bookmarkTypeObj.doAfterSelectHooks(ctx, exec); err != nil {
		return bookmarkTypeObj, err
	}

	return bookmarkTypeObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *BookmarkType) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("repository: no bookmark_types provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(bookmarkTypeColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	bookmarkTypeInsertCacheMut.RLock()
	cache, cached := bookmarkTypeInsertCache[key]
	bookmarkTypeInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			bookmarkTypeAllColumns,
			bookmarkTypeColumnsWithDefault,
			bookmarkTypeColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(bookmarkTypeType, bookmarkTypeMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(bookmarkTypeType, bookmarkTypeMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"bookmark_types\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"bookmark_types\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "repository: unable to insert into bookmark_types")
	}

	if !cached {
		bookmarkTypeInsertCacheMut.Lock()
		bookmarkTypeInsertCache[key] = cache
		bookmarkTypeInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the BookmarkType.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *BookmarkType) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	bookmarkTypeUpdateCacheMut.RLock()
	cache, cached := bookmarkTypeUpdateCache[key]
	bookmarkTypeUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			bookmarkTypeAllColumns,
			bookmarkTypePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("repository: unable to update bookmark_types, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"bookmark_types\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, bookmarkTypePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(bookmarkTypeType, bookmarkTypeMapping, append(wl, bookmarkTypePrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "repository: unable to update bookmark_types row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "repository: failed to get rows affected by update for bookmark_types")
	}

	if !cached {
		bookmarkTypeUpdateCacheMut.Lock()
		bookmarkTypeUpdateCache[key] = cache
		bookmarkTypeUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q bookmarkTypeQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "repository: unable to update all for bookmark_types")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "repository: unable to retrieve rows affected for bookmark_types")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o BookmarkTypeSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("repository: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), bookmarkTypePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"bookmark_types\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, bookmarkTypePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "repository: unable to update all in bookmarkType slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "repository: unable to retrieve rows affected all in update all bookmarkType")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *BookmarkType) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("repository: no bookmark_types provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(bookmarkTypeColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	bookmarkTypeUpsertCacheMut.RLock()
	cache, cached := bookmarkTypeUpsertCache[key]
	bookmarkTypeUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			bookmarkTypeAllColumns,
			bookmarkTypeColumnsWithDefault,
			bookmarkTypeColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			bookmarkTypeAllColumns,
			bookmarkTypePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("repository: unable to upsert bookmark_types, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(bookmarkTypePrimaryKeyColumns))
			copy(conflict, bookmarkTypePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"bookmark_types\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(bookmarkTypeType, bookmarkTypeMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(bookmarkTypeType, bookmarkTypeMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "repository: unable to upsert bookmark_types")
	}

	if !cached {
		bookmarkTypeUpsertCacheMut.Lock()
		bookmarkTypeUpsertCache[key] = cache
		bookmarkTypeUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single BookmarkType record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *BookmarkType) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("repository: no BookmarkType provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), bookmarkTypePrimaryKeyMapping)
	sql := "DELETE FROM \"bookmark_types\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "repository: unable to delete from bookmark_types")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "repository: failed to get rows affected by delete for bookmark_types")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q bookmarkTypeQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("repository: no bookmarkTypeQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "repository: unable to delete all from bookmark_types")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "repository: failed to get rows affected by deleteall for bookmark_types")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o BookmarkTypeSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(bookmarkTypeBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), bookmarkTypePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"bookmark_types\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, bookmarkTypePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "repository: unable to delete all from bookmarkType slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "repository: failed to get rows affected by deleteall for bookmark_types")
	}

	if len(bookmarkTypeAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *BookmarkType) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindBookmarkType(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *BookmarkTypeSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := BookmarkTypeSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), bookmarkTypePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"bookmark_types\".* FROM \"bookmark_types\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, bookmarkTypePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "repository: unable to reload all in BookmarkTypeSlice")
	}

	*o = slice

	return nil
}

// BookmarkTypeExists checks if the BookmarkType row exists.
func BookmarkTypeExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"bookmark_types\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "repository: unable to check if bookmark_types exists")
	}

	return exists, nil
}
