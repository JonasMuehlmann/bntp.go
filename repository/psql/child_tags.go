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
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// ChildTag is an object representing the database table.
type ChildTag struct {
	ID          int      `boil:"id" json:"id" toml:"id" yaml:"id"`
	ParentTagID null.Int `boil:"parent_tag_id" json:"parent_tag_id,omitempty" toml:"parent_tag_id" yaml:"parent_tag_id,omitempty"`
	ChildTagID  null.Int `boil:"child_tag_id" json:"child_tag_id,omitempty" toml:"child_tag_id" yaml:"child_tag_id,omitempty"`

	R *childTagR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L childTagL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ChildTagColumns = struct {
	ID          string
	ParentTagID string
	ChildTagID  string
}{
	ID:          "id",
	ParentTagID: "parent_tag_id",
	ChildTagID:  "child_tag_id",
}

var ChildTagTableColumns = struct {
	ID          string
	ParentTagID string
	ChildTagID  string
}{
	ID:          "child_tags.id",
	ParentTagID: "child_tags.parent_tag_id",
	ChildTagID:  "child_tags.child_tag_id",
}

// Generated where

var ChildTagWhere = struct {
	ID          whereHelperint
	ParentTagID whereHelpernull_Int
	ChildTagID  whereHelpernull_Int
}{
	ID:          whereHelperint{field: "\"child_tags\".\"id\""},
	ParentTagID: whereHelpernull_Int{field: "\"child_tags\".\"parent_tag_id\""},
	ChildTagID:  whereHelpernull_Int{field: "\"child_tags\".\"child_tag_id\""},
}

// ChildTagRels is where relationship names are stored.
var ChildTagRels = struct {
	ChildTag  string
	ParentTag string
}{
	ChildTag:  "ChildTag",
	ParentTag: "ParentTag",
}

// childTagR is where relationships are stored.
type childTagR struct {
	ChildTag  *Tag `boil:"ChildTag" json:"ChildTag" toml:"ChildTag" yaml:"ChildTag"`
	ParentTag *Tag `boil:"ParentTag" json:"ParentTag" toml:"ParentTag" yaml:"ParentTag"`
}

// NewStruct creates a new relationship struct
func (*childTagR) NewStruct() *childTagR {
	return &childTagR{}
}

func (r *childTagR) GetChildTag() *Tag {
	if r == nil {
		return nil
	}
	return r.ChildTag
}

func (r *childTagR) GetParentTag() *Tag {
	if r == nil {
		return nil
	}
	return r.ParentTag
}

// childTagL is where Load methods for each relationship are stored.
type childTagL struct{}

var (
	childTagAllColumns            = []string{"id", "parent_tag_id", "child_tag_id"}
	childTagColumnsWithoutDefault = []string{"id"}
	childTagColumnsWithDefault    = []string{"parent_tag_id", "child_tag_id"}
	childTagPrimaryKeyColumns     = []string{"id"}
	childTagGeneratedColumns      = []string{}
)

type (
	// ChildTagSlice is an alias for a slice of pointers to ChildTag.
	// This should almost always be used instead of []ChildTag.
	ChildTagSlice []*ChildTag
	// ChildTagHook is the signature for custom ChildTag hook methods
	ChildTagHook func(context.Context, boil.ContextExecutor, *ChildTag) error

	childTagQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	childTagType                 = reflect.TypeOf(&ChildTag{})
	childTagMapping              = queries.MakeStructMapping(childTagType)
	childTagPrimaryKeyMapping, _ = queries.BindMapping(childTagType, childTagMapping, childTagPrimaryKeyColumns)
	childTagInsertCacheMut       sync.RWMutex
	childTagInsertCache          = make(map[string]insertCache)
	childTagUpdateCacheMut       sync.RWMutex
	childTagUpdateCache          = make(map[string]updateCache)
	childTagUpsertCacheMut       sync.RWMutex
	childTagUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var childTagAfterSelectHooks []ChildTagHook

var childTagBeforeInsertHooks []ChildTagHook
var childTagAfterInsertHooks []ChildTagHook

var childTagBeforeUpdateHooks []ChildTagHook
var childTagAfterUpdateHooks []ChildTagHook

var childTagBeforeDeleteHooks []ChildTagHook
var childTagAfterDeleteHooks []ChildTagHook

var childTagBeforeUpsertHooks []ChildTagHook
var childTagAfterUpsertHooks []ChildTagHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *ChildTag) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range childTagAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *ChildTag) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range childTagBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *ChildTag) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range childTagAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *ChildTag) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range childTagBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *ChildTag) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range childTagAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *ChildTag) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range childTagBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *ChildTag) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range childTagAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *ChildTag) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range childTagBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *ChildTag) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range childTagAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddChildTagHook registers your hook function for all future operations.
func AddChildTagHook(hookPoint boil.HookPoint, childTagHook ChildTagHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		childTagAfterSelectHooks = append(childTagAfterSelectHooks, childTagHook)
	case boil.BeforeInsertHook:
		childTagBeforeInsertHooks = append(childTagBeforeInsertHooks, childTagHook)
	case boil.AfterInsertHook:
		childTagAfterInsertHooks = append(childTagAfterInsertHooks, childTagHook)
	case boil.BeforeUpdateHook:
		childTagBeforeUpdateHooks = append(childTagBeforeUpdateHooks, childTagHook)
	case boil.AfterUpdateHook:
		childTagAfterUpdateHooks = append(childTagAfterUpdateHooks, childTagHook)
	case boil.BeforeDeleteHook:
		childTagBeforeDeleteHooks = append(childTagBeforeDeleteHooks, childTagHook)
	case boil.AfterDeleteHook:
		childTagAfterDeleteHooks = append(childTagAfterDeleteHooks, childTagHook)
	case boil.BeforeUpsertHook:
		childTagBeforeUpsertHooks = append(childTagBeforeUpsertHooks, childTagHook)
	case boil.AfterUpsertHook:
		childTagAfterUpsertHooks = append(childTagAfterUpsertHooks, childTagHook)
	}
}

// One returns a single childTag record from the query.
func (q childTagQuery) One(ctx context.Context, exec boil.ContextExecutor) (*ChildTag, error) {
	o := &ChildTag{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for child_tags")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all ChildTag records from the query.
func (q childTagQuery) All(ctx context.Context, exec boil.ContextExecutor) (ChildTagSlice, error) {
	var o []*ChildTag

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to ChildTag slice")
	}

	if len(childTagAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all ChildTag records in the query.
func (q childTagQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count child_tags rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q childTagQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if child_tags exists")
	}

	return count > 0, nil
}

// ChildTag pointed to by the foreign key.
func (o *ChildTag) ChildTag(mods ...qm.QueryMod) tagQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.ChildTagID),
	}

	queryMods = append(queryMods, mods...)

	return Tags(queryMods...)
}

// ParentTag pointed to by the foreign key.
func (o *ChildTag) ParentTag(mods ...qm.QueryMod) tagQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.ParentTagID),
	}

	queryMods = append(queryMods, mods...)

	return Tags(queryMods...)
}

// LoadChildTag allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (childTagL) LoadChildTag(ctx context.Context, e boil.ContextExecutor, singular bool, maybeChildTag interface{}, mods queries.Applicator) error {
	var slice []*ChildTag
	var object *ChildTag

	if singular {
		object = maybeChildTag.(*ChildTag)
	} else {
		slice = *maybeChildTag.(*[]*ChildTag)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &childTagR{}
		}
		if !queries.IsNil(object.ChildTagID) {
			args = append(args, object.ChildTagID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &childTagR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.ChildTagID) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.ChildTagID) {
				args = append(args, obj.ChildTagID)
			}

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`tags`),
		qm.WhereIn(`tags.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Tag")
	}

	var resultSlice []*Tag
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Tag")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for tags")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for tags")
	}

	if len(childTagAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.ChildTag = foreign
		if foreign.R == nil {
			foreign.R = &tagR{}
		}
		foreign.R.ChildTagChildTags = append(foreign.R.ChildTagChildTags, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.ChildTagID, foreign.ID) {
				local.R.ChildTag = foreign
				if foreign.R == nil {
					foreign.R = &tagR{}
				}
				foreign.R.ChildTagChildTags = append(foreign.R.ChildTagChildTags, local)
				break
			}
		}
	}

	return nil
}

// LoadParentTag allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (childTagL) LoadParentTag(ctx context.Context, e boil.ContextExecutor, singular bool, maybeChildTag interface{}, mods queries.Applicator) error {
	var slice []*ChildTag
	var object *ChildTag

	if singular {
		object = maybeChildTag.(*ChildTag)
	} else {
		slice = *maybeChildTag.(*[]*ChildTag)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &childTagR{}
		}
		if !queries.IsNil(object.ParentTagID) {
			args = append(args, object.ParentTagID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &childTagR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.ParentTagID) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.ParentTagID) {
				args = append(args, obj.ParentTagID)
			}

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`tags`),
		qm.WhereIn(`tags.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Tag")
	}

	var resultSlice []*Tag
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Tag")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for tags")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for tags")
	}

	if len(childTagAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.ParentTag = foreign
		if foreign.R == nil {
			foreign.R = &tagR{}
		}
		foreign.R.ParentTagChildTags = append(foreign.R.ParentTagChildTags, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.ParentTagID, foreign.ID) {
				local.R.ParentTag = foreign
				if foreign.R == nil {
					foreign.R = &tagR{}
				}
				foreign.R.ParentTagChildTags = append(foreign.R.ParentTagChildTags, local)
				break
			}
		}
	}

	return nil
}

// SetChildTag of the childTag to the related item.
// Sets o.R.ChildTag to related.
// Adds o to related.R.ChildTagChildTags.
func (o *ChildTag) SetChildTag(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Tag) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"child_tags\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"child_tag_id"}),
		strmangle.WhereClause("\"", "\"", 2, childTagPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	queries.Assign(&o.ChildTagID, related.ID)
	if o.R == nil {
		o.R = &childTagR{
			ChildTag: related,
		}
	} else {
		o.R.ChildTag = related
	}

	if related.R == nil {
		related.R = &tagR{
			ChildTagChildTags: ChildTagSlice{o},
		}
	} else {
		related.R.ChildTagChildTags = append(related.R.ChildTagChildTags, o)
	}

	return nil
}

// RemoveChildTag relationship.
// Sets o.R.ChildTag to nil.
// Removes o from all passed in related items' relationships struct.
func (o *ChildTag) RemoveChildTag(ctx context.Context, exec boil.ContextExecutor, related *Tag) error {
	var err error

	queries.SetScanner(&o.ChildTagID, nil)
	if _, err = o.Update(ctx, exec, boil.Whitelist("child_tag_id")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	if o.R != nil {
		o.R.ChildTag = nil
	}
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.ChildTagChildTags {
		if queries.Equal(o.ChildTagID, ri.ChildTagID) {
			continue
		}

		ln := len(related.R.ChildTagChildTags)
		if ln > 1 && i < ln-1 {
			related.R.ChildTagChildTags[i] = related.R.ChildTagChildTags[ln-1]
		}
		related.R.ChildTagChildTags = related.R.ChildTagChildTags[:ln-1]
		break
	}
	return nil
}

// SetParentTag of the childTag to the related item.
// Sets o.R.ParentTag to related.
// Adds o to related.R.ParentTagChildTags.
func (o *ChildTag) SetParentTag(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Tag) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"child_tags\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"parent_tag_id"}),
		strmangle.WhereClause("\"", "\"", 2, childTagPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	queries.Assign(&o.ParentTagID, related.ID)
	if o.R == nil {
		o.R = &childTagR{
			ParentTag: related,
		}
	} else {
		o.R.ParentTag = related
	}

	if related.R == nil {
		related.R = &tagR{
			ParentTagChildTags: ChildTagSlice{o},
		}
	} else {
		related.R.ParentTagChildTags = append(related.R.ParentTagChildTags, o)
	}

	return nil
}

// RemoveParentTag relationship.
// Sets o.R.ParentTag to nil.
// Removes o from all passed in related items' relationships struct.
func (o *ChildTag) RemoveParentTag(ctx context.Context, exec boil.ContextExecutor, related *Tag) error {
	var err error

	queries.SetScanner(&o.ParentTagID, nil)
	if _, err = o.Update(ctx, exec, boil.Whitelist("parent_tag_id")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	if o.R != nil {
		o.R.ParentTag = nil
	}
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.ParentTagChildTags {
		if queries.Equal(o.ParentTagID, ri.ParentTagID) {
			continue
		}

		ln := len(related.R.ParentTagChildTags)
		if ln > 1 && i < ln-1 {
			related.R.ParentTagChildTags[i] = related.R.ParentTagChildTags[ln-1]
		}
		related.R.ParentTagChildTags = related.R.ParentTagChildTags[:ln-1]
		break
	}
	return nil
}

// ChildTags retrieves all the records using an executor.
func ChildTags(mods ...qm.QueryMod) childTagQuery {
	mods = append(mods, qm.From("\"child_tags\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"child_tags\".*"})
	}

	return childTagQuery{q}
}

// FindChildTag retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindChildTag(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*ChildTag, error) {
	childTagObj := &ChildTag{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"child_tags\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, childTagObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from child_tags")
	}

	if err = childTagObj.doAfterSelectHooks(ctx, exec); err != nil {
		return childTagObj, err
	}

	return childTagObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *ChildTag) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no child_tags provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(childTagColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	childTagInsertCacheMut.RLock()
	cache, cached := childTagInsertCache[key]
	childTagInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			childTagAllColumns,
			childTagColumnsWithDefault,
			childTagColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(childTagType, childTagMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(childTagType, childTagMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"child_tags\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"child_tags\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into child_tags")
	}

	if !cached {
		childTagInsertCacheMut.Lock()
		childTagInsertCache[key] = cache
		childTagInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the ChildTag.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *ChildTag) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	childTagUpdateCacheMut.RLock()
	cache, cached := childTagUpdateCache[key]
	childTagUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			childTagAllColumns,
			childTagPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update child_tags, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"child_tags\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, childTagPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(childTagType, childTagMapping, append(wl, childTagPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update child_tags row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for child_tags")
	}

	if !cached {
		childTagUpdateCacheMut.Lock()
		childTagUpdateCache[key] = cache
		childTagUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q childTagQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for child_tags")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for child_tags")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ChildTagSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), childTagPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"child_tags\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, childTagPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in childTag slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all childTag")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *ChildTag) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no child_tags provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(childTagColumnsWithDefault, o)

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

	childTagUpsertCacheMut.RLock()
	cache, cached := childTagUpsertCache[key]
	childTagUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			childTagAllColumns,
			childTagColumnsWithDefault,
			childTagColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			childTagAllColumns,
			childTagPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert child_tags, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(childTagPrimaryKeyColumns))
			copy(conflict, childTagPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"child_tags\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(childTagType, childTagMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(childTagType, childTagMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert child_tags")
	}

	if !cached {
		childTagUpsertCacheMut.Lock()
		childTagUpsertCache[key] = cache
		childTagUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single ChildTag record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *ChildTag) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no ChildTag provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), childTagPrimaryKeyMapping)
	sql := "DELETE FROM \"child_tags\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from child_tags")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for child_tags")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q childTagQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no childTagQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from child_tags")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for child_tags")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ChildTagSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(childTagBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), childTagPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"child_tags\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, childTagPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from childTag slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for child_tags")
	}

	if len(childTagAfterDeleteHooks) != 0 {
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
func (o *ChildTag) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindChildTag(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ChildTagSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ChildTagSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), childTagPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"child_tags\".* FROM \"child_tags\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, childTagPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ChildTagSlice")
	}

	*o = slice

	return nil
}

// ChildTagExists checks if the ChildTag row exists.
func ChildTagExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"child_tags\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if child_tags exists")
	}

	return exists, nil
}
